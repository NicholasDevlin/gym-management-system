package transactiondetail

import "gorm.io/gorm"

func ConvertReqToDto(input TransactionDetailReq) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:              input.Id,
		UUID:            input.UUID,
		TransactionId: input.TransactionId,
		MembershipPlanId: input.MembershipPlanId,
		Quantity: input.Quantity,
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
		Quantity: input.Quantity,
	}
}

func ConvertDtoToRes(input TransactionDetailDto) *TransactionDetailRes {
	return &TransactionDetailRes{
		UUID:            input.UUID,
		Quantity: input.Quantity,
	}
}
