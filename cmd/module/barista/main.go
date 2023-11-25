package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/barista/config"
	"github.com/mendesbarreto/go-my-coffe-shop/internal/barista/handler"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/auth"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/client"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/redis"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/logger"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/util"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var publicMethods []string = []string{}

// TODO: Refactory code to remove all warnings
func createUserContextAndCache(ctx context.Context, jwt string) (context.Context, error) {
	claims := &model.ModuleClaims{}
	err := redis.Get(ctx, jwt, claims)

	if err == nil {
		slog.Info("A redis cache was found: %v", claims.User)
		return context.WithValue(ctx, "user", claims.User), nil
	}

	token, claims, err := util.DecodeJWT(jwt)
	if err != nil {
		return nil, err
	}

	cfg := config.GetConfig()
	userClient, err := client.GetUserServiceClient(cfg.UserServiceHost, cfg.UserServicePort)
	if err != nil {
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.NotFound, "Missing headers")
	}

	outgoingCtx := metadata.NewOutgoingContext(ctx, md)
	updatedUser, err := userClient.GetMe(outgoingCtx, &gen.EmptyRequest{})
	if err != nil || updatedUser == nil {
		return nil, status.Errorf(codes.Unauthenticated, "User was not found %v", err.Error())
	}

	cacheDuration, err := util.GetDurationFromJWT(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "It was imposible to get the expiration from token %v", err.Error())
	}

	userId, err := primitive.ObjectIDFromHex(updatedUser.GetUserId())
	if err != nil || updatedUser == nil {
		return nil, status.Errorf(codes.Unauthenticated, "It was imposible to get user id from service client %v", err.Error())
	}

	claims.User = model.User{
		ID:    userId,
		Name:  updatedUser.GetName(),
		Email: updatedUser.GetEmail(),
	}
	err = redis.Save(ctx, jwt, claims, cacheDuration)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "It was impossible to use redis to save cache %v", err.Error())
	}

	slog.Info("[Authorization]", "jwt=", token.Raw, "user=", claims.User, "expDuration=", cacheDuration)
	return context.WithValue(ctx, "user", claims.User), nil
}

func main() {
	// This function serves to control the maximum number of operating system threads that can concurrently execute user-level Go code.
	// When a Go program starts, it automatically sets the number of threads to the maximum available on the system.
	// This effectively determines the number of Goroutines that can run in parallel.

	// In Kubernetes, all the CPU cores available on a node are visible to its pods. If you set a pod's CPU limit to 1 core,
	// but your node has 64 CPU cores, your Go application will utilize the full node resources and set GOMAXPROCS to 64.

	// The purpose of maxprocs.Set() is to ensure that only the number of CPU cores available to the pod is allocated for execution
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("Problem to set the threads available on the system", err)
	}

	slog.Info("Barista Module Starting up")

	slog.Info("Starting Third Parties")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	slog.Info("User Module context created")
	config := config.GetConfig()

	infra.SetupDependecies(ctx, config.Name, config.MongoDb.URI, config.Redis.URI)
	slog.Info("User Module Config loaded", config)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			auth.GetUnaryGrpcInterceptor(publicMethods, createUserContextAndCache),
			logger.GetUnaryGrpcInterceptor(),
		),
	)
	serverAddress := fmt.Sprintf("%s:%s", config.Host, config.Port)
	network := "tcp"

	handler.NewBaristaGRPCHandler(grpcServer, config)

	lis, err := net.Listen(network, serverAddress)
	if err != nil {
		slog.Error("Failed to start listen to address", err, "network", network, "serverAddress", serverAddress)
		cancel()
		infra.CleanUpDependcies(ctx)
		<-ctx.Done()
	}

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			slog.Error("Problem to initialize grpc server", "message=", err.Error())
		}
	}()

	slog.Info("User Module Server", "listen", serverAddress)
	defer func() {
		if err := lis.Close(); err != nil {
			slog.Error("Problem to close server", err, "network", network, "address", serverAddress)
		}
		cancel()
		infra.CleanUpDependcies(ctx)
		<-ctx.Done()
	}()

	// Create mux for grPC Gateway
	mux := runtime.NewServeMux()

	conn, err := grpc.DialContext(context.Background(), serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Problem to start grpc Gateway", "error message=", err.Error())
	}

	err = gen.RegisterProductServiceHandler(context.Background(), mux, conn)

	if err != nil {
		slog.Error("Problem to start grpc Gateway", "error message=", err.Error())
	}

	restServer := http.Server{
		Handler: logger.WithLogger(mux),
	}

	// TODO: Add the port to the configuration file from module user
	restServerAddress := fmt.Sprintf("%s:%s", config.Host, config.RestPort)
	slog.Info("Start to listen", "PORT", config.RestPort)
	restLis, err := net.Listen(network, restServerAddress)
	if err != nil {
		slog.Error("Problem to start listens maybe the port is in use", "PORT", config.RestPort, "error message=", err.Error())
	}

	err = restServer.Serve(restLis)
	if err != nil {
		slog.Error("Problem to start listens port 8082 maybe the port is in user", "error message=", err.Error())
	}

	defer func() {
		conn.Close()
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("Problem to start gRPC server", err, "network", network, "address", serverAddress)
	}
}
