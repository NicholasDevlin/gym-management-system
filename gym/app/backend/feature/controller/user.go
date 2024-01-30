package controller

import (
	"gym/app/backend/feature/services"
	"gym/app/backend/models/user"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/errors"
	"gym/app/backend/utils/middleware"
	"gym/app/backend/utils/pagination"
	"strconv"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type userController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *userController {
	return &userController{userService}
}

func (u *userController) RegisterUsers(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.RegisterUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.DisplayName, res.Role.Role)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(e, token)
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) GetAllUser(e echo.Context) error {
	var filter user.UserReq
	filter.DisplayName = e.QueryParam("name")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	pageSize, _ := strconv.Atoi(e.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecord, err := u.userService.GetAllUser(filter, page, pageSize)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	currentPage := pagination.CurrentPage(page, pageSize, totalRecord)
	nextPage := pagination.GetNextPage(currentPage, totalRecord)
	prevPage := pagination.GetPrevPage(currentPage)

	return baseresponse.NewSuccessPaginationResponse(e, res, currentPage, nextPage, prevPage, totalRecord)
}

func (u *userController) LoginUser(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.LoginUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.DisplayName, res.Role.Role)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(e, token)
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) GetUser(e echo.Context) error {
	var input user.UserReq

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := u.userService.GetUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) UpdateUser(e echo.Context) error {
	userId, role, err := middleware.ExtractToken(e)
	if role != consts.ADMIN {
		return  baseresponse.NewErrorResponseUnauthorize(e)
	}

	var input user.UserReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))

	if userId != uuid {
		return  baseresponse.NewErrorResponseUnauthorize(e)
	}
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	input.UUID = uuid

	res, err := u.userService.UpdateUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *userController) DeleteUser(e echo.Context) error {
	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}

	res, err := u.userService.DeleteUser(uuid)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	return baseresponse.NewSuccessResponse(e, res)
}