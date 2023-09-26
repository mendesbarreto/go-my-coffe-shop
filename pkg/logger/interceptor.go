package logger

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUnaryGrpcInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		result, err := handler(ctx, req)
		duration := time.Since(start)

		statusCode := codes.Unknown

		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		logger := slog.Info

		if err != nil {
			logger = slog.Error
		}

		logger("gRPC", "method", info.FullMethod, "req", req, "res", result, "duration", duration, "status", statusCode.String())

		return result, err
	}
}
