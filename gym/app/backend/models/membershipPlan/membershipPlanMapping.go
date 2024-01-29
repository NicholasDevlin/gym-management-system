package membershipplan

import "gorm.io/gorm"

func ConvertReqToDto(input MembershipPlanReq) *MembershipPlanDto {
	return &MembershipPlanDto{
		Id:             input.Id,
		UUID:           input.UUID,
		Name: input.Name,
		Description: input.Description,
		Duration: input.Duration,
		Price: input.Price,
	}
}

func ConvertDtoToModel(input MembershipPlanDto) *MembershipPlan {
	return &MembershipPlan{
		Model: gorm.Model{
			ID: input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:           input.UUID,
		Name: input.Name,
		Description: input.Description,
		Duration: input.Duration,
		Price: input.Price,
	}
}

func ConvertModelToDto(input MembershipPlan) *MembershipPlanDto {
	return &MembershipPlanDto{
		Id:             input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UUID:           input.UUID,
		Name: input.Name,
		Description: input.Description,
		Duration: input.Duration,
		Price: input.Price,
	}
}

func ConvertDtoToRes(input MembershipPlanDto) *MembershipPlanRes {
	return &MembershipPlanRes{
		Id:             input.Id,
		UUID:           input.UUID,
		Name: input.Name,
		Description: input.Description,
		Duration: input.Duration,
		Price: input.Price,
	}
}