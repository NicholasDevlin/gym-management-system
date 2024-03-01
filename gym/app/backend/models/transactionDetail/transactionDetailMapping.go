package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func ConvertReqToDto(input TransactionDetailReq) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:             input.Id,
		UUID:           input.UUID,
		TransactionId:  input.TransactionId,
		Quantity:       input.Quantity,
		UserUUID:       input.UserUUID,
		User:           *user.ConvertReqToDto(input.User),
		MembershipPlan: *membershipplan.ConvertReqToDto(input.MembershipPlan),
	}
}

func ConvertDtoToModel(input TransactionDetailDto) *TransactionDetail {
	return &TransactionDetail{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:             input.UUID,
		UserId:           input.UserId,
		TransactionId:    input.TransactionId,
		MembershipPlanId: input.MembershipPlanId,
		Quantity:         input.Quantity,
		User:             *user.ConvertDtoToModel(input.User),
		MembershipPlan:   *membershipplan.ConvertDtoToModel(input.MembershipPlan),
	}
}

func ConvertModelToDto(input TransactionDetail) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:                 input.ID,
		CreatedAt:          input.CreatedAt,
		UpdatedAt:          input.UpdatedAt,
		UUID:               input.UUID,
		TransactionId:      input.TransactionId,
		MembershipPlanId:   input.MembershipPlanId,
		MembershipPlanUUID: input.MembershipPlan.UUID,
		Quantity:           input.Quantity,
		UserUUID:           input.User.UUID,
		UserId:             input.UserId,
		User:               *user.ConvertModelToDto(input.User),
		MembershipPlan:     *membershipplan.ConvertModelToDto(input.MembershipPlan),
	}
}

func ConvertDtoToRes(input TransactionDetailDto) *TransactionDetailRes {
	return &TransactionDetailRes{
		UUID:               input.UUID,
		Quantity:           input.Quantity,
		MembershipPlanUUID: input.MembershipPlanUUID,
		Subtotal:           int64(input.Quantity) * input.MembershipPlan.Price,
		UserUUID:           input.UserUUID,
		User:               *user.ConvertDtoToRes(input.User),
		MembershipPlan:     *membershipplan.ConvertDtoToRes(input.MembershipPlan),
	}
}
