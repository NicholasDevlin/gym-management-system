package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repositories.NewUsersRepository(db)
	roleRepository := repositories.NewRoleRepository(db)
	service := services.NewUserService(repository, roleRepository)
	controller := controller.NewUserController(service)

	e.POST("/user/register", controller.RegisterUsers)
	e.POST("/user/login", controller.LoginUser)

	eJwt.GET("/user", controller.GetAllUser)
	eJwt.GET("/user/:id", controller.GetUser)
	eJwt.PUT("/user/:id", controller.UpdateUser)
	eJwt.DELETE("/user/:id", controller.DeleteUser)
}
