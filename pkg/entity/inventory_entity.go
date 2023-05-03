package entity

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Category        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Price           int
	Stock           int
	Description     string           `gorm:"type:varchar(255)"`
	PurchaseDetails []PurchaseDetail `gorm:"many2many:Item_Purchase;"`
	SaleDetails     []SaleDetail     `gorm:"many2many:Item_Sale;"`
}

type Items []Item
type Supplier struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Address   string
	Telp      string
	Purchases []Purchase
}

type Suppliers []Supplier
type Purchase struct {
	gorm.Model
	SupplierID      string `gorm:"type:varchar(36)"`
	TotalPrice      int
	Date            time.Time
	PurchaseDetails []PurchaseDetail
}

type Purchases []Purchase
type PurchaseDetail struct {
	gorm.Model
	PurchaseID string `gorm:"type:varchar(36)"`
	ItemID     string `gorm:"type:varchar(36)"`
	TotalItem  int
	Price      int
	UserID     string
}

type PurchaseDetails []PurchaseDetail
type Sale struct {
	gorm.Model
	TotalPrice  int
	Date        time.Time
	SaleDetails []SaleDetail
}

type Sales []Sale
type SaleDetail struct {
	gorm.Model
	SaleID    string `gorm:"type:varchar(36)"`
	ItemID    string `gorm:"type:varchar(36)"`
	TotalItem int
	Price     int
	UserID    string
	UserRole  string
}

type SaleDetails []SaleDetail
