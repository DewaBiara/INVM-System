package entity

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Category    string `gorm:"type:varchar(255);not null"`
	Price       int
	Stock       int
	Description string    `gorm:"type:varchar(255)"`
	Purchases   Purchases `gorm:"many2many:Item_Purchase;"`
	Sales       Sales     `gorm:"many2many:Item_Sale;"`
	CreatedBy   string    `gorm:"type:varchar(255)"`
	UpdatedBy   string    `gorm:"type:varchar(255)"`
}

type Items []Item
type Supplier struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Address   string
	Telp      string
	CreatedBy string `gorm:"type:varchar(255)"`
	UpdatedBy string `gorm:"type:varchar(255)"`
	Purchases Purchases
}

type Suppliers []Supplier
type Purchase struct {
	gorm.Model
	SupplierID uint
	TotalPrice int
	Date       time.Time
	Items      *Items `gorm:"many2many:Item_Purchase;"`
	TotalItem  int
	Price      int
	UserID     string
}

type Purchases []Purchase
type Sale struct {
	gorm.Model
	TotalPrice int
	Date       time.Time
	Items      *Items `gorm:"many2many:Item_Sale;"`
	TotalItem  int
	Price      int
	UserID     string
}

type Sales []Sale
