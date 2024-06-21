package repositories

import (
	"gym/app/backend/models/absensi"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IAbsensiRepository interface {
	SaveAbsensi(input *absensi.AbsensiDto) error
	GetAllAbsensi(filter absensi.AbsensiFilter) ([]absensi.AbsensiDto, error)
	GetAbsensi(filter absensi.AbsensiDto) (absensi.AbsensiDto, error) 
	// UpdateRole(data, input *absensi.AbsensiDto) error
	DeleteAbsensi(id uint) error
}

type absensiRepository struct {
	db *gorm.DB
}

func NewAbsensiRepository(db *gorm.DB) IAbsensiRepository {
	return &absensiRepository{db}
}

func (a *absensiRepository) SaveAbsensi(input *absensi.AbsensiDto) error {
	dataAbsensi := absensi.ConvertDtoToModel(*input)
	if dataAbsensi.UUID == uuid.Nil {
		dataAbsensi.UUID = uuid.NewV4()
	}

	err := a.db.Save(&dataAbsensi).Error
	if err != nil {
		return err
	}
	input = absensi.ConvertModelToDto(*dataAbsensi)
	return nil
}

func (a *absensiRepository) GetAllAbsensi(filter absensi.AbsensiFilter) ([]absensi.AbsensiDto, error) {
	var allAbsensi []absensi.Absensi
	var resAllAbsensi []absensi.AbsensiDto

	query := a.db.Preload("User").Model(&absensi.Absensi{})
	if !filter.DateFrom.IsZero() {
		dateOnly := time.Date(filter.DateFrom.Year(), filter.DateFrom.Month(), filter.DateFrom.Day(), 0, 0, 0, 0, filter.DateFrom.Location())
    query = query.Where("date >= ?", dateOnly)
	}
	if !filter.DateTo.IsZero() {
		dateOnly := time.Date(filter.DateFrom.Year(), filter.DateFrom.Month(), filter.DateFrom.Day(), 0, 0, 0, 0, filter.DateFrom.Location())
		dateOnly = dateOnly.Add(24 * time.Hour)
		query = query.Where("date <= ?", dateOnly)
	}

	query = query.Order("date desc")

	err := query.Find(&allAbsensi).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allAbsensi); i++ {
		role := absensi.ConvertModelToDto(allAbsensi[i])
		resAllAbsensi = append(resAllAbsensi, *role)
	}
	return resAllAbsensi, nil
}

func (a *absensiRepository) GetAbsensi(filter absensi.AbsensiDto) (absensi.AbsensiDto, error) {
	var model absensi.Absensi
	query := a.db.Model(&absensi.Absensi{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.First(&model).Error
	if err != nil {
		return absensi.AbsensiDto{}, err
	}
	return *absensi.ConvertModelToDto(model), nil
}

// func (r *absensiRepository) UpdateRole(data, input absensi.AbsensiDto) error {
// 	roleData := *absensi.ConvertDtoToModel(data)

// 	if input.Role != "" {
// 		roleData.Role = input.Role
// 	}

// 	if err := r.db.Save(&roleData).Error; err != nil {
// 		return absensi.AbsensiDto{}, err
// 	}
// 	return *absensi.ConvertModelToDto(roleData), nil
// }

func (r *absensiRepository) DeleteAbsensi(id uint) error {
	err := r.db.Delete(&absensi.Absensi{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
