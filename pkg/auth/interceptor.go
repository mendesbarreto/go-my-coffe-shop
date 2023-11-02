package auth

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/infra/redis"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
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

func validateJWT(token *jwt.Token) error {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return err
	}

	if exp == nil || exp.Before(time.Now()) {
		return status.Error(codes.Unauthenticated, "The token expired")
	}

	return nil
}

func GetDurationFromJWT(token *jwt.Token) (time.Duration, error) {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, err
	}
	return time.Since(exp.Time).Abs(), nil
}

func GetUnaryGrpcInterceptor(methods []string, getValueToCacheByUserId func(cxt context.Context, userId string) (interface{}, error)) grpc.UnaryServerInterceptor {
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

		claims := &model.ModuleClaims{}
		token, err := jwt.ParseWithClaims(*tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, status.Errorf(codes.Unauthenticated, "Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(config.GetConfig().AuthSecrete), nil
		})
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "The token providade is invalid: %v", err.Error())
		}

		if err = validateJWT(token); err != nil {
			return nil, err
		}

		valueToCache, err := getValueToCacheByUserId(ctx, claims.User.ID.String())
		slog.Info("<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>", valueToCache)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Problem to fetch user %v", err.Error())
		}

		redisCacheDuration, err := GetDurationFromJWT(token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "It was imposible to get the expiration from token %v", err.Error())
		}

		err = redis.Save(ctx, *tokenString, valueToCache, redisCacheDuration)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "It was impossible to use redis to save cache %v", err.Error())
		}

		slog.Info("[Authorization]", "jwt=", token.Raw, "user=", claims.User, "expDuration=", redisCacheDuration)

		return handler(ctx, req)
	}
}
