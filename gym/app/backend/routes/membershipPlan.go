package routes

import (
	"gym/app/backend/feature/controller"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/feature/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MembershipPlanRoute(e *echo.Echo, db *gorm.DB, eJwt *echo.Group) {
	repository := repositories.NewMembershipPlanRepository(db)
	service := services.NewMembershipPlanService(repository)
	controller := controller.NewMembershipPlanController(service)

	//access without token
	eJwt.POST("/membership-plan", controller.CreateMembershipPlan)
	e.GET("/membership-plan", controller.GetAllMembershipPlan)
	e.GET("/membership-plan/:id", controller.GetMembershipPlan)
	eJwt.PUT("/membership-plan/:id", controller.UpdateMembershipPlan)
	eJwt.DELETE("/membership-plan/:id", controller.DeleteMembershipPlan)
}
