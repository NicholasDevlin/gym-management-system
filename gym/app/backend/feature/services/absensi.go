package services

import (
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/absensi"
	"gym/app/backend/utils/errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type IAbsensiService interface {
	CreateAbsensi(input absensi.AbsensiReq) (absensi.AbsensiRes, error)
	// GetAllRole(filter absensi.AbsensiReq) ([]absensi.AbsensiRes, error)
	// GetRole(filter absensi.AbsensiReq) (absensi.AbsensiRes, error)
	// UpdateRole(input absensi.AbsensiReq) (absensi.AbsensiRes, error)
	// DeleteRole(id uint) (absensi.AbsensiRes, error)
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

func (a *absensiService) CreateAbsensi(input absensi.AbsensiReq) (absensi.AbsensiRes, error) {
	if input.UserUUID == uuid.Nil {
		return absensi.AbsensiRes{}, errors.ERR_USER_NOT_FOUND
	}
	if input.Date.IsZero() {
		input.Date = time.Now()
	}
	var err error
	dto := absensi.ConvertReqToDto(input)
	dto.User, err = a.userRepository.GetUser(dto.User)
	if err != nil {
		return absensi.AbsensiRes{}, errors.ERR_USER_NOT_FOUND
	}
	
	err = a.absensiRepository.CreateAbsensi(dto)
	if err != nil {
		return absensi.AbsensiRes{}, errors.ERR_CREATE_ABSENSI
	}
	return *absensi.ConvertDtoToRes(*dto), nil
}

// func (r *absensiService) GetAllRole(filter absensi.AbsensiReq) ([]absensi.AbsensiRes, error) {
// 	res, err := r.roleRepository.GetAllRole(*absensi.ConvertReqToDto(filter))
// 	if err != nil {
// 		return nil, errors.ERR_GET_DATA
// 	}
// 	var resRole []absensi.AbsensiRes
// 	for i := 0; i < len(res); i++ {
// 		roleVm := absensi.ConvertDtoToRes(res[i])
// 		resRole = append(resRole, *roleVm)
// 	}
// 	return resRole, nil
// }

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

// func (r *absensiService) DeleteRole(id uint) (absensi.AbsensiRes, error) {
// 	res, err := r.roleRepository.GetRole(absensi.RoleDto{Id: id})
// 	if err != nil {
// 		return absensi.AbsensiRes{}, errors.ERR_NOT_FOUND
// 	}

// 	res, err = r.roleRepository.DeleteRole(fmt.Sprint(id))
// 	if err != nil {
// 		return absensi.AbsensiRes{}, errors.ERR_DELETE_ROLE
// 	}
// 	return *absensi.ConvertDtoToRes(res), nil
// }
