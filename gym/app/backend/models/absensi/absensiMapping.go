package absensi

import (
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func ConvertReqToDto(input AbsensiReq) *AbsensiDto {
	return &AbsensiDto{
		UUID:     input.UUID,
		Date:     input.Date,
		UserUUID: input.UserUUID,
	}
}

func ConvertDtoToModel(input AbsensiDto) *Absensi {
	return &Absensi{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		Date:   input.Date,
		UserId: input.UserId,
		User:   *user.ConvertDtoToModel(input.User),
	}
}

func ConvertModelToDto(input Absensi) *AbsensiDto {
	return &AbsensiDto{
		Id:        input.ID,
		Date:      input.Date,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UserUUID:  input.User.UUID,
		User:      *user.ConvertModelToDto(input.User),
	}
}

func ConvertDtoToRes(input AbsensiDto) *AbsensiRes {
	return &AbsensiRes{
		Date:     input.Date,
		UserUUID: input.UserUUID,
		User:     *user.ConvertDtoToRes(input.User),
	}
}
