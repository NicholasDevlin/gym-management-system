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
		User:            *user.ConvertReqToDto(input.User),
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
		User:            *user.ConvertDtoToModel(input.User),
	}
}

func ConvertModelToDto(input Transaction) *TransactionDto {
	return &TransactionDto{
		Id:                input.ID,
		CreatedAt:         input.CreatedAt,
		UpdatedAt:         input.UpdatedAt,
		UUID:              input.UUID,
		UserId:            input.UserId,
		TransactionDate:   input.TransactionDate,
		TransactionNo:     input.TransactionNo,
		Status:            input.Status,
		UserUUID:          input.User.UUID,
		User:              *user.ConvertModelToDto(input.User),
		TransactionDetail: *ConvertModelToDtos(input.TransactionDetail),
	}
}

func ConvertDtoToRes(input TransactionDto) *TransactionRes {
	transactionDetail, total := ConvertDtosToRes(input.TransactionDetail)
	return &TransactionRes{
		UUID:              input.UUID,
		TransactionDate:   input.TransactionDate,
		TransactionNo:     input.TransactionNo,
		Status:            input.Status,
		UserUUID:          input.UserUUID,
		User:              *user.ConvertDtoToRes(input.User),
		TransactionDetail: *transactionDetail,
		Total:             total,
	}
}

func ConvertDtosToRes(input []transactiondetail.TransactionDetailDto) (*[]transactiondetail.TransactionDetailRes, int64) {
	var result []transactiondetail.TransactionDetailRes
	var total int64
	for i := range input {
		res := *transactiondetail.ConvertDtoToRes(input[i])
		total += res.Subtotal
		result = append(result, res)
	}
	return &result, total
}

func ConvertModelToDtos(input []transactiondetail.TransactionDetail) *[]transactiondetail.TransactionDetailDto {
	var result []transactiondetail.TransactionDetailDto
	for i := range input {
		res := *transactiondetail.ConvertModelToDto(input[i])
		result = append(result, res)
	}
	return &result
}
