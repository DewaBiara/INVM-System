package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/DewaBiara/INVM-System/pkg/utils"
	"gorm.io/gorm"
)

type PurchaseRepositoryImpl struct {
	db *gorm.DB
}

func NewPurchaseRepositoryImpl(db *gorm.DB) repository.PurchaseRepository {
	purchaseRepository := &PurchaseRepositoryImpl{
		db: db,
	}

	return purchaseRepository
}

func (u *PurchaseRepositoryImpl) CreatePurchase(ctx context.Context, purchase *entity.Purchase) error {
	err := u.db.WithContext(ctx).Create(purchase).Error
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

func (u *PurchaseRepositoryImpl) GetSinglePurchase(ctx context.Context, purchaseID string) (*entity.Purchase, error) {
	var purchase entity.Purchase
	err := u.db.WithContext(ctx).Select([]string{"id", "supplierid", "totalprice", "date"}).Where("id = ?", purchaseID).First(&purchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &purchase, nil
}

func (u *PurchaseRepositoryImpl) GetPagePurchase(ctx context.Context, limit int, offset int) (*entity.Purchases, error) {
	var purchases entity.Purchases
	err := u.db.WithContext(ctx).
		Select([]string{"id", "supplierid", "totalprice", "date"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&purchases).Error
	if err != nil {
		return nil, err
	}

	if len(purchases) == 0 {
		return nil, utils.ErrUserNotFound
	}

	return &purchases, nil
}

func (u *PurchaseRepositoryImpl) UpdatePurchase(ctx context.Context, purchase *entity.Purchase) error {
	result := u.db.WithContext(ctx).Model(&entity.Purchase{}).Where("id = ?", purchase.ID).Updates(purchase)
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

func (u *PurchaseRepositoryImpl) CreatePurchaseDetail(ctx context.Context, purchaseDetail *entity.PurchaseDetail) error {
	err := u.db.WithContext(ctx).Create(purchaseDetail).Error
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

func (u *PurchaseRepositoryImpl) GetSinglePurchaseDetail(ctx context.Context, purchaseDetailID string) (*entity.PurchaseDetail, error) {
	var purchaseDetail entity.PurchaseDetail
	err := u.db.WithContext(ctx).Select([]string{"id", "purchaseid", "itemid", "totalitem", "price", "userid"}).Where("id = ?", purchaseDetailID).First(&purchaseDetail).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &purchaseDetail, nil
}

func (u *PurchaseRepositoryImpl) GetPagePurchaseDetail(ctx context.Context, limit int, offset int) (*entity.PurchaseDetails, error) {
	var purchaseDetails entity.PurchaseDetails
	err := u.db.WithContext(ctx).
		Select([]string{"id", "purchaseid", "itemid", "totalitem", "price", "userid"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&purchaseDetails).Error
	if err != nil {
		return nil, err
	}

	if len(purchaseDetails) == 0 {
		return nil, utils.ErrUserNotFound
	}

	return &purchaseDetails, nil
}

func (u *PurchaseRepositoryImpl) UpdatePurchaseDetail(ctx context.Context, purchaseDetail *entity.PurchaseDetail) error {
	result := u.db.WithContext(ctx).Model(&entity.PurchaseDetail{}).Where("id = ?", purchaseDetail.ID).Updates(purchaseDetail)
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
