package impl

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
	"github.com/google/uuid"
)

type (
	ItemServiceImpl struct {
		itemRepository repository.ItemRepository
	}
)

func NewItemServiceImpl(itemRepository repository.ItemRepository) service.ItemService {
	return &ItemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (u *ItemServiceImpl) CreateItem(ctx context.Context, item *dto.CreateItemRequest) error {

	itemEntity := item.ToEntity()
	itemEntity.ID = uint(uuid.New().ID())

	err := u.itemRepository.CreateItem(ctx, itemEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *ItemServiceImpl) GetSingleItem(ctx context.Context, itemID string) (*dto.GetSingleItemResponse, error) {
	item, err := d.itemRepository.GetSingleItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	var itemResponse = dto.NewGetSingleItemResponse(item)

	return itemResponse, nil
}

func (u *ItemServiceImpl) GetPageItem(ctx context.Context, page int, limit int) (*dto.GetPageItemsResponse, error) {
	offset := (page - 1) * limit

	items, err := u.itemRepository.GetPageItem(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageItemsResponse(items), nil
}

func (u *ItemServiceImpl) UpdateItem(ctx context.Context, itemID uint, updateItem *dto.UpdateItemRequest) error {
	item := updateItem.ToEntity()
	item.ID = itemID

	return u.itemRepository.UpdateItem(ctx, item)
}

func (d *ItemServiceImpl) DeleteItem(ctx context.Context, itemID string) error {

	return d.itemRepository.DeleteItem(ctx, itemID)
}
