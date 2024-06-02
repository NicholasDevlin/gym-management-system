package repositories

import (
	"gym/app/backend/models/absensi"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IAbsensiRepository interface {
	CreateAbsensi(input *absensi.AbsensiDto) error
	// GetAllRole(filter absensi.AbsensiDto) ([]absensi.AbsensiDto, error)
	// GetRole(filter *absensi.AbsensiDto) error
	// UpdateRole(data, input *absensi.AbsensiDto) error
	// DeleteRole(id string) error
}

type absensiRepository struct {
	db *gorm.DB
}

func NewAbsensiRepository(db *gorm.DB) IAbsensiRepository {
	return &absensiRepository{db}
}

func (r *absensiRepository) CreateAbsensi(input *absensi.AbsensiDto) error {
	dataAbsensi := absensi.ConvertDtoToModel(*input)
	dataAbsensi.UUID = uuid.NewV4()
	err := r.db.Create(&dataAbsensi).Error
	if err != nil {
		return err
	}
	input = absensi.ConvertModelToDto(*dataAbsensi)
	return nil
}

// func (r *absensiRepository) GetAllRole(filter absensi.AbsensiDto) ([]absensi.AbsensiDto, error) {
// 	var allRole []absensi.Role
// 	var resAllRole []absensi.AbsensiDto

// 	query := r.db.Model(&absensi.Role{})
// 	if filter.Role != "" {
// 		query = query.Where("role LIKE ?", "%"+filter.Role+"%")
// 	}

// 	err := query.Find(&allRole).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i := 0; i < len(allRole); i++ {
// 		role := absensi.ConvertModelToDto(allRole[i])
// 		resAllRole = append(resAllRole, *role)
// 	}
// 	return resAllRole, nil
// }

// func (r *absensiRepository) GetRole(filter absensi.AbsensiDto) error {
// 	var model absensi.Role
// 	query := r.db.Model(&absensi.Role{})
// 	if filter.Id != 0 {
// 		query = query.Where("id = ?", filter.Id)
// 	}
// 	if filter.Role != "" {
// 		query = query.Where("role = ?", filter.Role)
// 	}

// 	err := query.First(&model).Error
// 	if err != nil {
// 		return absensi.AbsensiDto{}, err
// 	}
// 	return *absensi.ConvertModelToDto(model), nil
// }

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

// func (r *absensiRepository) DeleteRole(id string) error {
// 	roleData := absensi.Role{}

// 	err := r.db.Delete(&roleData, "id = ?", id).Error
// 	if err != nil {
// 		return absensi.AbsensiDto{}, err
// 	}

// 	return *absensi.ConvertModelToDto(roleData), nil
// }
