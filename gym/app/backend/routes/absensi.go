package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AbsensiRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repositories.NewAbsensiRepository(db)
	repositoryUser := repositories.NewUsersRepository(db)
	service := services.NewAbsensiService(repository, repositoryUser)
	controller := controller.NewAbsensiController(service)

	eJwt.POST("/absensi", controller.CreateAbsensi)
	eJwt.GET("/absensi", controller.GetAllAbsensi)
	// eJwt.GET("/role/:id", controller.GetRole)
	eJwt.PUT("/absensi/:id", controller.UpdateAbsensi)
	eJwt.DELETE("/absensi/:id", controller.DeleteAbsensi)
}
