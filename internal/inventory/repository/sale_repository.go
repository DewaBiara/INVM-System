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
}
