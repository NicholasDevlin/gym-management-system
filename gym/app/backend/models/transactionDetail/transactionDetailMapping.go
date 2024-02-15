package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	"gorm.io/gorm"
)

func ConvertReqToDto(input TransactionDetailReq) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:              input.Id,
		UUID:            input.UUID,
		TransactionId: input.TransactionId,
		Quantity: input.Quantity,
		MembershipPlan:  *membershipplan.ConvertReqToDto(input.MembershipPlan),
	}
}

func ConvertDtoToModel(input TransactionDetailDto) *TransactionDetail {
	return &TransactionDetail{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:            input.UUID,
		TransactionId: input.TransactionId,
		MembershipPlanId: input.MembershipPlanId,
		Quantity: input.Quantity,
		MembershipPlan: *membershipplan.ConvertDtoToModel(input.MembershipPlan),
	}
}

func ConvertModelToDto(input TransactionDetail) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:              input.ID,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
		UUID:            input.UUID,
		TransactionId: input.TransactionId,
		MembershipPlanId: input.MembershipPlanId,
		MembershipPlanUUID: input.MembershipPlan.UUID,
		Quantity: input.Quantity,
		MembershipPlan: *membershipplan.ConvertModelToDto(input.MembershipPlan),
	}
}

func ConvertDtoToRes(input TransactionDetailDto) *TransactionDetailRes {
	return &TransactionDetailRes{
		UUID:            input.UUID,
		Quantity: input.Quantity,
		MembershipPlanUUID: input.MembershipPlanUUID,
		Subtotal: int64(input.Quantity) * input.MembershipPlan.Price,
		MembershipPlan: *membershipplan.ConvertDtoToRes(input.MembershipPlan),
	}
}
