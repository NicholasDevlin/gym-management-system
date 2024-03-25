package services

import (
	"gym/app/backend/feature/repositories"
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/transaction"
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"
	"gym/app/backend/utils/consts"
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
	SaveTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error)
	DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error)
}

type transactionService struct {
	transactionRepository       repositories.ITransactionRepository
	userRepository              repositories.IUserRepository
	membershipPlanRepository    repositories.IMembershipPlanRepository
	transactionDetailRepository repositories.ITransactionDetailRepository
}

func NewTransactionService(repo repositories.ITransactionRepository, userRepo repositories.IUserRepository, membershipPlanRepo repositories.IMembershipPlanRepository, transactionDetailRepo repositories.ITransactionDetailRepository) *transactionService {
	return &transactionService{
		transactionRepository:       repo,
		userRepository:              userRepo,
		membershipPlanRepository:    membershipPlanRepo,
		transactionDetailRepository: transactionDetailRepo,
	}
}

func (t *transactionService) CreateTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error) {
	if input.TransactionDetail == nil {
		return transaction.TransactionRes{}, errors.ERR_TRANSACTION_DETAIL_EMPTY
	}
	entry := *transaction.ConvertReqToDto(input)
	entry.TransactionNo = generateTransactionNo()
	entry.TransactionDate = time.Now()
	resUser, err := t.userRepository.GetUser(entry.User)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	entry.UserId = resUser.Id
	entry.User = resUser
	res, err := t.transactionRepository.CreateTransaction(entry)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_CREATE_TRANSACTION
	}

	var transactionDetailDtos []transactiondetail.TransactionDetailDto
	var transactionDetailDto transactiondetail.TransactionDetailDto
	for i := range input.TransactionDetail {
		transactionDetailDto = *transactiondetail.ConvertReqToDto(input.TransactionDetail[i])
		transactionDetailDto.TransactionId = res.Id
		valid := transactionDetailValidation(transactionDetailDto)
		if !valid {
			continue
		}

		membershipPlan, err := t.membershipPlanRepository.GetMembershipPlan(membershipplan.MembershipPlanDto{UUID: input.TransactionDetail[i].MembershipPlanUUID})

		if err != nil {
			continue
		}
		transactionDetailDto.MembershipPlanId = membershipPlan.Id
		transactionDetailDto.MembershipPlan = membershipPlan
		res, err := t.transactionDetailRepository.SaveTransactionDetail(transactionDetailDto)
		if err == nil {
			transactionDetailDtos = append(transactionDetailDtos, res)
		}
	}
	res.TransactionDetail = transactionDetailDtos
	return *transaction.ConvertDtoToRes(res), nil
}

func (t *transactionService) GetAllTransaction(filter transaction.TransactionReq) ([]transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetAllTransaction(*transaction.ConvertReqToDto(filter))
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

func (t *transactionService) GetTransaction(filter transaction.TransactionReq) (transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetTransaction(*transaction.ConvertReqToDto(filter))

	if err != nil || (filter.Id == 0 && filter.UUID == uuid.Nil) {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (t *transactionService) SaveTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error) {
	var data transaction.TransactionDto
	var err error

	dto := *transaction.ConvertReqToDto(input)
	var transactionDetailDto transactiondetail.TransactionDetailDto
	for i := range dto.TransactionDetail {
		transactionDetailDto = dto.TransactionDetail[i]
		if transactionDetailDto.Quantity != len(transactionDetailDto.TransactionMemberDetail) {
			return transaction.TransactionRes{}, errors.ERR_TRANSACTION_MEMBER_EMPTY
		}

		if transactionDetailDto.Quantity == 0 {
			return transaction.TransactionRes{}, errors.ERR_QTY_EMPTY
		}

		membershipPlan, err := t.membershipPlanRepository.GetMembershipPlan(membershipplan.MembershipPlanDto{UUID: transactionDetailDto.MembershipPlanUUID})

		if err != nil {
			return transaction.TransactionRes{}, errors.ERR_MEMBERSHIP_PLAN_NOT_FOUND
		}
		transactionDetailDto.MembershipPlanId = membershipPlan.Id
		transactionDetailDto.MembershipPlan = membershipPlan

		for _, member := range transactionDetailDto.TransactionMemberDetail {
			user, err := t.userRepository.GetUser(user.UserDto{UUID: member.UserUUID})
			if err != nil {
				return transaction.TransactionRes{}, errors.ERR_USER_NOT_FOUND
			}
			member.UserId = user.Id
		}
		dto.TransactionDetail[i] = transactionDetailDto
	}

	if input.UUID != uuid.Nil {
		data, err = t.transactionRepository.GetTransaction(*transaction.ConvertReqToDto(input))
		if err != nil {
			return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
		}
	} else {
		data.TransactionNo = generateTransactionNo()
		data.TransactionDate = time.Now()
		data.Status = consts.WAITING_FOR_PAYMENT
		data.UUID = uuid.NewV4()
	}

	resUser, err := t.userRepository.GetUser(*&transaction.ConvertReqToDto(input).User)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	data.UserId = resUser.Id
	data.User = resUser

	res, err := t.transactionRepository.SaveTransaction(data, dto)
	if err != nil {
		return transaction.TransactionRes{}, err
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (t *transactionService) DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetTransaction(transaction.TransactionDto{UUID: id})
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}

	res, err = t.transactionRepository.DeleteTransaction(id.String())
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

func transactionDetailValidation(input transactiondetail.TransactionDetailDto) bool {
	if input.Quantity <= 0 {
		input.Quantity = 1
	}
	return true
}
