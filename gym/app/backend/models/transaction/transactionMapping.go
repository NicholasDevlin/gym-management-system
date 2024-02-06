package transaction

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func ConvertReqToDto(input TransactionReq) *TransactionDto {
	return &TransactionDto{
		Id:              input.Id,
		UUID:            input.UUID,
		UserUUID:        input.UserUUID,
		TransactionDate: input.TransactionDate,
		TransactionNo:   input.TransactionNo,
		Status:          input.Status,
	}
}

func ConvertDtoToModel(input TransactionDto) *Transaction {
	return &Transaction{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:            input.UUID,
		UserId:          input.UserId,
		TransactionDate: input.TransactionDate,
		TransactionNo:   input.TransactionNo,
		Status:          input.Status,
	}
}

func ConvertModelToDto(input Transaction) *TransactionDto {
	return &TransactionDto{
		Id:              input.ID,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
		UUID:            input.UUID,
		UserId:          input.UserId,
		TransactionDate: input.TransactionDate,
		TransactionNo:   input.TransactionNo,
		Status:          input.Status,
	}
}

func ConvertDtoToRes(input TransactionDto) *TransactionRes {
	return &TransactionRes{
		UUID:              input.UUID,
		TransactionDate:   input.TransactionDate,
		TransactionNo:     input.TransactionNo,
		Status:            input.Status,
		User:              *user.ConvertDtoToRes(input.User),
		TransactionDetail: *ConvertManyDtoToRes(input.TransactionDetail),
	}
}

func ConvertManyDtoToRes(input []transactiondetail.TransactionDetailDto) *[]transactiondetail.TransactionDetailRes {
	var result []transactiondetail.TransactionDetailRes
	for i := range input {
		res := *transactiondetail.ConvertDtoToRes(input[i])
		result = append(result, res)
	}
	return &result
}
