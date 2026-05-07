package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(secret),
	}
	return echojwt.WithConfig(config)
}

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
