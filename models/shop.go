package models

import (
	"time"
)

type Shop struct {
	ID          string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Email       string     `json:"email" gorm:"not null;type:varchar(100);index"`
	PhoneNumber string     `json:"phoneNumber" gorm:"not null;type:varchar(20);index"`
	Name        string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Address     string     `json:"address" gorm:"not null;type:varchar(250);index"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy   string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy   *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy   *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Shop) TableName() string {
	return "shop"
}

type Warehouse struct {
	Code      string     `json:"code" gorm:"primary_key;not null;type:varchar(100);index"`
	ShopID    string     `json:"shopId" gorm:"not null;type:varchar(100);index"`
	Name      string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Address   string     `json:"address" gorm:"not null;type:varchar(250);index"`
	Active    *bool      `json:"active" gorm:"not null;default:true"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Warehouse) TableName() string {
	return "warehouse"
}

type WarehouseStock struct {
	Code      string     `json:"code" gorm:"not null;type:varchar(100);index"`
	Invoice   string     `json:"invoice" gorm:"not null;type:varchar(100);index"`
	ProductID string     `json:"productId" gorm:"not null;type:varchar(100);index"`
	Qty       float64    `json:"qty" gorm:"index"`
	Price     float64    `json:"price" gorm:"index"`
	Amount    float64    `json:"amount" gorm:"index"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (WarehouseStock) TableName() string {
	return "warehouse_stock"
}

type (
	ShopRegister struct {
		Email       string `json:"email" binding:"required,email"`
		Name        string `json:"name" binding:"required,min=3"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}

	WarehouseRegister struct {
		ShopID  string `json:"shopId" binding:"required"`
		Name    string `json:"name" binding:"required,min=3"`
		Address string `json:"address" binding:"required"`
	}

	WarehouseActiveByCode struct {
		Active *bool `json:"active" binding:"required"`
	}

	WarehouseStockRegister struct {
		Code    string `json:"code" binding:"required"`
		Invoice string `json:"invoice" binding:"required"`
	}
)
