package models

import "time"

type Order struct {
	Invoice   string     `json:"invoice" gorm:"primary_key;not null;type:varchar(100);index"`
	UserID    string     `json:"userId" gorm:"not null;type:varchar(100);index"`
	Amount    float64    `json:"amount" gorm:"index"`
	ShopID    string     `json:"shopId" gorm:"not null;type:varchar(100);index"`
	Payment   *bool      `json:"payment" gorm:"not null;index;default:false"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Order) TableName() string {
	return "order"
}

type OrderDetail struct {
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

func (OrderDetail) TableName() string {
	return "order_detail"
}

type (
	OrderRegister struct {
		ShopID   string         `json:"shopId" binding:"required"`
		Payment  *bool          `json:"payment" binding:"required"`
		Products []OrderProduct `json:"products" binding:"required"`
	}

	OrderProduct struct {
		ProductID string  `json:"productId" binding:"required"`
		Qty       float64 `json:"qty" binding:"required"`
	}
)
