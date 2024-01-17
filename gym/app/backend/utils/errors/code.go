package errors

import (
	"net/http"
)

func GetCodeError(err error) int {
	switch err {
	case ERR_EMAIL_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_LOGIN:
		return http.StatusBadRequest
	case ERR_REGISTER:
		return http.StatusInternalServerError
	case ERR_PHONE_NUMBER_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PASSWORD_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_BCRYPT_PASSWORD:
		return http.StatusInternalServerError
	case ERR_TOKEN:
		return http.StatusInternalServerError
	case ERR_ROLE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_CREATE_ROLE:
		return http.StatusInternalServerError
	case ERR_GET_DATA:
		return http.StatusInternalServerError
	case ERR_NOT_FOUND:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
