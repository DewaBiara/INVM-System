package repository

import (
	"context"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type PurchaseRepository interface {
	CreatePurchase(ctx context.Context, purchase *entity.Purchase) error
	UpdatePurchase(ctx context.Context, purchase *entity.Purchase) error
	GetSinglePurchase(ctx context.Context, purchaseID string) (*entity.Purchase, error)
	GetPagePurchase(ctx context.Context, limit int, offset int) (*entity.Purchases, error)
	DeletePurchase(ctx context.Context, purchaseID string) error
}
