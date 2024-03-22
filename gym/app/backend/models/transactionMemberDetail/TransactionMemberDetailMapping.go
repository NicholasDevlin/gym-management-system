package transactionmemberdetail

import (
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func ConvertReqToDto(input TransactionMemberDetailReq) *TransactionMemberDetailDto {
	return &TransactionMemberDetailDto{
		UUID:                  input.UUID,
		UserUUID:              input.UserUUID,
		User:                  *user.ConvertReqToDto(input.User),
		TransactionDetailUUID: input.TransactionDetailUUID,
	}
}

func ConvertDtoToModel(input TransactionMemberDetailDto) *TransactionMemberDetail {
	return &TransactionMemberDetail{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:                input.UUID,
		UserId:              input.UserId,
		User:                *user.ConvertDtoToModel(input.User),
		TransactionDetailId: input.TransactionDetailId,
	}
}

func ConvertModelToDto(input TransactionMemberDetail) *TransactionMemberDetailDto {
	return &TransactionMemberDetailDto{
		Id:                  input.ID,
		CreatedAt:           input.CreatedAt,
		UpdatedAt:           input.UpdatedAt,
		UUID:                input.UUID,
		UserId:              input.UserId,
		User:                *user.ConvertModelToDto(input.User),
		UserUUID:            input.User.UUID,
		TransactionDetailId: input.TransactionDetailId,
	}
}

func ConvertDtoToRes(input TransactionMemberDetailDto) *TransactionMemberDetailRes {
	return &TransactionMemberDetailRes{
		UUID:                  input.UUID,
		UserUUID:              input.UserUUID,
		User:                  *user.ConvertDtoToRes(input.User),
		TransactionDetailUUID: input.TransactionDetailUUID,
	}
}
