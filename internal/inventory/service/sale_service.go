package service

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
)

type SaleService interface {
	CreateSale(ctx context.Context, sale *dto.CreateSaleRequest) error
	UpdateSale(ctx context.Context, saleID uint, updateSale *dto.UpdateSaleRequest) error
	GetSingleSale(ctx context.Context, saleID string) (*dto.GetSingleSaleResponse, error)
	GetPageSale(ctx context.Context, limit int, offset int) (*dto.GetPageSalesResponse, error)
	CreateSaleDetail(ctx context.Context, saleDetail *dto.CreateSaleDetailRequest) error
	UpdateSaleDetail(ctx context.Context, saleDetailID uint, updateSale *dto.UpdateSaleDetailRequest) error
	GetSingleSaleDetail(ctx context.Context, saleDetailID string) (*dto.GetSingleSaleDetailResponse, error)
	GetPageSaleDetail(ctx context.Context, limit int, offset int) (*dto.GetPageSaleDetailsResponse, error)
}
