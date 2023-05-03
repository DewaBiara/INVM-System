package mock

import (
	"context"

	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/stretchr/testify/mock"
)

type MockInventoryRepository struct {
	mock.Mock
}

func (m *MockInventoryRepository) CreateItem(ctx context.Context, item *entity.Item) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *MockInventoryRepository) GetSingleItem(ctx context.Context, itemID string) (*entity.Item, error) {
	args := m.Called(ctx, itemID)
	return args.Get(0).(*entity.Item), args.Error(1)
}

func (m *MockInventoryRepository) GetPageItem(ctx context.Context, limit int, offset int) (*entity.Items, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).(*entity.Items), args.Error(1)
}

func (m *MockInventoryRepository) UpdateItem(ctx context.Context, item *entity.Item) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}
