package auth

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getJWT(md metadata.MD) (*string, error) {
	authHeaders, exists := md["grpcgateway-authorization"]

	if !exists || len(authHeaders) == 0 || len(authHeaders[0]) == 0 {
		return nil, status.Error(codes.Unauthenticated, "No Authorization bearer was found")
	}

	values := strings.Split(authHeaders[0], " ")
	return &values[len(values)-1], nil
}

func hasNoAuthCheck(method string, methods []string) bool {
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}

func GetUnaryGrpcInterceptor(methods []string,
	createUserContextAndCache func(cxt context.Context, jwt string) (context.Context, error),
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if hasNoAuthCheck(info.FullMethod, methods) {
			slog.Info("[Authorization] No auth needed for %v", info.FullMethod)
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(codes.Unknown, "Failed to get metadata from context")
		}

		tokenString, err := getJWT(md)
		if err != nil {
			return nil, err
		}

		ctx, err = createUserContextAndCache(ctx, *tokenString)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
