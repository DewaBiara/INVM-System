package dto

import (
	"time"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type CreateSaleRequest struct {
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func (u *CreateSaleRequest) ToEntity() *entity.Sale {
	return &entity.Sale{
		TotalPrice: u.TotalPrice,
		Date:       u.Date,
	}
}

type UpdateSaleRequest struct {
	ID         uint      `json:"id" validate:"required"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func (u *UpdateSaleRequest) ToEntity() *entity.Sale {
	return &entity.Sale{
		TotalPrice: u.TotalPrice,
		Date:       u.Date,
	}
}

type GetSingleSaleResponse struct {
	ID         uint      `json:"id"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func NewGetSingleSaleResponse(sale *entity.Sale) *GetSingleSaleResponse {
	return &GetSingleSaleResponse{
		ID:         sale.ID,
		TotalPrice: sale.TotalPrice,
		Date:       sale.Date,
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

type CreateSaleDetailRequest struct {
	SaleID    uint   `json:"saleid"`
	ItemID    uint   `json:"itemid"`
	TotalItem int    `json:"totalitem"`
	Price     int    `json:"price"`
	UserID    string `json:"userid"`
	UserRole  string `json:"userrole"`
}

func (u *CreateSaleDetailRequest) ToEntity() *entity.SaleDetail {
	return &entity.SaleDetail{
		SaleID:    u.SaleID,
		ItemID:    u.ItemID,
		TotalItem: u.TotalItem,
		Price:     u.Price,
		UserID:    u.UserID,
		UserRole:  u.UserRole,
	}
}

type UpdateSaleDetailRequest struct {
	ID        uint   `json:"id" validate:"required"`
	SaleID    uint   `json:"saleid"`
	ItemID    uint   `json:"itemid"`
	TotalItem int    `json:"totalitem"`
	Price     int    `json:"price"`
	UserID    string `json:"userid"`
	UserRole  string `json:"userrole"`
}

func (u *UpdateSaleDetailRequest) ToEntity() *entity.SaleDetail {
	return &entity.SaleDetail{
		SaleID:    u.SaleID,
		ItemID:    u.ItemID,
		TotalItem: u.TotalItem,
		Price:     u.Price,
		UserID:    u.UserID,
		UserRole:  u.UserRole,
	}
}

type GetSingleSaleDetailResponse struct {
	ID        uint   `json:"id"`
	SaleID    uint   `json:"saleid"`
	ItemID    uint   `json:"itemid"`
	TotalItem int    `json:"totalitem"`
	Price     int    `json:"price"`
	UserID    string `json:"userid"`
	UserRole  string `json:"userrole"`
}

func NewGetSingleSaleDetailResponse(saleDetail *entity.SaleDetail) *GetSingleSaleDetailResponse {
	return &GetSingleSaleDetailResponse{
		ID:        saleDetail.ID,
		SaleID:    saleDetail.SaleID,
		ItemID:    saleDetail.ItemID,
		TotalItem: saleDetail.TotalItem,
		Price:     saleDetail.Price,
		UserID:    saleDetail.UserID,
		UserRole:  saleDetail.UserRole,
	}
}

type GetPageSaleDetailResponse struct {
	ID        uint   `json:"id"`
	SaleID    uint   `json:"saleid"`
	ItemID    uint   `json:"itemid"`
	TotalItem int    `json:"totalitem"`
	Price     int    `json:"price"`
	UserID    string `json:"userid"`
	UserRole  string `json:"userrole"`
}

func NewGetPageSaleDetailResponse(saleDetail *entity.SaleDetail) *GetPageSaleDetailResponse {
	return &GetPageSaleDetailResponse{
		ID:        saleDetail.ID,
		SaleID:    saleDetail.SaleID,
		ItemID:    saleDetail.ItemID,
		TotalItem: saleDetail.TotalItem,
		Price:     saleDetail.Price,
		UserID:    saleDetail.UserID,
		UserRole:  saleDetail.UserRole,
	}
}

type GetPageSaleDetailsResponse []GetPageSaleDetailResponse

func NewGetPageSaleDetailsResponse(saleDetails *entity.SaleDetails) *GetPageSaleDetailsResponse {
	var getPageSaleDetails GetPageSaleDetailsResponse
	for _, saleDetails := range *saleDetails {
		getPageSaleDetails = append(getPageSaleDetails, *NewGetPageSaleDetailResponse(&saleDetails))
	}
	return &getPageSaleDetails
}
