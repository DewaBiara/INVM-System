package service

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
)

type PurchaseService interface {
	CreatePurchase(ctx context.Context, purchase *dto.CreatePurchaseRequest) error
	UpdatePurchase(ctx context.Context, purchaseID uint, updatePurchase *dto.UpdatePurchaseRequest) error
	GetSinglePurchase(ctx context.Context, purchaseID string) (*dto.GetSinglePurchaseResponse, error)
	GetPagePurchase(ctx context.Context, limit int, offset int) (*dto.GetPagePurchasesResponse, error)
	CreatePurchaseDetail(ctx context.Context, purchaseDetail *dto.CreatePurchaseDetailRequest) error
	UpdatePurchaseDetail(ctx context.Context, purchaseDetailID uint, updatePurchase *dto.UpdatePurchaseDetailRequest) error
	GetSinglePurchaseDetail(ctx context.Context, purchaseDetailID string) (*dto.GetSinglePurchaseDetailResponse, error)
	GetPagePurchaseDetail(ctx context.Context, limit int, offset int) (*dto.GetPagePurchaseDetailsResponse, error)
}
