package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HasJwtExpired(token *jwt.Token) error {
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

func DecodeJWT(tokenString string) (*jwt.Token, *model.ModuleClaims, error) {
	claims := &model.ModuleClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.GetConfig().AuthSecrete), nil
	})
	if err != nil {
		return nil, nil, status.Errorf(codes.Unauthenticated, "The token providade is invalid: %v", err.Error())
	}

	if err = HasJwtExpired(token); err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}
