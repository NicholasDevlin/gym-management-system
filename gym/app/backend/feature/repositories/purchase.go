package repositories

import (
	"gym/app/backend/models/purchase"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IPurchaseRepository interface {
	CreatePurchase(input purchase.PurchaseDto) (purchase.PurchaseDto, error)
	GetAllPurchase(filter purchase.PurchaseDto) ([]purchase.PurchaseDto, error)
	GetPurchase(filter purchase.PurchaseDto) (purchase.PurchaseDto, error)
	UpdatePurchase(data, input purchase.PurchaseDto) (purchase.PurchaseDto, error)
	DeletePurchase(id string) (purchase.PurchaseDto, error)
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *purchaseRepository {
	return &purchaseRepository{db}
}


func (p *purchaseRepository) CreatePurchase(input purchase.PurchaseDto) (purchase.PurchaseDto, error) {
	dataPurchase := purchase.ConvertDtoToModel(input)
	dataPurchase.UUID = uuid.NewV4()
	err := p.db.Create(&dataPurchase).Error
	if err != nil {
		return purchase.PurchaseDto{}, err
	}
	return *purchase.ConvertModelToDto(*dataPurchase), nil
}

func (p *purchaseRepository) GetAllPurchase(filter purchase.PurchaseDto) ([]purchase.PurchaseDto, error) {
	var allPurchase []purchase.Purchase
	var resAllPurchase []purchase.PurchaseDto

	query := p.db.Model(&purchase.Purchase{})
	if filter.PurchaseNo != "" {
		query = query.Where("purchase_no LIKE ?", "%"+filter.PurchaseNo+"%")
	}

	err := query.Find(&allPurchase).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(allPurchase); i++ {
		purchase := purchase.ConvertModelToDto(allPurchase[i])
		resAllPurchase = append(resAllPurchase, *purchase)
	}
	return resAllPurchase, nil
}

func (p *purchaseRepository) GetPurchase(filter purchase.PurchaseDto) (purchase.PurchaseDto, error) {
	var model purchase.Purchase
	query := p.db.Model(&purchase.Purchase{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}

	err := query.First(&model).Error
	if err != nil {
		return purchase.PurchaseDto{}, err
	}
	return *purchase.ConvertModelToDto(model), nil
}

func (p *purchaseRepository) UpdatePurchase(data, input purchase.PurchaseDto) (purchase.PurchaseDto, error) {
	purchaseData := *purchase.ConvertDtoToModel(data)

	if !input.PurchaseDate.IsZero() {
		purchaseData.PurchaseDate = input.PurchaseDate
	}

	if err := p.db.Save(&purchaseData).Error; err != nil {
		return purchase.PurchaseDto{}, err
	}
	return *purchase.ConvertModelToDto(purchaseData), nil
}

func (p *purchaseRepository) DeletePurchase(id string) (purchase.PurchaseDto, error) {
	purchaseData := purchase.Purchase{}

	err := p.db.Delete(&purchaseData, "uuid = ?", id).Error
	if err != nil {
		return purchase.PurchaseDto{}, err
	}

	return *purchase.ConvertModelToDto(purchaseData), nil
}