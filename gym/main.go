package main

import (
	"fmt"
	"gym/app/backend/config"
	"gym/app/backend/drivers/mysql"
	"gym/app/backend/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig, dbConfig := config.InitConfig()
	db := mysql.StartDB(dbConfig)

	var e *echo.Echo
	e = echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	routes.Route(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}
