package user

import (
	"gym/app/backend/entity/role"

	"gorm.io/gorm"
)

func ConvertReqToDto(input UserReq) *UserDto {
	return &UserDto{
		Id:             input.Id,
		UUID:           input.UUID,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
		RoleId: input.RoleId,
		Role: role.RoleDto{
			Role: input.Role.Role,
		},
	}
}

func ConvertDtoToModel(input UserDto) *User {
	return &User{
		Model: gorm.Model{
			ID: input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:           input.UUID,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
		RoleId: input.RoleId,
		Role: role.Role{
			Role: input.Role.Role,
		},
	}
}

func ConvertModelToDto(input User) *UserDto {
	return &UserDto{
		Id:             input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UUID:           input.UUID,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
		RoleId: input.RoleId,
		Role: role.RoleDto{
			Role: input.Role.Role,
		},
	}
}

func ConvertDtoToRes(input UserDto) *UserRes {
	return &UserRes{
		Id:             input.Id,
		UUID:           input.UUID,
		DisplayName:    input.DisplayName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Gender:         input.Gender,
		BirthDate:      input.BirthDate,
		GoogleID:       input.GoogleID,
		ProfilePicture: input.ProfilePicture,
		IsGoogleUser:   input.IsGoogleUser,
		RoleId: input.RoleId,
		Role: role.RoleRes{
			Role: input.Role.Role,
		},
	}
}
