package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/handler"
	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/model"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/auth"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/logger"
	"github.com/mendesbarreto/go-my-coffe-shop/proto/gen"
)

var publicMethods []string = []string{
	"/api.v1.mycoffeshop.user.UserService/SignIn",
	"/api.v1.mycoffeshop.user.UserService/SignUp",
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

	slog.Info("User Module Starting up")

	slog.Info("Starting Third Parties")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	slog.Info("User Module context created")
	config := config.GetConfig()

	infra.SetupDependecies(ctx, config)
	slog.Info("User Module Config loaded", config)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			auth.GetUnaryGrpcInterceptor(publicMethods, func() (interface{}, error) {
				return model.User{}, nil
			}),
			logger.GetUnaryGrpcInterceptor(),
		),
	)
	serverAddress := fmt.Sprintf("%s:%s", config.Host, config.Port)
	network := "tcp"

	handler.NewUserGRPCHandler(grpcServer, config)

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

	err = gen.RegisterUserServiceHandler(context.Background(), mux, conn)

	if err != nil {
		slog.Error("Problem to start grpc Gateway", "error message=", err.Error())
	}

	restServer := http.Server{
		Handler: logger.WithLogger(mux),
	}

	// TODO: Add the port to the configuration file from module user
	slog.Info("Start to listen :8081")
	restLis, err := net.Listen(network, "0.0.0.0:8081")
	if err != nil {
		slog.Error("Problem to start listens port 8081 maybe the port is in user", "error message=", err.Error())
	}

	err = restServer.Serve(restLis)
	if err != nil {
		slog.Error("Problem to start listens port 8081 maybe the port is in user", "error message=", err.Error())
	}

	defer func() {
		conn.Close()
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("Problem to start gRPC server", err, "network", network, "address", serverAddress)
	}
}
