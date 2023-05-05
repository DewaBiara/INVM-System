package repository

import (
	"context"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type SupplierRepository interface {
	CreateSupplier(ctx context.Context, supplier *entity.Supplier) error
	UpdateSupplier(ctx context.Context, supplier *entity.Supplier) error
	GetSingleSupplier(ctx context.Context, supplierID string) (*entity.Supplier, error)
	GetPageSupplier(ctx context.Context, limit int, offset int) (*entity.Suppliers, error)
}
