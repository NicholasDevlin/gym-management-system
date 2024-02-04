package services

import (
	"gym/app/backend/feature/repositories"
	"gym/app/backend/models/purchase"
	"gym/app/backend/utils/errors"
	"math/rand"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type IPurchaseService interface {
	CreatePurchase(input purchase.PurchaseReq) (purchase.PurchaseRes, error)
	GetAllPurchase(filter purchase.PurchaseReq) ([]purchase.PurchaseRes, error)
	GetPurchase(filter purchase.PurchaseReq) (purchase.PurchaseRes, error)
	UpdatePurchase(input purchase.PurchaseReq) (purchase.PurchaseRes, error)
	DeletePurchase(id uuid.UUID) (purchase.PurchaseRes, error)
}

type purchaseService struct {
	purchaseRepository repositories.IPurchaseRepository
	userRepository     repositories.IUserRepository
}

func NewPurchaseService(repo repositories.IPurchaseRepository, userRepo repositories.IUserRepository) *purchaseService {
	return &purchaseService{
		purchaseRepository: repo,
		userRepository:     userRepo,
	}
}

func (p *purchaseService) CreatePurchase(input purchase.PurchaseReq) (purchase.PurchaseRes, error) {
	entry := *purchase.ConvertReqToDto(input)
	entry.PurchaseNo = generatePurchaseNo()
	entry.PurchaseDate = time.Now()
	resUser, err := p.userRepository.GetUser(entry.User)
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_NOT_FOUND
	}
	entry.UserId = resUser.Id
	res, err := p.purchaseRepository.CreatePurchase(entry)
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_CREATE_PURCHASE
	}
	return *purchase.ConvertDtoToRes(res), nil
}

func (p *purchaseService) GetAllPurchase(filter purchase.PurchaseReq) ([]purchase.PurchaseRes, error) {
	res, err := p.purchaseRepository.GetAllPurchase(*purchase.ConvertReqToDto(filter))
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	var resPurchase []purchase.PurchaseRes
	for i := 0; i < len(res); i++ {
		roleVm := purchase.ConvertDtoToRes(res[i])
		resPurchase = append(resPurchase, *roleVm)
	}
	return resPurchase, nil
}

func (p *purchaseService) GetPurchase(filter purchase.PurchaseReq) (purchase.PurchaseRes, error) {
	res, err := p.purchaseRepository.GetPurchase(*purchase.ConvertReqToDto(filter))

	if err != nil || (filter.Id == 0 && filter.UUID == uuid.Nil) {
		return purchase.PurchaseRes{}, errors.ERR_NOT_FOUND
	}
	return *purchase.ConvertDtoToRes(res), nil
}

func (p *purchaseService) UpdatePurchase(input purchase.PurchaseReq) (purchase.PurchaseRes, error) {
	res, err := p.purchaseRepository.GetPurchase(*purchase.ConvertReqToDto(input))
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_NOT_FOUND
	}
	res, err = p.purchaseRepository.UpdatePurchase(res, *purchase.ConvertReqToDto(input))
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_UPDATE_ROLE
	}
	return *purchase.ConvertDtoToRes(res), nil
}

func (p *purchaseService) DeletePurchase(id uuid.UUID) (purchase.PurchaseRes, error) {
	res, err := p.purchaseRepository.GetPurchase(purchase.PurchaseDto{UUID: id})
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_NOT_FOUND
	}

	res, err = p.purchaseRepository.DeletePurchase(id.String())
	if err != nil {
		return purchase.PurchaseRes{}, errors.ERR_DELETE_USER
	}
	return *purchase.ConvertDtoToRes(res), nil
}

func generatePurchaseNo() string {
	var purchaseNo string
	now := time.Now().Format("020106-150405")
	randomNumber := strconv.Itoa(rand.Intn(8999) + 1000)
	purchaseNo = randomNumber + "/" + now + "/GYM"
	return purchaseNo
}
