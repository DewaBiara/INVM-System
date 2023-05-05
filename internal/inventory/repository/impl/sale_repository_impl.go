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
	err := u.db.WithContext(ctx).Select([]string{"id", "totalprice", "date"}).Where("id = ?", saleID).First(&sale).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &sale, nil
}

func (u *SaleRepositoryImpl) GetPageSale(ctx context.Context, limit int, offset int) (*entity.Sales, error) {
	var sales entity.Sales
	err := u.db.WithContext(ctx).
		Select([]string{"id", "totalprice", "date"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&sales).Error
	if err != nil {
		return nil, err
	}

	if len(sales) == 0 {
		return nil, utils.ErrUserNotFound
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

func (u *SaleRepositoryImpl) CreateSaleDetail(ctx context.Context, saleDetail *entity.SaleDetail) error {
	err := u.db.WithContext(ctx).Create(saleDetail).Error
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

func (u *SaleRepositoryImpl) GetSingleSaleDetail(ctx context.Context, saleDetailID string) (*entity.SaleDetail, error) {
	var saleDetail entity.SaleDetail
	err := u.db.WithContext(ctx).Select([]string{"id", "saleid", "itemid", "totalitem", "price", "userid"}).Where("id = ?", saleDetailID).First(&saleDetail).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &saleDetail, nil
}

func (u *SaleRepositoryImpl) GetPageSaleDetail(ctx context.Context, limit int, offset int) (*entity.SaleDetails, error) {
	var saleDetails entity.SaleDetails
	err := u.db.WithContext(ctx).
		Select([]string{"id", "saleid", "itemid", "totalitem", "price", "userid"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&saleDetails).Error
	if err != nil {
		return nil, err
	}

	if len(saleDetails) == 0 {
		return nil, utils.ErrUserNotFound
	}

	return &saleDetails, nil
}

func (u *SaleRepositoryImpl) UpdateSaleDetail(ctx context.Context, saleDetail *entity.SaleDetail) error {
	result := u.db.WithContext(ctx).Model(&entity.SaleDetail{}).Where("id = ?", saleDetail.ID).Updates(saleDetail)
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
