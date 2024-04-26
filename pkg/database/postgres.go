package database

import (
	"context"
	"log"

	"github.com/muharik19/shop-backend/models"
	"github.com/muharik19/shop-backend/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbGorm *gorm.DB

func InitGorm(ctx context.Context) (*gorm.DB, error) {
	var err error
	connectionString := util.Getenv("GORM_CONNECTION")
	DbGorm, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	DbGorm.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Shop{},
		&models.Warehouse{},
		&models.WarehouseStock{},
	)

	return DbGorm, nil
}
