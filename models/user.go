package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID          string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Email       string     `json:"email" gorm:"not null;type:varchar(100);index"`
	PhoneNumber string     `json:"phoneNumber" gorm:"not null;type:varchar(20);index"`
	Name        string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Password    string     `json:"password" gorm:"type:varchar(150);index"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy   string     `json:"createdBy" gorm:"not null;type:varchar(150)"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"default:null"`
	UpdatedBy   *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150);default:null"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"default:null"`
	DeletedBy   *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150);default:null"`
}

func (User) TableName() string {
	return "user"
}

type (
	UserRegister struct {
		Email       string `json:"email" binding:"required,email"`
		Name        string `json:"name" binding:"required,min=3"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	UserClaims struct {
		jwt.StandardClaims
		ID          string `json:"id"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
	}
)
