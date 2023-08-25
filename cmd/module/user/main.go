package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"log/slog"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	handler "github.com/mendesbarreto/go-my-coffe-shop/internal/user"
)

func main() {
	//
	// This function is used to keep the number of operating system threads that can execute user-level Go code at any one time.
	// When Go starts it's automatically set the number of thread to max available on the system. That means the number of goroutines that can actually run in parallel.
	// In Kubernetes, all the available CPU cores on the node are visible by its pods If you set a pod CPU limit to 1 core
	// but your node has 64 cores of CPU, your Go app will grab the actual node resource and set GOMAXPROC to 64
	// So the maxprocs Set() helps to assign only the number of cores available for the pod.
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("Problem to set the threads available on the system", err)
	}

	slog.Info("User Module Starting up")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	slog.Info("User Module context created")
	config := config.GetConfig()

	slog.Info("User Module Config loaded", config)

	grpcServer := grpc.NewServer()
	serverAddress := fmt.Sprintf("%s:%s", config.Grcp.Host, config.Grcp.Port)
	network := "tcp"

	handler.NewUserGRPCHandler(grpcServer, config)

	lis, err := net.Listen(network, serverAddress)
	if err != nil {
		slog.Info("Failed to start listen to address", err, "network", network, "serverAddress", serverAddress)
		cancel()
		<-ctx.Done()
	}

	slog.Info("User Module Server", "listen", serverAddress)

	defer func() {
		if err := lis.Close(); err != nil {
			slog.Error("Problem to close server", err, "network", network, "address", serverAddress)
		}
		cancel()
		<-ctx.Done()
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("Problem to start gRPC server", err, "network", network, "address", serverAddress)
	}
}
