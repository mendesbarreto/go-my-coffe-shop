package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GenerateJwt(user *model.User) (string, error) {
	expTime := time.Now().Add(8760 * time.Hour)

	if user == nil {
		return "", status.Error(codes.NotFound, "User was not found")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		model.ModuleClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
			User: *user,
		})

	tokenString, err := token.SignedString([]byte(config.GetConfig().AuthSecrete))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
