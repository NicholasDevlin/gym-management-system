package routes

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, db *gorm.DB) {
	godotenv.Load("")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

}
