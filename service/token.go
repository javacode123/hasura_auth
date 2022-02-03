package service

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hasura_auth/config"
	"github.com/hasura_auth/model"
	"time"
)

func CreateToken(ctx context.Context, claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":                          time.Now().Add(config.GlobalConfig.JWTExpireIn).Unix(),
		"iat":                          time.Now().Unix(),
		"iss":                          config.GlobalConfig.AppName,
		"https://hasura.io/jwt/claims": claims,
	})
	return token.SignedString([]byte(config.GlobalConfig.JWTSecret))
}
