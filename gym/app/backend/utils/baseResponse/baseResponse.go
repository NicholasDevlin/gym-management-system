package baseresponse

import (
	"gym/app/backend/utils/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Success  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	BaseResponse BaseResponse `json:"response"`
	CurrentPage int `json:"currentPage"`
	NextPage int `json:"nextPage"`
	PrevPage int `json:"prevPage"`
	AllPages int `json:"allPages"`
}

func NewSuccessPaginationResponse(c echo.Context, data interface{}, currentPage, nextPage, prevPage, allPages int) error {
	return c.JSON(http.StatusOK, Pagination{
		CurrentPage: currentPage,
		NextPage: nextPage,
		PrevPage: prevPage,
		AllPages: allPages,
		BaseResponse: BaseResponse{
			Success:  true,
			Message: "Success",
			Data:    data,
		},
	})
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Success:  true,
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(c echo.Context, err error) error {
	return c.JSON(errors.GetCodeError(err), BaseResponse{
		Success:  false,
		Message: err.Error(),
		Data:    nil,
	})
}