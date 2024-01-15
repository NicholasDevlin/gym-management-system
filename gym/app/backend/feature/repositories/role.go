package repositories

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/utils/errors"

	"gorm.io/gorm"
)

type IRoleRepository interface {
	CreateRole(input role.RoleDto) (role.RoleDto, error)
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
	err = r.db.Preload("Role").First(&dataRole, "id = ?", dataRole.ID).Error
	if err != nil {
		return role.RoleDto{}, errors.ERR_LOGIN
	}
	return *role.ConvertModelToDto(*dataRole), nil
}

// func (r *roleRepository) GetAllRole(filter role.RoleDto) ([]role.RoleDto, error){

// 	// return
// }
