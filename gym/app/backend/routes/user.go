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
	service := services.NewUserService(repository)
	controller := controller.NewUserController(service)

	//access without token
	e.POST("/users/register", controller.RegisterUsers)
	// e.POST("/users/login", controller.LoginUsers)

	// // superadmin can access
	// eJwt.GET("/users", controller.GetAllUser)
	// eJwt.GET("/users/:id", controller.GetUser)
	// eJwt.PUT("/users/:id", controller.UpdateUser)
	// eJwt.DELETE("/users/:id", controller.DeleteUser)
}
