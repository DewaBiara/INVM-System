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
	CreatePurchaseDetail(ctx context.Context, purchaseDetail *entity.PurchaseDetail) error
	UpdatePurchaseDetail(ctx context.Context, purchaseDetail *entity.PurchaseDetail) error
	GetSinglePurchaseDetail(ctx context.Context, purchaseDetailID string) (*entity.PurchaseDetail, error)
	GetPagePurchaseDetail(ctx context.Context, limit int, offset int) (*entity.PurchaseDetails, error)
}
