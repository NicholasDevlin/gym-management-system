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
		CheckOut: input.CheckOut,
	}
}

func ConvertDtoToModel(input AbsensiDto) *Absensi {
	return &Absensi{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:     input.UUID,
		Date:     input.Date,
		UserId:   input.User.Id,
		User:     *user.ConvertDtoToModel(input.User),
		CheckOut: input.CheckOut,
	}
}

func ConvertModelToDto(input Absensi) *AbsensiDto {
	return &AbsensiDto{
		Id:        input.ID,
		UUID:      input.UUID,
		Date:      input.Date,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UserUUID:  input.User.UUID,
		User:      *user.ConvertModelToDto(input.User),
		CheckOut:  input.CheckOut,
	}
}

func ConvertDtoToRes(input AbsensiDto) *AbsensiRes {
	return &AbsensiRes{
		UUID:     input.UUID,
		Date:     input.Date,
		UserUUID: input.UserUUID,
		CheckOut: input.CheckOut,
		User:     *user.ConvertDtoToRes(input.User),
	}
}
