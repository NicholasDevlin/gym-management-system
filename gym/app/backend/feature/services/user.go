package services

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/utils/bcrypt"
	"gym/app/backend/utils/errors"
)

type IUserService interface {
	RegisterUser(data user.UserReq) (user.UserRes, error)
	// LoginUser(data *request.User) (user.CreaRestors, error)
	// GetAllUser(nameFilter string, page, pageSize int) ([]user.UserRes, map[string]int, error)
	// GetUser(id string) (user.UserRes, error)
	// UpdateUser(id string, input request.User) (user.UserRes, error)
	// DeleteUser(id string) (user.UserRes, error)
	// CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	// GetNextPage(currentPage, allPages int) int
	// GetPrevPage(currentPage int) int
	// CountUsersByRole(roleId uint) (int, error)
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) RegisterUser(data user.UserReq) (user.UserRes, error) {
	if data.Email == "" {
		return user.UserRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.PhoneNumber == "" {
		return user.UserRes{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	if data.Password == "" {
		return user.UserRes{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	hashPass, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return user.UserRes{}, errors.ERR_BCRYPT_PASSWORD
	}

	data.Password = hashPass
	res, err := u.UserRepo.RegisterUser(*user.ConvertReqToDto(data))
	if err != nil {
		return user.UserRes{}, err
	}
	return *user.ConvertDtoToRes(res), nil
}
