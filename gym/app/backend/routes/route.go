package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	jwt "gym/app/backend/utils/middleware"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	eJwt := e.Group("")
	eJwt.Use(jwt.JWTMiddleware())
	UserRoute(e, db, eJwt)
	RoleRoute(e, db, eJwt)
	MembershipPlanRoute(e, db, eJwt)
}
