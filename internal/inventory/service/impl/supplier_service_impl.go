package impl

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
	"github.com/google/uuid"
)

type (
	SupplierServiceImpl struct {
		supplierRepository repository.SupplierRepository
	}
)

func NewSupplierServiceImpl(supplierRepository repository.SupplierRepository) service.SupplierService {
	return &SupplierServiceImpl{
		supplierRepository: supplierRepository,
	}
}

func (u *SupplierServiceImpl) CreateSupplier(ctx context.Context, supplier *dto.CreateSupplierRequest) error {

	supplierEntity := supplier.ToEntity()
	supplierEntity.ID = uint(uuid.New().ID())

	err := u.supplierRepository.CreateSupplier(ctx, supplierEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *SupplierServiceImpl) GetSingleSupplier(ctx context.Context, supplierID string) (*dto.GetSingleSupplierResponse, error) {
	supplier, err := d.supplierRepository.GetSingleSupplier(ctx, supplierID)
	if err != nil {
		return nil, err
	}

	var supplierResponse = dto.NewGetSingleSupplierResponse(supplier)

	return supplierResponse, nil
}

func (u *SupplierServiceImpl) GetPageSupplier(ctx context.Context, page int, limit int) (*dto.GetPageSuppliersResponse, error) {
	offset := (page - 1) * limit

	suppliers, err := u.supplierRepository.GetPageSupplier(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageSuppliersResponse(suppliers), nil
}

func (u *SupplierServiceImpl) UpdateSupplier(ctx context.Context, supplierID uint, updateSupplier *dto.UpdateSupplierRequest) error {
	supplier := updateSupplier.ToEntity()
	supplier.ID = supplierID

	return u.supplierRepository.UpdateSupplier(ctx, supplier)
}
