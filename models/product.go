package models

import "time"

type Product struct {
	ID        string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Name      string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Price     float64    `json:"price" gorm:"index"`
	Stock     float64    `json:"stock" gorm:"index"`
	ShopID    string     `json:"shopId" gorm:"not null;type:varchar(100);index"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Product) TableName() string {
	return "product"
}

type (
	ProductRegister struct {
		Name   string  `json:"name" binding:"required,min=3"`
		Price  float64 `json:"price" binding:"required"`
		Stock  float64 `json:"stock" binding:"required"`
		ShopID string  `json:"shopId" binding:"required"`
	}

	ListProduct struct {
		Page      int       `json:"page"`
		Limit     int       `json:"limit"`
		Total     int       `json:"total"`
		TotalPage int       `json:"totalPage"`
		Products  []Product `json:"products"`
	}
)
