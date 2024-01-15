package controller

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/feature/services"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/errors"
	"gym/app/backend/utils/middleware"

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