package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type CreatePurchaseRequest struct {
	SupplierID uint          `json:"supplierid" validate:"required"`
	TotalPrice int           `json:"totalprice"`
	Date       CustomTime    `json:"date"`
	Items      *ItemRequests `json:"items"`
	TotalItem  int           `json:"totalitem"`
	Price      int           `json:"price"`
	UserID     string        `json:"userid"`
}

func (u *CreatePurchaseRequest) ToEntity() *entity.Purchase {
	return &entity.Purchase{
		SupplierID: u.SupplierID,
		TotalPrice: u.TotalPrice,
		Date:       u.Date.Time,
		Items:      u.Items.ToEntity(),
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type UpdatePurchaseRequest struct {
	ID         uint          `json:"id" validate:"required"`
	SupplierID uint          `json:"supplierid" validate:"required"`
	TotalPrice int           `json:"totalprice"`
	Date       CustomTime    `json:"date"`
	Items      *ItemRequests `json:"items"`
	TotalItem  int           `json:"totalitem"`
	Price      int           `json:"price"`
	UserID     string        `json:"userid"`
}

func (u *UpdatePurchaseRequest) ToEntity() *entity.Purchase {
	return &entity.Purchase{
		SupplierID: u.SupplierID,
		TotalPrice: u.TotalPrice,
		Date:       u.Date.Time,
		Items:      u.Items.ToEntity(),
		TotalItem:  u.TotalItem,
		Price:      u.Price,
		UserID:     u.UserID,
	}
}

type GetSinglePurchaseResponse struct {
	ID         uint           `json:"id"`
	SupplierID uint           `json:"supplierid" validate:"required"`
	TotalPrice int            `json:"totalprice"`
	Date       time.Time      `json:"date"`
	Items      *ItemResponses `json:"items"`
	TotalItem  int            `json:"totalitem"`
	Price      int            `json:"price"`
	UserID     string         `json:"userid"`
}

func NewGetSinglePurchaseResponse(purchase *entity.Purchase) *GetSinglePurchaseResponse {
	return &GetSinglePurchaseResponse{
		ID:         purchase.ID,
		SupplierID: purchase.SupplierID,
		TotalPrice: purchase.TotalPrice,
		Date:       purchase.Date,
		Items:      NewItemResponses(purchase.Items),
		TotalItem:  purchase.TotalItem,
		Price:      purchase.Price,
		UserID:     purchase.UserID,
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
