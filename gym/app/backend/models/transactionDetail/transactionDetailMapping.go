package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func ConvertReqToDto(input TransactionDetailReq) *TransactionDetailDto {
	return &TransactionDetailDto{
		UUID:                    input.UUID,
		TransactionId:           input.TransactionId,
		Quantity:                input.Quantity,
		Price:                   input.Price,
		MembershipPlanUUID:      input.MembershipPlanUUID,
		Deleted:                 input.Deleted,
		MembershipPlan:          *membershipplan.ConvertReqToDto(input.MembershipPlan),
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
		IsDel: soft_delete.DeletedAt(input.Deleted),
		UUID:                    input.UUID,
		TransactionId:           input.TransactionId,
		MembershipPlanId:        input.MembershipPlanId,
		Quantity:                input.Quantity,
		Price:                   input.Price,
		MembershipPlan:          *membershipplan.ConvertDtoToModel(input.MembershipPlan),
		TransactionMemberDetail: *ConvertDtosToModel(input.TransactionMemberDetail),
	}
}

func ConvertModelToDto(input TransactionDetail) *TransactionDetailDto {
	return &TransactionDetailDto{
		Id:                      input.ID,
		CreatedAt:               input.CreatedAt,
		UpdatedAt:               input.UpdatedAt,
		UUID:                    input.UUID,
		TransactionId:           input.TransactionId,
		MembershipPlanId:        input.MembershipPlanId,
		MembershipPlanUUID:      input.MembershipPlan.UUID,
		Quantity:                input.Quantity,
		Price:                   input.Price,
		MembershipPlan:          *membershipplan.ConvertModelToDto(input.MembershipPlan),
		TransactionMemberDetail: *ConvertModelToDtos(input.TransactionMemberDetail),
	}
}

func ConvertDtoToRes(input TransactionDetailDto) *TransactionDetailRes {
	transactionMember, subtotal := ConvertDtosToRes((input.Quantity * input.Price), input.TransactionMemberDetail)
	return &TransactionDetailRes{
		UUID:                    input.UUID,
		Quantity:                input.Quantity,
		MembershipPlanUUID:      input.MembershipPlanUUID,
		Price:                   input.Price,
		Subtotal:                subtotal,
		MembershipPlan:          *membershipplan.ConvertDtoToRes(input.MembershipPlan),
		TransactionMemberDetail: *transactionMember,
	}
}

func ConvertDtosToRes(subtotal int, input []transactionmemberdetail.TransactionMemberDetailDto) (*[]transactionmemberdetail.TransactionMemberDetailRes, int) {
	var result []transactionmemberdetail.TransactionMemberDetailRes
	for i := range input {
		res := *transactionmemberdetail.ConvertDtoToRes(input[i])
		result = append(result, res)
		subtotal += int(res.AdditionalPrice)
	}
	return &result, subtotal
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

func ConvertDtosToModel(input []transactionmemberdetail.TransactionMemberDetailDto) *[]transactionmemberdetail.TransactionMemberDetail {
	var result []transactionmemberdetail.TransactionMemberDetail
	for i := range input {
		res := *transactionmemberdetail.ConvertDtoToModel(input[i])
		result = append(result, res)
	}
	return &result
}
