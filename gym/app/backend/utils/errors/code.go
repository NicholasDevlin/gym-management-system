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
	case ERR_EMAIL_NOT_FOUND:
		return http.StatusNotFound
	case ERR_WRONG_PASSWORD:
		return http.StatusBadRequest
	case ERR_UPDATE_USER:
		return http.StatusInternalServerError
	case ERR_DELETE_USER:
		return http.StatusInternalServerError
	case ERR_UPDATE_ROLE:
		return http.StatusInternalServerError
	case ERR_DELETE_ROLE:
		return http.StatusInternalServerError
	case ERR_MEMBERSHIP_NAME_EMPTY:
		return http.StatusBadRequest
	case ERR_MEMBERSHIP_DURATION_EMPTY:
		return http.StatusBadRequest
	case ERR_MEMBERSHIP_PRICE_EMPTY:
		return http.StatusBadRequest
	case ERR_DELETE_MEMBERSHIP_PLAN:
		return http.StatusInternalServerError
	case ERR_UNAUTHORIZE:
		return http.StatusUnauthorized
	case ERR_CREATE_MEMBERSHIP_PLAN:
		return http.StatusInternalServerError
	case ERR_CREATE_PURCHASE:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
