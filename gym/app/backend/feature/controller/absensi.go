package controller

import (
	"gym/app/backend/feature/services"
	"gym/app/backend/models/absensi"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/middleware"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type absensiController struct {
	absensiService services.IAbsensiService
}

func NewAbsensiController(absensiService services.IAbsensiService) *absensiController {
	return &absensiController{absensiService}
}

func (a *absensiController) CreateAbsensi(e echo.Context) error {
	var input absensi.AbsensiReq
	e.Bind(&input)

	res, err := a.absensiService.SaveAbsensi(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (a *absensiController) GetAllAbsensi(e echo.Context) error {
	var input absensi.AbsensiFilter
	if err := e.Bind(&input); err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := a.absensiService.GetAllAbsensi(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

// func (r *absensiController) GetRole(e echo.Context) error {
// 	var input role.RoleReq

// 	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}
// 	input.Id = uint(id)
// 	res, err := r.absensiService.GetRole(input)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}

// 	return baseresponse.NewSuccessResponse(e, res)
// }

func (a *absensiController) UpdateAbsensi(e echo.Context) error {
	var input absensi.AbsensiReq
	
	if err := e.Bind(&input); err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	_, role, err := middleware.ExtractToken(e)
	if role != consts.ADMIN {
		return baseresponse.NewErrorResponseUnauthorize(e)
	} else if err != nil {
		return baseresponse.NewErrorResponseUnauthorize(e)
	}

	if e.Param("id") != "" {
		input.UUID, _ = uuid.FromString(e.Param("id"))
	}

	res, err := a.absensiService.SaveAbsensi(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (a *absensiController) DeleteAbsensi(e echo.Context) error {
	_, role, err := middleware.ExtractToken(e)

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	if role != consts.ADMIN {
		return baseresponse.NewErrorResponseUnauthorize(e)
	}
	err = a.absensiService.DeleteAbsensi(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, absensi.AbsensiRes{})
}
