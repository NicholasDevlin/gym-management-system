package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoleRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repository)
	controller := controller.NewRoleController(service)

	//access without token
	e.POST("/role", controller.CreateRole)
	e.GET("/role", controller.GetAllRole)
}
