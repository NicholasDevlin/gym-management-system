package role

import "gorm.io/gorm"

func ConvertReqToDto(input RoleReq) *RoleDto {
	return &RoleDto{
		Id:   input.Id,
		Role: input.Role,
	}
}

func ConvertDtoToModel(input RoleDto) *Role {
	return &Role{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		Role: input.Role,
	}
}

func ConvertModelToDto(input Role) *RoleDto {
	return &RoleDto{
		Id:        input.ID,
		Role:      input.Role,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ConvertDtoToRes(input RoleDto) *RoleRes {
	return &RoleRes{
		Role: input.Role,
	}
}
