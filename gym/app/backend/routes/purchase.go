package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PurchaseRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repositories.NewPurchaseRepository(db)
	userRepository := repositories.NewUsersRepository(db)
	service := services.NewPurchaseService(repository, userRepository)
	controller := controller.NewPurchaseController(service)

	eJwt.POST("/purchase", controller.CreatePurchase)
	eJwt.GET("/purchase", controller.GetAllPurchase)
	eJwt.GET("/purchase/:id", controller.GetPurchase)
	eJwt.PUT("/purchase/:id", controller.UpdatePurchase)
	eJwt.DELETE("/purchase/:id", controller.DeletePurchase)
}