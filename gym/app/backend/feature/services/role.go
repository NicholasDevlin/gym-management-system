package services

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/utils/errors"
)

type IRoleService interface {
	CreateRole(input role.RoleReq) (role.RoleRes, error)
	GetAllRole(filter role.RoleReq) ([]role.RoleRes, error)
	GetRole(filter role.RoleReq) (role.RoleRes, error)
}

type roleService struct {
	roleRepository repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) *roleService {
	return &roleService{roleRepository: repo}
}

func (r *roleService) CreateRole(input role.RoleReq) (role.RoleRes, error) {
	if input.Role == "" {
		return role.RoleRes{}, errors.ERR_BCRYPT_PASSWORD
	}
	res, err := r.roleRepository.CreateRole(*role.ConvertReqToDto(input))
	if err != nil {
		return role.RoleRes{}, errors.ERR_BCRYPT_PASSWORD
	}
	return *role.ConvertDtoToRes(res), nil
}

func (r *roleService) GetAllRole(filter role.RoleReq) ([]role.RoleRes, error) {
	res, err := r.roleRepository.GetAllRole(*role.ConvertReqToDto(filter))
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	var resRole []role.RoleRes
	for i := 0; i < len(res); i++ {
		roleVm := role.ConvertDtoToRes(res[i])
		resRole = append(resRole, *roleVm)
	}
	return resRole, nil
}

func (r *roleService) GetRole(filter role.RoleReq) (role.RoleRes, error) {
	res, err := r.roleRepository.GetRole(*role.ConvertReqToDto(filter))

	if err != nil {
		return role.RoleRes{}, errors.ERR_NOT_FOUND
	}
	return *role.ConvertDtoToRes(res), nil
}
