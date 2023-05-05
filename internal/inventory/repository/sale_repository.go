package repository

import (
	"context"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type SaleRepository interface {
	CreateSale(ctx context.Context, sale *entity.Sale) error
	UpdateSale(ctx context.Context, sale *entity.Sale) error
	GetSingleSale(ctx context.Context, saleID string) (*entity.Sale, error)
	GetPageSale(ctx context.Context, limit int, offset int) (*entity.Sales, error)
	CreateSaleDetail(ctx context.Context, saleDetail *entity.SaleDetail) error
	UpdateSaleDetail(ctx context.Context, saleDetail *entity.SaleDetail) error
	GetSingleSaleDetail(ctx context.Context, saleDetailID string) (*entity.SaleDetail, error)
	GetPageSaleDetail(ctx context.Context, limit int, offset int) (*entity.SaleDetails, error)
}
