package controller

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/feature/services"
	baseresponse "gym/app/backend/utils/baseResponse"
	"strconv"

	"github.com/labstack/echo/v4"
)

type roleController struct {
	roleService services.IRoleService
}

func NewRoleController(roleService services.IRoleService) *roleController {
	return &roleController{roleService}
}

func (r *roleController) CreateRole(e echo.Context) error {
	var input role.RoleReq
	e.Bind(&input)

	res, err := r.roleService.CreateRole(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (r *roleController) GetAllRole(e echo.Context) error {
	var input role.RoleReq
	input.Role = e.QueryParam("role")

	res, err := r.roleService.GetAllRole(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (r *roleController) GetRole(e echo.Context) error {
	var input role.RoleReq

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.Id = uint(id)
	res, err := r.roleService.GetRole(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}

func (r *roleController) UpdateRole(e echo.Context) error {
	var input role.RoleReq
	e.Bind(&input)
	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.Id = uint(id)

	res, err := r.roleService.UpdateRole(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (r *roleController) DeleteUser(e echo.Context) error {
	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := r.roleService.DeleteRole(uint(id))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}
