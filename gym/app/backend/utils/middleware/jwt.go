package middleware

import (
	"errors"
	// "net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func CreateToken(userId uuid.UUID, name, role string) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{}
	claims["uuid"] = userId
	claims["role"] = role
	// claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	// claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func JWTMiddleware() echo.MiddlewareFunc {
	godotenv.Load()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
		SigningMethod: "HS256",
	})
}

// func SetTokenCookie(e echo.Context, token string) {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "token"
// 	cookie.Value = token
// 	cookie.Path = "/"

// 	e.SetCookie(cookie)
// }

func ExtractToken(e echo.Context) (uuid.UUID, string, error) {
	user, ok := e.Get("user").(*jwt.Token)

	if !ok {
		return uuid.UUID{}, "", errors.New("invalid token claims")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, "", errors.New("invalid token claims")
	}

	userId, ok := claims["uuid"].(string)

	if !ok {
		return uuid.UUID{}, "", errors.New("invalid token claims")
	}
	userUUID, _ := uuid.FromString(userId)

	role, ok := claims["role"].(string)
	if !ok {
		return uuid.UUID{}, "", errors.New("invalid token claims")
	}

	// name, okName := claims["name"].(string)
	// if !okName {
	// 	return 0, 0, 0, "", "", errors.New("invalid token claims")
	// }

	// role, okRole := claims["role"].(string)
	// if !okRole {
	// 	return 0, 0, 0, "", errors.New("invalid token claims")
	// }

	return userUUID, role, nil
}
