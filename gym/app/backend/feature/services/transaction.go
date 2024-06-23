package services

import (
	"gym/app/backend/feature/repositories"
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/transaction"
	transactiondetail "gym/app/backend/models/transactionDetail"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"
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
	GetAllTransaction(filter transaction.TransactionFilter) ([]transaction.TransactionRes, error)
	GetTransaction(filter transaction.TransactionFilter) (transaction.TransactionRes, error)
	SaveTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error)
	DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error)
}

type transactionService struct {
	transactionRepository       repositories.ITransactionRepository
	userRepository              repositories.IUserRepository
	membershipPlanRepository    repositories.IMembershipPlanRepository
	transactionDetailRepository repositories.ITransactionDetailRepository
	transactionDetailMember     repositories.ITransactionMemberDetailRepository
}

func NewTransactionService(
	repo repositories.ITransactionRepository,
	userRepo repositories.IUserRepository,
	membershipPlanRepo repositories.IMembershipPlanRepository,
	transactionDetailRepo repositories.ITransactionDetailRepository,
	transactionDetailMember repositories.ITransactionMemberDetailRepository,
) *transactionService {
	return &transactionService{
		transactionRepository:       repo,
		userRepository:              userRepo,
		membershipPlanRepository:    membershipPlanRepo,
		transactionDetailRepository: transactionDetailRepo,
		transactionDetailMember:     transactionDetailMember,
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

func (t *transactionService) GetAllTransaction(filter transaction.TransactionFilter) ([]transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetAllTransaction(filter)
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

func (t *transactionService) GetTransaction(filter transaction.TransactionFilter) (transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetTransaction(filter)

	if err != nil || (filter.TransactionId == 0 && filter.UUID == uuid.Nil && filter.MemberUUID == uuid.Nil) {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (t *transactionService) SaveTransaction(input transaction.TransactionReq) (transaction.TransactionRes, error) {
	var err error
	var data transaction.TransactionDto
	dto := *transaction.ConvertReqToDto(input)
	var transactionDetailDto transactiondetail.TransactionDetailDto
	isComplete := input.Status == consts.COMPLETE
	for i := range dto.TransactionDetail {
		transactionDetailDto = dto.TransactionDetail[i]
		if transactionDetailDto.Deleted != 0 && transactionDetailDto.UUID == uuid.Nil {
			continue
		}
		transactionDetailDto, err = SaveTransactionDetail(t, transactionDetailDto)
		if err != nil {
			return transaction.TransactionRes{}, err
		}

		for j, member := range transactionDetailDto.TransactionMemberDetail {
			if member.Deleted != 0 && member.UUID == uuid.Nil {
				continue
			}
			transactionDetailDto.TransactionMemberDetail[j], err = SaveTransactionDetailMember(t, member, transactionDetailDto.MembershipPlan.Duration, isComplete)
			if err != nil {
				return transaction.TransactionRes{}, err
			}
		}
		dto.TransactionDetail[i] = transactionDetailDto
	}

	if input.UUID != uuid.Nil {
		data, err = t.transactionRepository.GetTransaction(*&transaction.TransactionFilter{UUID: input.UUID})
		if data.Status == consts.COMPLETE {
			return transaction.TransactionRes{}, errors.ERR_COMPLETED_TRANSACTION
		}
		if err != nil {
			return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
		}
		if !input.TransactionDate.IsZero() {
			data.TransactionDate = input.TransactionDate
		}
		if input.Status != "" {
			data.Status = input.Status
		}
	} else {
		data.TransactionNo = generateTransactionNo()
		data.TransactionDate = time.Now()
		if input.Status == "" {
			data.Status = consts.WAITING_FOR_PAYMENT
		}
		data.UUID = uuid.NewV4()
	}
	data.TransactionDetail = dto.TransactionDetail

	resUser, err := t.userRepository.GetUser(*&transaction.ConvertReqToDto(input).User)
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	data.UserId = resUser.Id
	data.User = resUser

	res, err := t.transactionRepository.SaveTransaction(data)
	if err != nil {
		return transaction.TransactionRes{}, err
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func (t *transactionService) DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error) {
	res, err := t.transactionRepository.GetTransaction(transaction.TransactionFilter{UUID: id})
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_NOT_FOUND
	}
	isComplete := res.Status == consts.COMPLETE
	if isComplete {
		return transaction.TransactionRes{}, errors.ERR_COMPLETED_TRANSACTION
	}
	res, err = t.transactionRepository.DeleteTransaction(id.String())
	if err != nil {
		return transaction.TransactionRes{}, errors.ERR_DELETE_TRANSACTION
	}
	return *transaction.ConvertDtoToRes(res), nil
}

func generateTransactionNo() string {
	var TransactionNo string
	now := time.Now().Format("020106")
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

func SaveTransactionDetail(t *transactionService, input transactiondetail.TransactionDetailDto) (transactiondetail.TransactionDetailDto, error) {
	var existing transactiondetail.TransactionDetailDto
	if input.UUID == uuid.Nil {
		existing.UUID = uuid.NewV4()
	} else {
		existing, _ = t.transactionDetailRepository.GetTransactionDetail(input)
	}

	if input.Quantity == 0 {
		return transactiondetail.TransactionDetailDto{}, errors.ERR_QTY_EMPTY
	}

	membershipPlan, err := t.membershipPlanRepository.GetMembershipPlan(membershipplan.MembershipPlanDto{UUID: input.MembershipPlanUUID})

	if err != nil {
		return transactiondetail.TransactionDetailDto{}, errors.ERR_MEMBERSHIP_PLAN_NOT_FOUND
	}
	if input.Quantity != 0 {
		qty := len(input.TransactionMemberDetail)
		if qty != input.Quantity {
			qty = input.Quantity
		}
		existing.Quantity = qty
	}
	if input.Price != 0 {
		existing.Price = input.Price
	}
	existing.Deleted = input.Deleted
	existing.MembershipPlanId = membershipPlan.Id
	existing.MembershipPlan = membershipPlan
	existing.Price = int(membershipPlan.Price)
	existing.MembershipPlanUUID = membershipPlan.UUID
	existing.TransactionMemberDetail = input.TransactionMemberDetail
	return existing, nil
}

func SaveTransactionDetailMember(t *transactionService, input transactionmemberdetail.TransactionMemberDetailDto, duration int, isComplete bool) (transactionmemberdetail.TransactionMemberDetailDto, error) {
	var existing transactionmemberdetail.TransactionMemberDetailDto

	if input.UUID == uuid.Nil {
		existing.UUID = uuid.NewV4()
	} else {
		existing, _ = t.transactionDetailMember.GetTransactionMemberDetail(input)
	}

	user, err := t.userRepository.GetUser(user.UserDto{UUID: input.UserUUID})
	if err != nil {
		return transactionmemberdetail.TransactionMemberDetailDto{}, errors.ERR_USER_NOT_FOUND
	}

	if err != nil {
		return transactionmemberdetail.TransactionMemberDetailDto{}, errors.ERR_USER_NOT_FOUND
	}

	if isComplete {
		if user.SubscriptionExpirationDate != nil && !user.SubscriptionExpirationDate.IsZero() {
			membershipDuration := user.SubscriptionExpirationDate.AddDate(0, 0, duration)
			if user.SubscriptionExpirationDate.Before(time.Now()) {
				membershipDuration = time.Now().AddDate(0, 0, duration)
			}
			user.SubscriptionExpirationDate = &membershipDuration
		} else {
			membershipDuration := time.Now().AddDate(0, 0, duration)
			user.SubscriptionExpirationDate = &membershipDuration
		}
	}
	if input.AdditionalPrice != 0 {
		existing.AdditionalPrice = input.AdditionalPrice
	}
	existing.Deleted = input.Deleted
	existing.User = user
	existing.UserId = user.Id

	return existing, nil
}
