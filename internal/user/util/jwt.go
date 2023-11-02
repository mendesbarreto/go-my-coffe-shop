package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mendesbarreto/go-my-coffe-shop/cmd/module/user/config"
	"github.com/mendesbarreto/go-my-coffe-shop/internal/user/model"
)

func GenerateJwt(user *model.User) (string, error) {
	expTime := time.Now().Add(8760 * time.Hour)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		model.UserModuleClains{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
			User: model.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		})

	tokenString, err := token.SignedString([]byte(config.GetConfig().AuthSecrete))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
