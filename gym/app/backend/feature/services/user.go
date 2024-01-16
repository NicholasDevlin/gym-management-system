package services

import (
	"gym/app/backend/entity/user"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/utils/bcrypt"
	"gym/app/backend/utils/errors"
)

type IUserService interface {
	RegisterUser(input user.UserReq) (user.UserRes, error)
	// LoginUser(input *request.User) (user.CreaRestors, error)
	GetAllUser(filter user.UserReq, page, pageSize int) ([]user.UserRes, int, error) 
	// GetUser(id string) (user.UserRes, error)
	// UpdateUser(id string, input request.User) (user.UserRes, error)
	// DeleteUser(id string) (user.UserRes, error)
	// CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	// GetNextPage(currentPage, allPages int) int
	// GetPrevPage(currentPage int) int
	// CountUsersByRole(roleId uint) (int, error)
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *userService {
	return &userService{userRepository: repo}
}

func (u *userService) RegisterUser(input user.UserReq) (user.UserRes, error) {
	if input.Email == "" {
		return user.UserRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if input.PhoneNumber == "" {
		return user.UserRes{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	if input.Password == "" {
		return user.UserRes{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	hashPass, err := bcrypt.HashPassword(input.Password)
	if err != nil {
		return user.UserRes{}, errors.ERR_BCRYPT_PASSWORD
	}

	input.Password = hashPass
	res, err := u.userRepository.RegisterUser(*user.ConvertReqToDto(input))
	if err != nil {
		return user.UserRes{}, err
	}
	return *user.ConvertDtoToRes(res), nil
}

func (u *userService) GetAllUser(filter user.UserReq, page, pageSize int) ([]user.UserRes, int, error) {
	res, totalRecord, err := u.userRepository.GetAllUser(*user.ConvertReqToDto(filter), page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	var resUser []user.UserRes
	for i := 0; i < len(res); i++ {
		roleVm := user.ConvertDtoToRes(res[i])
		resUser = append(resUser, *roleVm)
	}
	return resUser, totalRecord, nil
}
