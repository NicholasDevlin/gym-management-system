package controller

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/feature/services"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/errors"
	"gym/app/backend/utils/middleware"
	"gym/app/backend/utils/pagination"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) RegisterUsers(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.RegisterUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	var role string
	token, err := middleware.CreateToken(res.Id, res.DisplayName, role)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(e, token)
	return baseresponse.NewSuccessResponse(e, res)
}

func (u *UserController) GetAllUser(e echo.Context) error {
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
