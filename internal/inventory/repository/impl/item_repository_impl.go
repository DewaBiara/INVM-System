package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/DewaBiara/INVM-System/pkg/utils"
	"gorm.io/gorm"
)

type ItemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepositoryImpl(db *gorm.DB) repository.ItemRepository {
	itemRepository := &ItemRepositoryImpl{
		db: db,
	}

	return itemRepository
}

func (u *ItemRepositoryImpl) CreateItem(ctx context.Context, item *entity.Item) error {
	err := u.db.WithContext(ctx).Create(item).Error
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

func (u *ItemRepositoryImpl) GetSingleItem(ctx context.Context, itemID string) (*entity.Item, error) {
	var item entity.Item
	err := u.db.WithContext(ctx).Select([]string{"id", "name", "category", "price", "stock", "description", "createdby", "updatedby"}).
		Where("id = ?", itemID).First(&item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrItemNotFound
		}

		return nil, err
	}

	return &item, nil
}

func (u *ItemRepositoryImpl) GetPageItem(ctx context.Context, limit int, offset int) (*entity.Items, error) {
	var items entity.Items
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&items).Error
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, utils.ErrItemNotFound
	}

	return &items, nil
}

func (u *ItemRepositoryImpl) UpdateItem(ctx context.Context, item *entity.Item) error {
	result := u.db.WithContext(ctx).Model(&entity.Item{}).Where("id = ?", item.ID).Updates(item)
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
		return utils.ErrItemNotFound
	}

	return nil
}

func (d *ItemRepositoryImpl) DeleteItem(ctx context.Context, itemID string) error {
	result := d.db.WithContext(ctx).
		Select("Item").
		Delete(&entity.Item{}, "id = ?", itemID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrItemNotFound
	}

	return nil
}
