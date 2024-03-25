package repositories

import (
	transactionDetail "gym/app/backend/models/transactionDetail"
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ITransactionDetailRepository interface {
	SaveTransactionDetail(input transactionDetail.TransactionDetailDto) (transactionDetail.TransactionDetailDto, error)
	GetAllTransactionDetail(filter transactionDetail.TransactionDetailDto) ([]transactionDetail.TransactionDetailDto, error)
	GetTransactionDetail(filter transactionDetail.TransactionDetailDto) (transactionDetail.TransactionDetailDto, error)
	UpdateTransactionDetail(data, input transactionDetail.TransactionDetailDto) (transactionDetail.TransactionDetailDto, error)
	DeleteTransactionDetail(id string) (transactionDetail.TransactionDetailDto, error)
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) *transactionDetailRepository {
	return &transactionDetailRepository{db}
}

func (td *transactionDetailRepository) SaveTransactionDetail(input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	var data transactionDetail.TransactionDetail
	if input.UUID == uuid.Nil {
		data.UUID = uuid.NewV4()
		} else {
			dataTransactionDetail, err := td.GetTransactionDetail(input)
			if err != nil {
				return transactiondetail.TransactionDetailDto{}, errors.ERR_TRANSACTION_DETAIL_NOT_FOUND
			}
			data = *transactionDetail.ConvertDtoToModel(dataTransactionDetail)
		}
	if input.TransactionId != 0 {
		data.TransactionId = input.TransactionId
	}
	if input.MembershipPlanId != 0 {
		data.MembershipPlanId = input.MembershipPlanId
	}
	if input.Quantity != 0 {
		data.Quantity = input.Quantity
	}
	err := td.db.Save(&data).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}
	return *transactiondetail.ConvertModelToDto(data), nil
}

func (td *transactionDetailRepository) GetAllTransactionDetail(filter transactiondetail.TransactionDetailDto) ([]transactiondetail.TransactionDetailDto, error) {
	var allTransactionDetail []transactiondetail.TransactionDetail
	var resAllTransactionDetail []transactiondetail.TransactionDetailDto

	err := td.db.Preload("MembershipPlan").Find(&allTransactionDetail).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allTransactionDetail); i++ {
		transactionDetail := transactiondetail.ConvertModelToDto(allTransactionDetail[i])
		resAllTransactionDetail = append(resAllTransactionDetail, *transactionDetail)
	}
	return resAllTransactionDetail, nil
}

func (td *transactionDetailRepository) GetTransactionDetail(filter transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	if filter.UUID == uuid.Nil && filter.Id == 0 {
		return transactiondetail.TransactionDetailDto{}, nil
	}
	var model transactiondetail.TransactionDetail
	query := td.db.Model(&transactiondetail.TransactionDetail{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.Preload("User").First(&model).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}
	return *transactiondetail.ConvertModelToDto(model), nil
}

func (td *transactionDetailRepository) UpdateTransactionDetail(data, input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	transactionDetailData := *transactiondetail.ConvertDtoToModel(data)

	if input.MembershipPlanId != 0 {
		transactionDetailData.MembershipPlanId = input.MembershipPlanId
	}
	if input.Quantity != 0 {
		transactionDetailData.Quantity = input.Quantity
	}

	if err := td.db.Save(&transactionDetailData).Error; err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}
	return *transactiondetail.ConvertModelToDto(transactionDetailData), nil
}

func (td *transactionDetailRepository) DeleteTransactionDetail(id string) (transactiondetail.TransactionDetailDto, error) {
	transactionDetailData := transactiondetail.TransactionDetail{}

	err := td.db.Delete(&transactionDetailData, "uuid = ?", id).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}

	return *transactiondetail.ConvertModelToDto(transactionDetailData), nil
}
