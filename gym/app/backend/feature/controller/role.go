package controller

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/feature/services"
	baseresponse "gym/app/backend/utils/baseResponse"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	roleService services.IRoleService
}

func NewRoleController(roleService services.IRoleService) *RoleController {
	return &RoleController{roleService}
}

func (u *RoleController) CreateRole(e echo.Context) error {
	var input role.RoleReq
	e.Bind(&input)

	res, err := u.roleService.CreateRole(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	return baseresponse.NewSuccessResponse(e, res)
}
