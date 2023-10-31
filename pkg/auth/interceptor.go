package auth

import (
	"context"
	"log/slog"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
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

func GetUnaryGrpcInterceptor(methods []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if hasNoAuthCheck(info.FullMethod, methods) {
			slog.Info("[Authorization] No auth needed for %v", info.FullMethod)
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(codes.Unknown, "failed to get metadata from context")
		}

		tokenString, err := getJWT(md)
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(*tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, status.Errorf(codes.Unauthenticated, "Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(config.GetConfig().AuthSecrete), nil
		})
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "The token providade is invalid: %v", err.Error())
		}

		clains, err := token.Claims.GetAudience()
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "The token providade does not have any clains %v", err.Error())
		}

		// TODO: HERE WE NEED TO FETCH THE USER. Maybe a Factory??

		slog.Info("[Authorization]", "jwt=", token.Raw, "user=", clains)

		return handler(ctx, req)
	}
}
