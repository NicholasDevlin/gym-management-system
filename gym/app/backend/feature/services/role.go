package services

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/utils/errors"
)

type IRoleService interface {
	CreateRole(input role.RoleReq) (role.RoleRes, error)
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
