package middlewares

import (
	"alta/temanpetani/app/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId uint64, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["userRole"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func ExtractToken(e echo.Context) (uint64, string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		userRole := claims["userRole"].(string)
		return uint64(userId), userRole, nil
	}
	return 0, "", errors.New("failed to extract jwt-token")
}
