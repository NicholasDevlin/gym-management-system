package errors

import "errors"

var (
	ERR_EMAIL_IS_EMPTY        = errors.New("Email is empty")
	ERR_LOGIN                 = errors.New("Failed to signing")
	ERR_REGISTER              = errors.New("Failed to register")
	ERR_CREATE_ROLE           = errors.New("Failed to create new Role")
	ERR_PASSWORD_IS_EMPTY     = errors.New("Password is empty")
	ERR_PHONE_NUMBER_IS_EMPTY = errors.New("Phone number is empty")
	ERR_ROLE_IS_EMPTY         = errors.New("Role is empty")
	ERR_BCRYPT_PASSWORD       = errors.New("Failed to bcrypt password")
	ERR_TOKEN                 = errors.New("Failed to create new token")
	ERR_GET_DATA              = errors.New("Failed to get data from database")
	ERR_NOT_FOUND             = errors.New("Data not Found")
)
