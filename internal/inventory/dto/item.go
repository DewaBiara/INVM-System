package dto

import (
	"github.com/DewaBiara/INVM-System/pkg/entity"
	"gorm.io/gorm"
)

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
}

func (u *CreateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		Category:    u.Category,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
		CreatedBy:   u.CreatedBy,
	}
}

type UpdateItemRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	UpdatedBy   string `json:"updatedby"`
}

func (u *UpdateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		Category:    u.Category,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
		UpdatedBy:   u.UpdatedBy,
	}
}

type GetSingleItemResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
	UpdatedBy   string `json:"updatedby"`
}

func NewGetSingleItemResponse(item *entity.Item) *GetSingleItemResponse {
	return &GetSingleItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Category:    item.Category,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
	}
}

type GetPageItemResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
	UpdatedBy   string `json:"updatedby"`
}

func NewGetPageItemResponse(item *entity.Item) *GetPageItemResponse {
	return &GetPageItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Category:    item.Category,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
	}
}

type GetPageItemsResponse []GetPageItemResponse

func NewGetPageItemsResponse(items *entity.Items) *GetPageItemsResponse {
	var getPageItems GetPageItemsResponse
	for _, items := range *items {
		getPageItems = append(getPageItems, *NewGetPageItemResponse(&items))
	}
	return &getPageItems
}

type ItemRequest struct {
	ItemID uint `json:"itemid"`
}

type ItemRequests []ItemRequest

func (r ItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Model: gorm.Model{
			ID: r.ItemID,
		},
	}
}

func (r *ItemRequests) ToEntity() *entity.Items {
	items := entity.Items{}
	for _, item := range *r {
		items = append(items, *item.ToEntity())
	}

	return &items
}

type ItemResponse struct {
	ItemID   uint   `json:"itemid"`
	Name     string `json:"name"`
	Category string `json:"catagory"`
}

func NewItemResponse(item *entity.Item) *ItemResponse {
	return &ItemResponse{
		ItemID:   item.ID,
		Name:     item.Name,
		Category: item.Category,
	}
}

type ItemResponses []ItemResponse

func NewItemResponses(items *entity.Items) *ItemResponses {
	var itemResponses ItemResponses
	for _, items := range *items {
		itemResponses = append(itemResponses, *NewItemResponse(&items))
	}
	return &itemResponses
}
