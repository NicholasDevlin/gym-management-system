package repositories

import (
	"gym/app/backend/models/transaction"
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/utils/errors"

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
	db                                *gorm.DB
	transactionDetailRepository       ITransactionDetailRepository
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

	err := query.Preload("User").Preload("TransactionDetail").Preload("TransactionDetail.MembershipPlan").Preload("TransactionDetail.TransactionMemberDetail").Preload("TransactionDetail.TransactionMemberDetail.User").Find(&allTransaction).Error
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

	err := query.Preload("User").Preload("TransactionDetail").Preload("TransactionDetail.MembershipPlan").Preload("TransactionDetail.TransactionMemberDetail").Preload("TransactionDetail.TransactionMemberDetail.User").First(&model).Error
	if err != nil {
		return transaction.TransactionDto{}, err
	}
	return *transaction.ConvertModelToDto(model), nil
}

func (t *transactionRepository) SaveTransaction(data, input transaction.TransactionDto) (transaction.TransactionDto, error) {
	var err error
	tx := t.db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()

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
		if data.Id != 0 {
			err = errors.ERR_UPDATE_TRANSACTION
			tx.SavePoint("transaction")
		} else {
			err = errors.ERR_CREATE_TRANSACTION
		}
		return transaction.TransactionDto{}, err
	}

	// save transaction detail
	var transactionDetailData transactiondetail.TransactionDetail
	var existingDetail transactiondetail.TransactionDetailDto
	for _, detail := range input.TransactionDetail {
		detail.TransactionId = transactionData.ID
		if detail.UUID != uuid.Nil {
			existingDetail, err = t.transactionDetailRepository.GetTransactionDetail(detail)
			if err != nil {
				return transaction.TransactionDto{}, errors.ERR_TRANSACTION_DETAIL_NOT_FOUND
			}
		}
		transactionDetailData = SaveTransactionDetail(existingDetail, detail)
		err = tx.Save(&transactionDetailData).Error
		if err != nil {
			tx.RollbackTo("transaction")
			err = errors.ERR_CREATE_TRANSACTION_DETAIL
			return transaction.TransactionDto{}, err
		}
		// // save transaction detail member
		// for _, member := range detail.TransactionMemberDetail {
		// 	fmt.Println("member")
		// 	member.TransactionDetailId = detailRes.Id
		// 	_, err := t.transactionMemberDetailRepository.SaveTransactionMemberDetail(member)
		// 	if err != nil {
		// 		// tx.Rollback()
		// 		return transaction.TransactionDto{}, err
		// 	}
		// member = memberRes
		//}
		// detail = detailRes
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

func SaveTransactionDetail(existing, input transactiondetail.TransactionDetailDto) transactiondetail.TransactionDetail {
	var data transactiondetail.TransactionDetail
	if input.UUID == uuid.Nil {
		data.UUID = uuid.NewV4()
	} else {
		data = *transactiondetail.ConvertDtoToModel(existing)
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
	return data
}
