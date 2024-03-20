package repositories

import (
	transactionDetail "gym/app/backend/models/transactionDetail"
	transactiondetail "gym/app/backend/models/transactionDetail"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ITransactionDetailRepository interface {
	CreateTransactionDetail(input transactionDetail.TransactionDetailDto) (transactionDetail.TransactionDetailDto, error)
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

func (td *transactionDetailRepository) CreateTransactionDetail(input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	dataTransactionDetail := transactiondetail.ConvertDtoToModel(input)
	dataTransactionDetail.UUID = uuid.NewV4()
	err := td.db.Create(&dataTransactionDetail).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}
	return *transactiondetail.ConvertModelToDto(*dataTransactionDetail), nil
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
