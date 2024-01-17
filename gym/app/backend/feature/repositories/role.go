package repositories

import (
	"gym/app/backend/entity/role"

	"gorm.io/gorm"
)

type IRoleRepository interface {
	CreateRole(input role.RoleDto) (role.RoleDto, error)
	GetAllRole(filter role.RoleDto) ([]role.RoleDto, error)
	GetRole(filter role.RoleDto) (role.RoleDto, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) CreateRole(input role.RoleDto) (role.RoleDto, error) {
	dataRole := role.ConvertDtoToModel(input)
	err := r.db.Create(&dataRole).Error
	if err != nil {
		return role.RoleDto{}, err
	}
	return *role.ConvertModelToDto(*dataRole), nil
}

func (r *roleRepository) GetAllRole(filter role.RoleDto) ([]role.RoleDto, error) {
	var allRole []role.Role
	var resAllRole []role.RoleDto

	query := r.db.Model(&role.Role{})
	if filter.Role != "" {
		query = query.Where("role LIKE ?", "%"+filter.Role+"%")
	}

	err := query.Find(&allRole).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allRole); i++ {
		roleVm := role.ConvertModelToDto(allRole[i])
		resAllRole = append(resAllRole, *roleVm)
	}
	return resAllRole, nil
}

func (r *roleRepository) GetRole(filter role.RoleDto) (role.RoleDto, error) {
	var model role.Role
	query := r.db.Model(&role.Role{})
	if filter.Role != "" {
		query = query.Where("role LIKE ?", "%"+filter.Role+"%")
	}
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}

	err := query.First(&model).Error
	if err != nil {
		return role.RoleDto{}, err
	}
	return *role.ConvertModelToDto(model), nil
}
