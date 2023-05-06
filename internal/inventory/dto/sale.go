package dto

import (
	"time"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type CreateSaleRequest struct {
	TotalPrice int           `json:"totalprice"`
	Date       CustomTime    `json:"date"`
	Items      *ItemRequests `json:"items"`
	TotalItem  int           `json:"totalitem"`
	Price      int           `json:"price"`
	UserID     string        `json:"userid"`
}

func (u *CreateSaleRequest) ToEntity() *entity.Sale {
	return &entity.Sale{
		TotalPrice: u.TotalPrice,
		Date:       u.Date.Time,
		Items:      u.Items.ToEntity(),
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type UpdateSaleRequest struct {
	ID         uint          `json:"id" validate:"required"`
	TotalPrice int           `json:"totalprice"`
	Date       CustomTime    `json:"date"`
	Items      *ItemRequests `json:"items"`
	TotalItem  int           `json:"totalitem"`
	Price      int           `json:"price"`
	UserID     string        `json:"userid"`
}

func (u *UpdateSaleRequest) ToEntity() *entity.Sale {
	return &entity.Sale{
		TotalPrice: u.TotalPrice,
		Date:       u.Date.Time,
		Items:      u.Items.ToEntity(),
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type GetSingleSaleResponse struct {
	ID         uint           `json:"id"`
	TotalPrice int            `json:"totalprice"`
	Date       time.Time      `json:"date"`
	Items      *ItemResponses `json:"items"`
	TotalItem  int            `json:"totalitem"`
	Price      int            `json:"price"`
	UserID     string         `json:"userid"`
}

func NewGetSingleSaleResponse(sale *entity.Sale) *GetSingleSaleResponse {
	return &GetSingleSaleResponse{
		ID:         sale.ID,
		TotalPrice: sale.TotalPrice,
		Date:       sale.Date,
		Items:      NewItemResponses(sale.Items),
		TotalItem:  sale.TotalItem,
		Price:      sale.Price,
		UserID:     sale.UserID,
	}
}

type GetPageSaleResponse struct {
	ID         uint      `json:"id"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func NewGetPageSaleResponse(sale *entity.Sale) *GetPageSaleResponse {
	return &GetPageSaleResponse{
		ID:         sale.ID,
		TotalPrice: sale.TotalPrice,
		Date:       sale.Date,
	}
}

type GetPageSalesResponse []GetPageSaleResponse

func NewGetPageSalesResponse(sales *entity.Sales) *GetPageSalesResponse {
	var getPageSales GetPageSalesResponse
	for _, sales := range *sales {
		getPageSales = append(getPageSales, *NewGetPageSaleResponse(&sales))
	}
	return &getPageSales
}
