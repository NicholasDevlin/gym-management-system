package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	detailRepo := repositories.NewTransactionDetailRepository(db)
	detailMemberRepo := repositories.NewTransactionMemberDetailRepository(db)
	repository := repositories.NewTransactionRepository(db, detailRepo, detailMemberRepo)
	userRepository := repositories.NewUsersRepository(db)
	membershipPlanRepository := repositories.NewMembershipPlanRepository(db)
	transactionDetailRepository := repositories.NewTransactionDetailRepository(db)
	service := services.NewTransactionService(repository, userRepository, membershipPlanRepository, transactionDetailRepository)
	controller := controller.NewTransactionController(service)

	eJwt.POST("/transaction", controller.SaveTransaction)
	eJwt.GET("/transaction", controller.GetAllTransaction)
	eJwt.GET("/transaction/:id", controller.GetTransaction)
	eJwt.POST("/transaction/:id", controller.SaveTransaction)
	eJwt.DELETE("/transaction/:id", controller.DeleteTransaction)
}
