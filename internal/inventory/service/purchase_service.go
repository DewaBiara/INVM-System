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
	DeletePurchase(ctx context.Context, purchaseID string) error
}
