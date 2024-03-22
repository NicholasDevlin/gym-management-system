package controller

import (
	"gym/app/backend/feature/services"
	"gym/app/backend/models/transaction"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/middleware"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type transactionController struct {
	transactionService services.ITransactionService
}

func NewTransactionController(transactionService services.ITransactionService) *transactionController {
	return &transactionController{transactionService}
}

func (t *transactionController) CreateTransaction(e echo.Context) error {
	userUUID, role, err := middleware.ExtractToken(e)
	var input transaction.TransactionReq
	e.Bind(&input)
	if role == consts.USER {
		input.User.UUID = userUUID
	}

	res, err := t.transactionService.CreateTransaction(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (t *transactionController) GetAllTransaction(e echo.Context) error {
	var input transaction.TransactionReq
	// input.TransactionDate.Date() = e.QueryParam("TransactionDate")

	res, err := t.transactionService.GetAllTransaction(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (t *transactionController) GetTransaction(e echo.Context) error {
	var input transaction.TransactionReq

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := t.transactionService.GetTransaction(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (t *transactionController) SaveTransaction(e echo.Context) error {
	var input transaction.TransactionReq
	e.Bind(&input)
	if e.Param("id") != "" {
		input.UUID, _ = uuid.FromString(e.Param("id"))
	}

	res, err := t.transactionService.SaveTransaction(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (t *transactionController) DeleteTransaction(e echo.Context) error {
	_, _, err := middleware.ExtractToken(e)

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := t.transactionService.DeleteTransaction(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}
