package dto

import (
	"time"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type CreatePurchaseRequest struct {
	SupplierID uint      `json:"supplierid" validate:"required"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func (u *CreatePurchaseRequest) ToEntity() *entity.Purchase {
	return &entity.Purchase{
		SupplierID: u.SupplierID,
		TotalPrice: u.TotalPrice,
		Date:       u.Date,
	}
}

type UpdatePurchaseRequest struct {
	ID         uint      `json:"id" validate:"required"`
	SupplierID uint      `json:"supplierid" validate:"required"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func (u *UpdatePurchaseRequest) ToEntity() *entity.Purchase {
	return &entity.Purchase{
		SupplierID: u.SupplierID,
		TotalPrice: u.TotalPrice,
		Date:       u.Date,
	}
}

type GetSinglePurchaseResponse struct {
	ID         uint      `json:"id"`
	SupplierID uint      `json:"supplierid" validate:"required"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func NewGetSinglePurchaseResponse(purchase *entity.Purchase) *GetSinglePurchaseResponse {
	return &GetSinglePurchaseResponse{
		ID:         purchase.ID,
		SupplierID: purchase.SupplierID,
		TotalPrice: purchase.TotalPrice,
		Date:       purchase.Date,
	}
}

type GetPagePurchaseResponse struct {
	ID         uint      `json:"id"`
	SupplierID uint      `json:"supplierid" validate:"required"`
	TotalPrice int       `json:"totalprice"`
	Date       time.Time `json:"date"`
}

func NewGetPagePurchaseResponse(purchase *entity.Purchase) *GetPagePurchaseResponse {
	return &GetPagePurchaseResponse{
		ID:         purchase.ID,
		SupplierID: purchase.SupplierID,
		TotalPrice: purchase.TotalPrice,
		Date:       purchase.Date,
	}
}

type GetPagePurchasesResponse []GetPagePurchaseResponse

func NewGetPagePurchasesResponse(purchases *entity.Purchases) *GetPagePurchasesResponse {
	var getPagePurchases GetPagePurchasesResponse
	for _, purchases := range *purchases {
		getPagePurchases = append(getPagePurchases, *NewGetPagePurchaseResponse(&purchases))
	}
	return &getPagePurchases
}

type CreatePurchaseDetailRequest struct {
	PurchaseID uint   `json:"purchaseid"`
	ItemID     uint   `json:"itemid"`
	TotalItem  int    `json:"totalitem"`
	Price      int    `json:"price"`
	UserID     string `json:"userid"`
}

func (u *CreatePurchaseDetailRequest) ToEntity() *entity.PurchaseDetail {
	return &entity.PurchaseDetail{
		PurchaseID: u.PurchaseID,
		ItemID:     u.ItemID,
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type UpdatePurchaseDetailRequest struct {
	ID         uint   `json:"id" validate:"required"`
	PurchaseID uint   `json:"purchaseid"`
	ItemID     uint   `json:"itemid"`
	TotalItem  int    `json:"totalitem"`
	Price      int    `json:"price"`
	UserID     string `json:"userid"`
}

func (u *UpdatePurchaseDetailRequest) ToEntity() *entity.PurchaseDetail {
	return &entity.PurchaseDetail{
		PurchaseID: u.PurchaseID,
		ItemID:     u.ItemID,
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type GetSinglePurchaseDetailResponse struct {
	ID         uint   `json:"id"`
	PurchaseID uint   `json:"purchaseid"`
	ItemID     uint   `json:"itemid"`
	TotalItem  int    `json:"totalitem"`
	Price      int    `json:"price"`
	UserID     string `json:"userid"`
}

func NewGetSinglePurchaseDetailResponse(purchaseDetail *entity.PurchaseDetail) *GetSinglePurchaseDetailResponse {
	return &GetSinglePurchaseDetailResponse{
		ID:         purchaseDetail.ID,
		PurchaseID: purchaseDetail.PurchaseID,
		ItemID:     purchaseDetail.ItemID,
		TotalItem:  purchaseDetail.TotalItem,
		Price:      purchaseDetail.Price,
		UserID:     purchaseDetail.UserID,
	}
}

type GetPagePurchaseDetailResponse struct {
	ID         uint   `json:"id"`
	PurchaseID uint   `json:"purchaseid"`
	ItemID     uint   `json:"itemid"`
	TotalItem  int    `json:"totalitem"`
	Price      int    `json:"price"`
	UserID     string `json:"userid"`
}

func NewGetPagePurchaseDetailResponse(purchaseDetail *entity.PurchaseDetail) *GetPagePurchaseDetailResponse {
	return &GetPagePurchaseDetailResponse{
		ID:         purchaseDetail.ID,
		PurchaseID: purchaseDetail.PurchaseID,
		ItemID:     purchaseDetail.ItemID,
		TotalItem:  purchaseDetail.TotalItem,
		Price:      purchaseDetail.Price,
		UserID:     purchaseDetail.UserID,
	}
}

type GetPagePurchaseDetailsResponse []GetPagePurchaseDetailResponse

func NewGetPagePurchaseDetailsResponse(purchaseDetails *entity.PurchaseDetails) *GetPagePurchaseDetailsResponse {
	var getPagePurchaseDetails GetPagePurchaseDetailsResponse
	for _, purchaseDetails := range *purchaseDetails {
		getPagePurchaseDetails = append(getPagePurchaseDetails, *NewGetPagePurchaseDetailResponse(&purchaseDetails))
	}
	return &getPagePurchaseDetails
}
