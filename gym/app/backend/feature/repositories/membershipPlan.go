package repositories

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IMembershipPlanRepository interface {
	CreateMembershipPlan(input membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error)
	GetAllMembershipPlan(filter membershipplan.MembershipPlanDto) ([]membershipplan.MembershipPlanDto, error)
	GetMembershipPlan(filter membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error)
	UpdateMembershipPlan(data, input membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error)
	DeleteMembershipPlan(id string) (membershipplan.MembershipPlanDto, error)
}

type membershipPlanRepository struct {
	db *gorm.DB
}

func NewMembershipPlanRepository(db *gorm.DB) *membershipPlanRepository {
	return &membershipPlanRepository{db}
}

func (mp *membershipPlanRepository) CreateMembershipPlan(input membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error) {
	dataMembershipPlan := membershipplan.ConvertDtoToModel(input)
	dataMembershipPlan.UUID = uuid.NewV4()
	err := mp.db.Create(&dataMembershipPlan).Error
	if err != nil {
		return membershipplan.MembershipPlanDto{}, err
	}
	return *membershipplan.ConvertModelToDto(*dataMembershipPlan), nil
}

func (mp *membershipPlanRepository) GetAllMembershipPlan(filter membershipplan.MembershipPlanDto) ([]membershipplan.MembershipPlanDto, error) {
	var allPlan []membershipplan.MembershipPlan
	var resAllPlan []membershipplan.MembershipPlanDto

	query := mp.db.Model(&membershipplan.MembershipPlan{})
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	err := query.Find(&allPlan).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allPlan); i++ {
		membershipPlan := membershipplan.ConvertModelToDto(allPlan[i])
		resAllPlan = append(resAllPlan, *membershipPlan)
	}
	return resAllPlan, nil
}

func (mp *membershipPlanRepository) GetMembershipPlan(filter membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error) {
	var model membershipplan.MembershipPlan
	query := mp.db.Model(&membershipplan.MembershipPlan{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.First(&model).Error
	if err != nil {
		return membershipplan.MembershipPlanDto{}, err
	}
	return *membershipplan.ConvertModelToDto(model), nil
}

func (mp *membershipPlanRepository) UpdateMembershipPlan(data, input membershipplan.MembershipPlanDto) (membershipplan.MembershipPlanDto, error) {
	membershipPlanData := *membershipplan.ConvertDtoToModel(data)

	if input.Name != "" {
		membershipPlanData.Name = input.Name
	}
	if input.Description != "" {
		membershipPlanData.Description = input.Description
	}
	if input.Duration != 0 {
		membershipPlanData.Duration = input.Duration
	}
	if input.Price != 0 {
		membershipPlanData.Price = input.Price
	}

	if err := mp.db.Save(&membershipPlanData).Error; err != nil {
		return membershipplan.MembershipPlanDto{}, err
	}
	return *membershipplan.ConvertModelToDto(membershipPlanData), nil
}

func (mp *membershipPlanRepository) DeleteMembershipPlan(id string) (membershipplan.MembershipPlanDto, error) {
	membershipPlanData := membershipplan.MembershipPlan{}

	err := mp.db.Delete(&membershipPlanData, "uuid = ?", id).Error
	if err != nil {
		return membershipplan.MembershipPlanDto{}, err
	}

	return *membershipplan.ConvertModelToDto(membershipPlanData), nil
}
