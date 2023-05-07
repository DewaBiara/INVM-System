package service

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
)

type ItemService interface {
	CreateItem(ctx context.Context, item *dto.CreateItemRequest) error
	UpdateItem(ctx context.Context, itemID uint, updateItem *dto.UpdateItemRequest) error
	GetSingleItem(ctx context.Context, itemID string) (*dto.GetSingleItemResponse, error)
	GetPageItem(ctx context.Context, limit int, offset int) (*dto.GetPageItemsResponse, error)
	DeleteItem(ctx context.Context, itemID string) error
}
