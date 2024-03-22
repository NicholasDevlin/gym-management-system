package repositories

import (
	"gym/app/backend/models/transaction"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(input transaction.TransactionDto) (transaction.TransactionDto, error)
	GetAllTransaction(filter transaction.TransactionDto) ([]transaction.TransactionDto, error)
	GetTransaction(filter transaction.TransactionDto) (transaction.TransactionDto, error)
	SaveTransaction(data, input transaction.TransactionDto) (transaction.TransactionDto, error)
	DeleteTransaction(id string) (transaction.TransactionDto, error)
}

type transactionRepository struct {
	db *gorm.DB
	transactionDetailRepository ITransactionDetailRepository
	transactionMemberDetailRepository ITransactionMemberDetailRepository
}

func NewTransactionRepository(db *gorm.DB, transactionDetailRepo ITransactionDetailRepository, transactionMemberDetailRepo ITransactionMemberDetailRepository) *transactionRepository {
	return &transactionRepository{
		db,
		transactionDetailRepo,
		transactionMemberDetailRepo,
	}
}

func (t *transactionRepository) CreateTransaction(input transaction.TransactionDto) (transaction.TransactionDto, error) {
	dataTransaction := transaction.ConvertDtoToModel(input)
	dataTransaction.UUID = uuid.NewV4()
	err := t.db.Create(&dataTransaction).Error
	if err != nil {
		return transaction.TransactionDto{}, err
	}
	return *transaction.ConvertModelToDto(*dataTransaction), nil
}

func (t *transactionRepository) GetAllTransaction(filter transaction.TransactionDto) ([]transaction.TransactionDto, error) {
	var allTransaction []transaction.Transaction
	var resAllTransaction []transaction.TransactionDto

	query := t.db.Model(&transaction.Transaction{})
	if filter.TransactionNo != "" {
		query = query.Where("transaction_no LIKE ?", "%"+filter.TransactionNo+"%")
	}

	err := query.Preload("User").Preload("TransactionDetail").Preload("TransactionDetail.MembershipPlan").Find(&allTransaction).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allTransaction); i++ {
		transaction := transaction.ConvertModelToDto(allTransaction[i])
		resAllTransaction = append(resAllTransaction, *transaction)
	}
	return resAllTransaction, nil
}

func (t *transactionRepository) GetTransaction(filter transaction.TransactionDto) (transaction.TransactionDto, error) {
	var model transaction.Transaction
	query := t.db.Model(&transaction.Transaction{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.Preload("User").Preload("TransactionDetail").Preload("TransactionDetail.MembershipPlan").First(&model).Error
	if err != nil {
		return transaction.TransactionDto{}, err
	}
	return *transaction.ConvertModelToDto(model), nil
}

func (t *transactionRepository) SaveTransaction(data, input transaction.TransactionDto) (transaction.TransactionDto, error) {
	tx := t.db.Begin()
	transactionData := *transaction.ConvertDtoToModel(data)

	if !input.TransactionDate.IsZero() {
		transactionData.TransactionDate = input.TransactionDate
	}
	if input.Status != "" {
		transactionData.Status = input.Status
	}
	if input.UserId != 0 {
		transactionData.UserId = input.UserId
	}

	if err := tx.Save(&transactionData).Error; err != nil {
		tx.Rollback()
		return transaction.TransactionDto{}, err
	}

	// save transaction detail
	for _, detail := range input.TransactionDetail {
		detailRes, err := t.transactionDetailRepository.SaveTransactionDetail(detail)
		if err != nil {
			tx.Rollback()
			return transaction.TransactionDto{}, err
		}
		// save transaction detail member 
		for _, member := range detail.TransactionMemberDetail {
			member.TransactionDetailId = detailRes.Id
			memberRes, err := t.transactionMemberDetailRepository.SaveTransactionMemberDetail(member)
			if err != nil {
				tx.Rollback()
				return transaction.TransactionDto{}, err
			}
			member = memberRes
		}
		detail = detailRes
	}

	tx.Commit()
	return *transaction.ConvertModelToDto(transactionData), nil
}

func (t *transactionRepository) DeleteTransaction(id string) (transaction.TransactionDto, error) {
	transactionData := transaction.Transaction{}

	err := t.db.Delete(&transactionData, "uuid = ?", id).Error
	if err != nil {
		return transaction.TransactionDto{}, err
	}

	return *transaction.ConvertModelToDto(transactionData), nil
}
