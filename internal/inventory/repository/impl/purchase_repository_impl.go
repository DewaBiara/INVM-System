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

	err := u.db.WithContext(ctx).
		Where("id = ?", purchaseID).
		Preload("Items").
		First(&purchase).Error
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

func (d *PurchaseRepositoryImpl) DeletePurchase(ctx context.Context, purchaseID string) error {
	result := d.db.WithContext(ctx).
		Select("Purchase").
		Delete(&entity.Purchase{}, "id = ?", purchaseID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrDocumentNotFound
	}

	return nil
}
