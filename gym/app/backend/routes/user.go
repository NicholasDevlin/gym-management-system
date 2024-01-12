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
	handler := controller.NewUserHandler(service)

	//access without token
	e.POST("/users/register", handler.RegisterUsers)
	// e.POST("/users/login", handler.LoginUsers)

	// // superadmin can access
	// eJwt.GET("/users", handler.GetAllUser)
	// eJwt.GET("/users/:id", handler.GetUser)
	// eJwt.PUT("/users/:id", handler.UpdateUser)
	// eJwt.DELETE("/users/:id", handler.DeleteUser)
}