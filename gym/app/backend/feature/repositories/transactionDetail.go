package repositories

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	transactionDetaildetail "gym/app/backend/models/transactionDetail"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ITransactionDetailDetailRepository interface {
	CreateTransactionDetailDetail(input transactionDetaildetail.TransactionDetailDto) (transactionDetaildetail.TransactionDetailDto, error)
	GetAllTransactionDetailDetail(filter transactionDetaildetail.TransactionDetailDto) ([]transactionDetaildetail.TransactionDetailDto, error)
	GetTransactionDetailDetail(filter transactionDetaildetail.TransactionDetailDto) (transactionDetaildetail.TransactionDetailDto, error)
	UpdateTransactionDetailDetail(data, input transactionDetaildetail.TransactionDetailDto) (transactionDetaildetail.TransactionDetailDto, error)
	DeleteTransactionDetailDetail(id string) (transactionDetaildetail.TransactionDetailDto, error)
}

type transactionDetailDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailDetailRepository(db *gorm.DB) *transactionDetailDetailRepository {
	return &transactionDetailDetailRepository{db}
}


func (td *transactionDetailDetailRepository) CreateTransactionDetail(input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	dataTransactionDetail := transactiondetail.ConvertDtoToModel(input)
	dataTransactionDetail.UUID = uuid.NewV4()
	err := td.db.Create(&dataTransactionDetail).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}
	return *transactiondetail.ConvertModelToDto(*dataTransactionDetail), nil
}

func (td *transactionDetailDetailRepository) GetAllTransactionDetail(filter transactiondetail.TransactionDetailDto) ([]transactiondetail.TransactionDetailDto, error) {
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

func (td *transactionDetailDetailRepository) GetTransactionDetail(filter transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
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

func (td *transactionDetailDetailRepository) UpdateTransactionDetail(data, input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
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

func (td *transactionDetailDetailRepository) DeleteTransactionDetail(id string) (transactiondetail.TransactionDetailDto, error) {
	transactionDetailData := transactiondetail.TransactionDetail{}

	err := td.db.Delete(&transactionDetailData, "uuid = ?", id).Error
	if err != nil {
		return transactiondetail.TransactionDetailDto{}, err
	}

	return *transactiondetail.ConvertModelToDto(transactionDetailData), nil
}
