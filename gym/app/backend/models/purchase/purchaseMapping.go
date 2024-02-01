package purchase

import "gorm.io/gorm"

func ConvertReqToDto(input PurchaseReq) *PurchaseDto {
	return &PurchaseDto{
		Id:           input.Id,
		UUID:         input.UUID,
		UserId:       input.UserId,
		PurchaseDate: input.PurchaseDate,
		PurchaseNo:   input.PurchaseNo,
	}
}

func ConvertDtoToModel(input PurchaseDto) *Purchase {
	return &Purchase{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:         input.UUID,
		UserId:       input.UserId,
		PurchaseDate: input.PurchaseDate,
		PurchaseNo:   input.PurchaseNo,
	}
}

func ConvertModelToDto(input Purchase) *PurchaseDto {
	return &PurchaseDto{
		Id:           input.ID,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
		UUID:         input.UUID,
		UserId:       input.UserId,
		PurchaseDate: input.PurchaseDate,
		PurchaseNo:   input.PurchaseNo,
	}
}

func ConvertDtoToRes(input PurchaseDto) *PurchaseRes {
	return &PurchaseRes{
		Id:           input.Id,
		UUID:         input.UUID,
		UserId:       input.UserId,
		PurchaseDate: input.PurchaseDate,
		PurchaseNo:   input.PurchaseNo,
	}
}
