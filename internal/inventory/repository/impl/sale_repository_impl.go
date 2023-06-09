package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/DewaBiara/INVM-System/pkg/utils"
	"gorm.io/gorm"
)

type SaleRepositoryImpl struct {
	db *gorm.DB
}

func NewSaleRepositoryImpl(db *gorm.DB) repository.SaleRepository {
	saleRepository := &SaleRepositoryImpl{
		db: db,
	}

	return saleRepository
}

func (u *SaleRepositoryImpl) CreateSale(ctx context.Context, sale *entity.Sale) error {
	err := u.db.WithContext(ctx).Create(sale).Error
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

func (u *SaleRepositoryImpl) GetSingleSale(ctx context.Context, saleID string) (*entity.Sale, error) {
	var sale entity.Sale
	err := u.db.WithContext(ctx).
		Preload("Items").
		Where("id = ?", saleID).
		First(&sale).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrSaleNotFound
		}

		return nil, err
	}

	return &sale, nil
}

func (u *SaleRepositoryImpl) GetPageSale(ctx context.Context, limit int, offset int) (*entity.Sales, error) {
	var sales entity.Sales
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&sales).Error
	if err != nil {
		return nil, err
	}

	if len(sales) == 0 {
		return nil, utils.ErrSaleNotFound
	}

	return &sales, nil
}

func (u *SaleRepositoryImpl) UpdateSale(ctx context.Context, sale *entity.Sale) error {
	result := u.db.WithContext(ctx).Model(&entity.Sale{}).Where("id = ?", sale.ID).Updates(sale)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrSaleNotFound
	}

	return nil
}

func (d *SaleRepositoryImpl) DeleteSale(ctx context.Context, saleID string) error {
	result := d.db.WithContext(ctx).
		Select("Sale").
		Delete(&entity.Sale{}, "id = ?", saleID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrSaleNotFound
	}

	return nil
}
