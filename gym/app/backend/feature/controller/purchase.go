package controller

import (
	"gym/app/backend/feature/services"
	"gym/app/backend/models/purchase"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/middleware"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type purchaseController struct {
	purchaseService services.IPurchaseService
}

func NewPurchaseController(purchaseService services.IPurchaseService) *purchaseController {
	return &purchaseController{purchaseService}
}

func (p *purchaseController) CreatePurchase(e echo.Context) error {
	userUUID, role, err := middleware.ExtractToken(e)
	var input purchase.PurchaseReq
	e.Bind(&input)
	if role == consts.USER {
		input.User.UUID = userUUID
	}

	res, err := p.purchaseService.CreatePurchase(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (p *purchaseController) GetAllPurchase(e echo.Context) error {
	var input purchase.PurchaseReq
	// input.PurchaseDate.Date() = e.QueryParam("purchaseDate")

	res, err := p.purchaseService.GetAllPurchase(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (p *purchaseController) GetPurchase(e echo.Context) error {
	var input purchase.PurchaseReq

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := p.purchaseService.GetPurchase(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (p *purchaseController) UpdatePurchase(e echo.Context) error {
	var input purchase.PurchaseReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := p.purchaseService.UpdatePurchase(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (p *purchaseController) DeletePurchase(e echo.Context) error {
	_, _, err := middleware.ExtractToken(e)

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := p.purchaseService.DeletePurchase(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}
