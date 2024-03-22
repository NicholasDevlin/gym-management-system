package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"

	"gorm.io/gorm"
)

func ConvertReqToDto(input TransactionDetailReq) *TransactionDetailDto {
	return &TransactionDetailDto{
		UUID:           input.UUID,
		TransactionId:  input.TransactionId,
		Quantity:       input.Quantity,
		MembershipPlan: *membershipplan.ConvertReqToDto(input.MembershipPlan),
		TransactionMemberDetail: *ConvertReqToDtos(input.TransactionMemberDetail),
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
		TransactionId:    input.TransactionId,
		MembershipPlanId: input.MembershipPlanId,
		Quantity:         input.Quantity,
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
		MembershipPlan:     *membershipplan.ConvertModelToDto(input.MembershipPlan),
		TransactionMemberDetail: *ConvertModelToDtos(input.TransactionMemberDetail),
	}
}

func ConvertDtoToRes(input TransactionDetailDto) *TransactionDetailRes {
	return &TransactionDetailRes{
		UUID:               input.UUID,
		Quantity:           input.Quantity,
		MembershipPlanUUID: input.MembershipPlanUUID,
		Subtotal:           int64(input.Quantity) * input.MembershipPlan.Price,
		MembershipPlan:     *membershipplan.ConvertDtoToRes(input.MembershipPlan),
		TransactionMemberDetail: *ConvertDtosToRes(input.TransactionMemberDetail),
	}
}

func ConvertDtosToRes(input []transactionmemberdetail.TransactionMemberDetailDto) (*[]transactionmemberdetail.TransactionMemberDetailRes) {
	var result []transactionmemberdetail.TransactionMemberDetailRes
	for i := range input {
		res := *transactionmemberdetail.ConvertDtoToRes(input[i])
		result = append(result, res)
	}
	return &result
}

func ConvertModelToDtos(input []transactionmemberdetail.TransactionMemberDetail) *[]transactionmemberdetail.TransactionMemberDetailDto {
	var result []transactionmemberdetail.TransactionMemberDetailDto
	for i := range input {
		res := *transactionmemberdetail.ConvertModelToDto(input[i])
		result = append(result, res)
	}
	return &result
}

func ConvertReqToDtos(input []transactionmemberdetail.TransactionMemberDetailReq) *[]transactionmemberdetail.TransactionMemberDetailDto {
	var result []transactionmemberdetail.TransactionMemberDetailDto
	for i := range input {
		res := *transactionmemberdetail.ConvertReqToDto(input[i])
		result = append(result, res)
	}
	return &result
}
