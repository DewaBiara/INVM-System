package repository

import (
	"context"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, item *entity.Item) error
	UpdateItem(ctx context.Context, item *entity.Item) error
	GetSingleItem(ctx context.Context, itemID string) (*entity.Item, error)
	GetPageItem(ctx context.Context, limit int, offset int) (*entity.Items, error)
}
