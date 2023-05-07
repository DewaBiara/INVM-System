package impl

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
	"github.com/google/uuid"
)

type (
	PurchaseServiceImpl struct {
		purchaseRepository repository.PurchaseRepository
	}
)

func NewPurchaseServiceImpl(purchaseRepository repository.PurchaseRepository) service.PurchaseService {
	return &PurchaseServiceImpl{
		purchaseRepository: purchaseRepository,
	}
}

func (u *PurchaseServiceImpl) CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseRequest) error {

	purchaseEntity := purchase.ToEntity()
	purchaseEntity.ID = uint(uuid.New().ID())

	err := u.purchaseRepository.CreatePurchase(ctx, purchaseEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *PurchaseServiceImpl) GetSinglePurchase(ctx context.Context, purchaseID string) (*dto.GetSinglePurchaseResponse, error) {
	purchase, err := d.purchaseRepository.GetSinglePurchase(ctx, purchaseID)
	if err != nil {
		return nil, err
	}

	var purchaseResponse = dto.NewGetSinglePurchaseResponse(purchase)

	return purchaseResponse, nil
}

func (u *PurchaseServiceImpl) GetPagePurchase(ctx context.Context, page int, limit int) (*dto.GetPagePurchasesResponse, error) {
	offset := (page - 1) * limit

	purchases, err := u.purchaseRepository.GetPagePurchase(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPagePurchasesResponse(purchases), nil
}

func (u *PurchaseServiceImpl) UpdatePurchase(ctx context.Context, purchaseID uint, updatePurchase *dto.UpdatePurchaseRequest) error {
	purchase := updatePurchase.ToEntity()
	purchase.ID = purchaseID

	return u.purchaseRepository.UpdatePurchase(ctx, purchase)
}

func (d *PurchaseServiceImpl) DeletePurchase(ctx context.Context, purchaseID string) error {

	return d.purchaseRepository.DeletePurchase(ctx, purchaseID)
}
