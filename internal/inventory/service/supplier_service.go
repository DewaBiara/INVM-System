package service

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
)

type SupplierService interface {
	CreateSupplier(ctx context.Context, supplier *dto.CreateSupplierRequest) error
	UpdateSupplier(ctx context.Context, supplierID uint, updateSupplier *dto.UpdateSupplierRequest) error
	GetSingleSupplier(ctx context.Context, supplierID string) (*dto.GetSingleSupplierResponse, error)
	GetPageSupplier(ctx context.Context, limit int, offset int) (*dto.GetPageSuppliersResponse, error)
	DeleteSupplier(ctx context.Context, supplierID string) error
}
