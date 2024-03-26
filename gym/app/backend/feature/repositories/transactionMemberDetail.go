package repositories

import (
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ITransactionMemberDetailRepository interface {
	SaveTransactionMemberDetail(input transactionmemberdetail.TransactionMemberDetailDto) (transactionmemberdetail.TransactionMemberDetailDto, error)
	GetTransactionMemberDetail(filter transactionmemberdetail.TransactionMemberDetailDto) (transactionmemberdetail.TransactionMemberDetailDto, error)
	// DeleteTransactionMemberDetail(id string) (transactionmemberdetail.TransactionMemberDetailDto, error)
}

type transactionMemberDetailRepository struct {
	db *gorm.DB
}

func NewTransactionMemberDetailRepository(db *gorm.DB) *transactionMemberDetailRepository {
	return &transactionMemberDetailRepository{db}
}

func (td *transactionMemberDetailRepository) SaveTransactionMemberDetail(input transactionmemberdetail.TransactionMemberDetailDto) (transactionmemberdetail.TransactionMemberDetailDto, error) {
	dataTransactionMemberDetail := transactionmemberdetail.ConvertDtoToModel(input)
	if input.UUID != uuid.Nil {
		err := td.db.First(&dataTransactionMemberDetail, "uuid = ?", input.UUID)
		if err != nil {
			dataTransactionMemberDetail.UUID = uuid.NewV4()
		}
	} else {
		dataTransactionMemberDetail.UUID = uuid.NewV4()
	}
	if input.TransactionDetailId != 0 {
		dataTransactionMemberDetail.TransactionDetailId = input.TransactionDetailId
	}
	if input.UserId != 0 {
		dataTransactionMemberDetail.UserId = input.UserId
	}

	err := td.db.Save(&dataTransactionMemberDetail).Error
	if err != nil {
		return transactionmemberdetail.TransactionMemberDetailDto{}, err
	}
	return *transactionmemberdetail.ConvertModelToDto(*dataTransactionMemberDetail), nil
}

func (td *transactionMemberDetailRepository) GetTransactionMemberDetail(filter transactionmemberdetail.TransactionMemberDetailDto) (transactionmemberdetail.TransactionMemberDetailDto, error) {
	if filter.UUID == uuid.Nil && filter.Id == 0 {
		return transactionmemberdetail.TransactionMemberDetailDto{}, nil
	}
	var model transactionmemberdetail.TransactionMemberDetail
	query := td.db.Model(&transactionmemberdetail.TransactionMemberDetail{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.Preload("User").First(&model).Error
	if err != nil {
		return transactionmemberdetail.TransactionMemberDetailDto{}, err
	}
	return *transactionmemberdetail.ConvertModelToDto(model), nil
}
