package controller

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/feature/services"
	baseresponse "gym/app/backend/utils/baseResponse"
	"gym/app/backend/utils/errors"
	"gym/app/backend/utils/middleware"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) RegisterUsers(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.RegisterUser(input)
	if err != nil {
		return baseresponse.NewErrorResponse(e, err)
	}
	var role string
	token, err := middleware.CreateToken(res.Id, role, 0)
	if err != nil {
		return baseresponse.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(e, token)
	return baseresponse.NewSuccessResponse(e, res)
}
