package services

import (
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/role"
	"gym/app/backend/models/user"
	"gym/app/backend/utils/bcrypt"
	"gym/app/backend/utils/consts"
	"gym/app/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IUserService interface {
	RegisterUser(input user.UserReq) (user.UserRes, error)
	LoginUser(input user.UserReq) (user.UserRes, error)
	GetAllUser(filter user.UserReq, page, pageSize int) ([]user.UserRes, int, error)
	GetUser(filter user.UserReq) (user.UserRes, error)
	UpdateUser(input user.UserReq) (user.UserRes, error)
	DeleteUser(id uuid.UUID) (user.UserRes, error)
}

type userService struct {
	userRepository repositories.IUserRepository
	roleRepository repositories.IRoleRepository
}

func NewUserService(repo repositories.IUserRepository, roleRepo repositories.IRoleRepository) *userService {
	return &userService{
		userRepository: repo,
		roleRepository: roleRepo,
	}
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

	if input.Role.Role == "" {
		input.Role.Role = consts.USER
	}

	roleRes, err := u.roleRepository.GetRole(role.RoleDto{Role: input.Role.Role, Id: input.RoleId})
	if err != nil {
		return user.UserRes{}, errors.ERR_GET_DATA
	}

	input.RoleId = roleRes.Id
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

func (u *userService) LoginUser(data user.UserReq) (user.UserRes, error) {
	if data.Email == "" {
		return user.UserRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.Password == "" {
		return user.UserRes{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	res, err := u.userRepository.LoginUser(*user.ConvertReqToDto(data))
	if err != nil {
		return user.UserRes{}, err
	}
	return *user.ConvertDtoToRes(res), nil
}

func (u *userService) GetUser(filter user.UserReq) (user.UserRes, error) {
	res, err := u.userRepository.GetUser(*user.ConvertReqToDto(filter))
	if err != nil || (filter.Id == 0 && filter.UUID == uuid.Nil) {
		return user.UserRes{}, errors.ERR_NOT_FOUND
	}
	return *user.ConvertDtoToRes(res), nil
}

func (u *userService) UpdateUser(input user.UserReq) (user.UserRes, error) {
	res, err := u.userRepository.GetUser(*user.ConvertReqToDto(input))
	if err != nil {
		return user.UserRes{}, errors.ERR_NOT_FOUND
	}
	roleRes, err := u.roleRepository.GetRole(role.RoleDto{Role: input.Role.Role})
	if err != nil {
		return user.UserRes{}, errors.ERR_GET_DATA
	}

	input.RoleId = roleRes.Id
	res, err = u.userRepository.UpdateUser(res, *user.ConvertReqToDto(input))
	if err != nil {
		return user.UserRes{}, errors.ERR_UPDATE_USER
	}
	return *user.ConvertDtoToRes(res), nil
}

func (u *userService) DeleteUser(id uuid.UUID) (user.UserRes, error) {
	res, err := u.userRepository.GetUser(user.UserDto{UUID: id})
	if err != nil {
		return user.UserRes{}, errors.ERR_NOT_FOUND
	}

	res, err = u.userRepository.DeleteUser(id.String())
	if err != nil {
		return user.UserRes{}, errors.ERR_DELETE_USER
	}
	return *user.ConvertDtoToRes(res), nil
}
