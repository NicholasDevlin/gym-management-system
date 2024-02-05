package transaction

import "gorm.io/gorm"

func ConvertReqToDto(input TransactionReq) *TransactionDto {
	return &TransactionDto{
		Id:              input.Id,
		UUID:            input.UUID,
		UserId:          input.UserId,
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
		UUID:            input.UUID,
		TransactionDate: input.TransactionDate,
		TransactionNo:   input.TransactionNo,
		Status:          input.Status,
	}
}
