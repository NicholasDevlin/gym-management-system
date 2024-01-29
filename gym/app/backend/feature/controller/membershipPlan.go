package controller

import (
	"gym/app/backend/feature/services"
	membershipplan "gym/app/backend/models/membershipPlan"
	baseresponse "gym/app/backend/utils/baseResponse"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type membershipPlanController struct {
	membershipPlanService services.IMembershipPlanService
}

func NewMembershipPlanController(membershipPlanService services.IMembershipPlanService) *membershipPlanController {
	return &membershipPlanController{membershipPlanService}
}

func (mp *membershipPlanController) CreateMembershipPlan(e echo.Context) error {
	var input membershipplan.MembershipPlanReq
	e.Bind(&input)

	res, err := mp.membershipPlanService.CreateMembershipPlan(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (mp *membershipPlanController) GetAllMembershipPlan(e echo.Context) error {
	var input membershipplan.MembershipPlanReq
	input.Name = e.QueryParam("role")

	res, err := mp.membershipPlanService.GetAllMembershipPlan(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (mp *membershipPlanController) GetMembershipPlan(e echo.Context) error {
	var input membershipplan.MembershipPlanReq

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := mp.membershipPlanService.GetMembershipPlan(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (mp *membershipPlanController) UpdateMembershipPlan(e echo.Context) error {
	var input membershipplan.MembershipPlanReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := mp.membershipPlanService.UpdateMembershipPlan(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (mp *membershipPlanController) DeleteMembershipPlan(e echo.Context) error {
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := mp.membershipPlanService.DeleteMembershipPlan(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}