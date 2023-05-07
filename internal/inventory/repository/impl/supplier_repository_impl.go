package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/DewaBiara/INVM-System/pkg/utils"
	"gorm.io/gorm"
)

type SupplierRepositoryImpl struct {
	db *gorm.DB
}

func NewSupplierRepositoryImpl(db *gorm.DB) repository.SupplierRepository {
	supplierRepository := &SupplierRepositoryImpl{
		db: db,
	}

	return supplierRepository
}

func (u *SupplierRepositoryImpl) CreateSupplier(ctx context.Context, supplier *entity.Supplier) error {
	err := u.db.WithContext(ctx).Create(supplier).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *SupplierRepositoryImpl) GetSingleSupplier(ctx context.Context, supplierID string) (*entity.Supplier, error) {
	var supplier entity.Supplier
	err := u.db.WithContext(ctx).Select([]string{"id", "name", "address", "telp"}).Where("id = ?", supplierID).First(&supplier).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &supplier, nil
}

func (u *SupplierRepositoryImpl) GetPageSupplier(ctx context.Context, limit int, offset int) (*entity.Suppliers, error) {
	var suppliers entity.Suppliers
	err := u.db.WithContext(ctx).
		Select([]string{"id", "name", "address", "telp"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&suppliers).Error
	if err != nil {
		return nil, err
	}

	if len(suppliers) == 0 {
		return nil, utils.ErrUserNotFound
	}

	return &suppliers, nil
}

func (u *SupplierRepositoryImpl) UpdateSupplier(ctx context.Context, supplier *entity.Supplier) error {
	result := u.db.WithContext(ctx).Model(&entity.Supplier{}).Where("id = ?", supplier.ID).Updates(supplier)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrUsernameAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

func (d *SupplierRepositoryImpl) DeleteSupplier(ctx context.Context, supplierID string) error {
	result := d.db.WithContext(ctx).
		Select("Supplier").
		Delete(&entity.Supplier{}, "id = ?", supplierID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrDocumentNotFound
	}

	return nil
}
