package services

import (
	"fmt"
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/role"
	"gym/app/backend/utils/errors"
)

type IRoleService interface {
	CreateRole(input role.RoleReq) (role.RoleRes, error)
	GetAllRole(filter role.RoleReq) ([]role.RoleRes, error)
	GetRole(filter role.RoleReq) (role.RoleRes, error)
	UpdateRole(input role.RoleReq) (role.RoleRes, error)
	DeleteRole(id uint) (role.RoleRes, error)
}

type roleService struct {
	roleRepository repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) *roleService {
	return &roleService{roleRepository: repo}
}

func (r *roleService) CreateRole(input role.RoleReq) (role.RoleRes, error) {
	if input.Role == "" {
		return role.RoleRes{}, errors.ERR_ROLE_IS_EMPTY
	}
	res, err := r.roleRepository.CreateRole(*role.ConvertReqToDto(input))
	if err != nil {
		return role.RoleRes{}, errors.ERR_CREATE_ROLE
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

	if err != nil || filter.Id == 0 {
		return role.RoleRes{}, errors.ERR_NOT_FOUND
	}
	return *role.ConvertDtoToRes(res), nil
}

func (r *roleService) UpdateRole(input role.RoleReq) (role.RoleRes, error) {
	res, err := r.roleRepository.GetRole(*role.ConvertReqToDto(input))
	if err != nil {
		return role.RoleRes{}, errors.ERR_NOT_FOUND
	}
	res, err = r.roleRepository.UpdateRole(res, *role.ConvertReqToDto(input))
	if err != nil {
		return role.RoleRes{}, errors.ERR_UPDATE_ROLE
	}
	return *role.ConvertDtoToRes(res), nil
}

func (r *roleService) DeleteRole(id uint) (role.RoleRes, error) {
	res, err := r.roleRepository.GetRole(role.RoleDto{Id: id})
	if err != nil {
		return role.RoleRes{}, errors.ERR_NOT_FOUND
	}

	res, err = r.roleRepository.DeleteRole(fmt.Sprint(id))
	if err != nil {
		return role.RoleRes{}, errors.ERR_DELETE_ROLE
	}
	return *role.ConvertDtoToRes(res), nil
}
