package services

import (
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/absensi"
	"gym/app/backend/models/user"
	"gym/app/backend/utils/errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type IAbsensiService interface {
	SaveAbsensi(input absensi.AbsensiReq) (absensi.AbsensiRes, error)
	GetAllAbsensi(filter absensi.AbsensiFilter) ([]absensi.AbsensiRes, error)
	// GetRole(filter absensi.AbsensiReq) (absensi.AbsensiRes, error)
	// UpdateRole(input absensi.AbsensiReq) (absensi.AbsensiRes, error)
	DeleteAbsensi(id uuid.UUID) error
}

type absensiService struct {
	absensiRepository repositories.IAbsensiRepository
	userRepository    repositories.IUserRepository
}

func NewAbsensiService(repo repositories.IAbsensiRepository, userRepo repositories.IUserRepository) IAbsensiService {
	return &absensiService{
		absensiRepository: repo,
		userRepository:    userRepo,
	}
}

func (a *absensiService) SaveAbsensi(input absensi.AbsensiReq) (absensi.AbsensiRes, error) {
	if input.UserUUID == uuid.Nil {
		return absensi.AbsensiRes{}, errors.ERR_USER_NOT_FOUND
	}
	if input.Date.IsZero() {
		input.Date = time.Now()
	}
	
	var err error
	dto := absensi.ConvertReqToDto(input)

	var existing absensi.AbsensiDto
	if input.UUID != uuid.Nil {
		existing, err = a.absensiRepository.GetAbsensi(*dto)
		if err != nil {
			return absensi.AbsensiRes{}, errors.ERR_CREATE_ABSENSI
		}
	}
	existing.UserUUID = dto.UserUUID
	existing.CheckOut = dto.CheckOut
	existing.Date = dto.Date
	existing.User, err = a.userRepository.GetUser(user.UserDto{UUID: input.UserUUID})
	if err != nil {
		return absensi.AbsensiRes{}, errors.ERR_USER_NOT_FOUND
	}

	err = a.absensiRepository.SaveAbsensi(&existing)
	if err != nil {
		return absensi.AbsensiRes{}, errors.ERR_CREATE_ABSENSI
	}
	return *absensi.ConvertDtoToRes(*dto), nil
}

func (a *absensiService) GetAllAbsensi(filter absensi.AbsensiFilter) ([]absensi.AbsensiRes, error) {
	res, err := a.absensiRepository.GetAllAbsensi(filter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	var resRole []absensi.AbsensiRes
	for i := 0; i < len(res); i++ {
		roleVm := absensi.ConvertDtoToRes(res[i])
		resRole = append(resRole, *roleVm)
	}
	return resRole, nil
}

// func (r *absensiService) GetRole(filter absensi.AbsensiReq) (absensi.AbsensiRes, error) {
// 	res, err := r.roleRepository.GetRole(*absensi.ConvertReqToDto(filter))

// 	if err != nil || filter.Id == 0 {
// 		return absensi.AbsensiRes{}, errors.ERR_NOT_FOUND
// 	}
// 	return *absensi.ConvertDtoToRes(res), nil
// }

// func (r *absensiService) UpdateRole(input absensi.AbsensiReq) (absensi.AbsensiRes, error) {
// 	res, err := r.roleRepository.GetRole(*absensi.ConvertReqToDto(input))
// 	if err != nil {
// 		return absensi.AbsensiRes{}, errors.ERR_NOT_FOUND
// 	}
// 	res, err = r.roleRepository.UpdateRole(res, *absensi.ConvertReqToDto(input))
// 	if err != nil {
// 		return absensi.AbsensiRes{}, errors.ERR_UPDATE_ROLE
// 	}
// 	return *absensi.ConvertDtoToRes(res), nil
// }

func (a *absensiService) DeleteAbsensi(id uuid.UUID) error {
	res, err := a.absensiRepository.GetAbsensi(absensi.AbsensiDto{UUID: id})
	if err != nil {
		return errors.ERR_NOT_FOUND
	}

	err = a.absensiRepository.DeleteAbsensi(res.Id)
	if err != nil {
		return errors.ERR_DELETE_ROLE
	}
	return nil
}
