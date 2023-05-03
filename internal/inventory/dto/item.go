package dto

import "github.com/DewaBiara/INVM-System/pkg/entity"

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

func (u *CreateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		Category:    u.Category,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
	}
}

type UpdateItemRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

func (u *UpdateItemRequest) ToEntity() *entity.Item {
	return &entity.Item{
		Name:        u.Name,
		Category:    u.Category,
		Price:       u.Price,
		Stock:       u.Stock,
		Description: u.Description,
	}
}

type GetSingleItemResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

func NewGetSingleItemResponse(item *entity.Item) *GetSingleItemResponse {
	return &GetSingleItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Category:    item.Category,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
	}
}

type GetPageItemResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

func NewGetPageItemResponse(item *entity.Item) *GetPageItemResponse {
	return &GetPageItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Category:    item.Category,
		Price:       item.Price,
		Stock:       item.Stock,
		Description: item.Description,
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

type CreateSupplier struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"category" validate:"required"`
	Telp    string `json:"price"`
}

func (u *CreateSupplier) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:    u.Name,
		Address: u.Address,
		Telp:    u.Telp,
	}
}

type UpdateSupplier struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"category" validate:"required"`
	Telp    string `json:"price"`
}

func (u *UpdateSupplier) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:    u.Name,
		Address: u.Address,
		Telp:    u.Telp,
	}
}

type GetSingleSupplier struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"category" validate:"required"`
	Telp    string `json:"price"`
}

func NewGetSingleSupplier(supplier *entity.Supplier) *GetSingleSupplier {
	return &GetSingleSupplier{
		ID:      supplier.ID,
		Name:    supplier.Name,
		Address: supplier.Address,
		Telp:    supplier.Telp,
	}
}

type GetPageSupplier []GetSingleSupplier

func NewGetPageSupplier(suppliers *entity.Suppliers) *GetPageSupplier {
	var getPageSupplier GetPageSupplier
	for _, supplier := range *suppliers {
		getPageSupplier = append(getPageSupplier, *NewGetSingleSupplier(&supplier))
	}
	return &getPageSupplier
}
