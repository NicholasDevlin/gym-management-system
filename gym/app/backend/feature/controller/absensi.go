package controller

import (
	"gym/app/backend/feature/services"
	"gym/app/backend/models/absensi"
	baseresponse "gym/app/backend/utils/baseResponse"

	"github.com/labstack/echo/v4"
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

	res, err := a.absensiService.CreateAbsensi(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

// func (r *absensiController) GetAllRole(e echo.Context) error {
// 	var input role.RoleReq
// 	input.Role = e.QueryParam("role")

// 	res, err := r.absensiService.GetAllRole(input)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}

// 	return baseresponse.NewSuccessResponse(e, res)
// }

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

// func (r *absensiController) UpdateRole(e echo.Context) error {
// 	var input role.RoleReq
// 	e.Bind(&input)
// 	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}
// 	input.Id = uint(id)

// 	res, err := r.absensiService.UpdateRole(input)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}
// 	return baseresponse.NewSuccessResponse(e, res)
// }

// func (r *absensiController) DeleteUser(e echo.Context) error {
// 	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}

// 	res, err := r.absensiService.DeleteRole(uint(id))
// 	if err != nil {
// 		return baseresponse.NewErrorResponse(e, err)
// 	}
// 	return baseresponse.NewSuccessResponse(e, res)
// }
