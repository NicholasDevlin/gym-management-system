package services

import (
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/transaction"
	"gym/app/backend/utils/errors"
	"math/rand"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ITransactionService interface {
	CreateTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error)
	GetAllTransaction(filter transaction.TransactionReq) ([]transaction.TransactionRes, error)
	GetTransaction(filter transaction.TransactionReq) (transaction.TransactionRes, error)
	UpdateTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error)
	DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error)
}

type transactionService struct {
	transactionRepository repositories.ITransactionRepository
	userRepository     repositories.IUserRepository
}

func NewTransactionService(repo repositories.ITransactionRepository, userRepo repositories.IUserRepository) *transactionService {
	return &transactionService{
		transactionRepository: repo,
		userRepository:     userRepo,
	}
}

func (p *transactionService) CreateTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error) {
	entry := *transaction.ConvertReqToDto(input)
	entry.TransactionNo = generateTransactionNo()
	entry.TransactionDate = time.Now()
	resUser, err := p.userRepository.GetUser(entry.User)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	entry.UserId = resUser.Id
	res, err := p.transactionRepository.CreateTransaction(entry)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_CREATE_TRANSACTION
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (p *transactionService) GetAllTransaction(filter transaction.TransactionReq) ([]transaction.TransactionRes, error) {
	res, err := p.transactionRepository.GetAllTransaction(*transaction.ConvertReqToDto(filter))
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	var resTransaction []transaction.TransactionRes
	for i := 0; i < len(res); i++ {
		roleVm := transaction.ConvertDtoToRes(res[i])
		resTransaction = append(resTransaction, *roleVm)
	}
	return resTransaction, nil
}

func (p *transactionService) GetTransaction(filter transaction.TransactionReq) (transaction.TransactionRes, error) {
	res, err := p.transactionRepository.GetTransaction(*transaction.ConvertReqToDto(filter))

	if err != nil || (filter.Id == 0 && filter.UUID == uuid.Nil) {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (p *transactionService) UpdateTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error) {
	res, err := p.transactionRepository.GetTransaction(*transaction.ConvertReqToDto(input))
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	res, err = p.transactionRepository.UpdateTransaction(res, *transaction.ConvertReqToDto(input))
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_UPDATE_TRANSACTION
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (p *transactionService) DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error) {
	res, err := p.transactionRepository.GetTransaction(transaction.TransactionDto{UUID: id})
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}

	res, err = p.transactionRepository.DeleteTransaction(id.String())
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_DELETE_TRANSACTION
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func generateTransactionNo() string {
	var TransactionNo string
	now := time.Now().Format("020106-150405")
	randomNumber := strconv.Itoa(rand.Intn(8999) + 1000)
	TransactionNo = randomNumber + "/" + now + "/GYM"
	return TransactionNo
}
