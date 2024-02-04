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

	eJwt.POST("/role", controller.CreateRole)
	eJwt.GET("/role", controller.GetAllRole)
	eJwt.GET("/role/:id", controller.GetRole)
	eJwt.PUT("/role/:id", controller.UpdateRole)
	eJwt.DELETE("/role/:id", controller.DeleteUser)
}
