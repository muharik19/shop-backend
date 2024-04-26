package repositories

import (
	"fmt"
	"sync"

	"github.com/muharik19/shop-backend/models"
	"gorm.io/gorm"
)

type ShopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{
		DB: db,
	}
}

func (repository *ShopRepository) AddShop(shop *models.Shop) error {
	return repository.DB.Create(shop).Error
}

func (repository *ShopRepository) GetShopById(id string) (models.Shop, error) {
	var shop models.Shop
	if err := repository.DB.Where(&models.Shop{ID: id}).Where("deleted_at isnull").First(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func (repository *ShopRepository) GetShopByEmail(email string) (models.Shop, error) {
	var shop models.Shop
	if err := repository.DB.Where(&models.Shop{Email: email}).Where("deleted_at isnull").First(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func (repository *ShopRepository) GetShopByPhoneNumber(phoneNumber string) (models.Shop, error) {
	var shop models.Shop
	if err := repository.DB.Where(&models.Shop{PhoneNumber: phoneNumber}).Where("deleted_at isnull").First(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func (repository *ShopRepository) AddWarehouse(warehouse *models.Warehouse) error {
	return repository.DB.Create(warehouse).Error
}

func (repository *ShopRepository) UpdateWarehouse(warehouse *models.Warehouse) error {
	return repository.DB.Updates(warehouse).Error
}

func (repository *ShopRepository) GetWarehouseByCode(code string) (models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := repository.DB.Where(&models.Warehouse{Code: code}).Where("deleted_at isnull").First(&warehouse).Error; err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

func (repository *ShopRepository) GetWarehouseByShopId(shopID, name string) (models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := repository.DB.Where(&models.Warehouse{ShopID: shopID, Name: name}).Where("deleted_at isnull").First(&warehouse).Error; err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

func (repository *ShopRepository) AddWarehouseStock(warehouseStock *[]models.WarehouseStock, wg *sync.WaitGroup, mu *sync.Mutex) error {
	defer wg.Done()

	// Acquire a lock to ensure safe concurrent access
	mu.Lock()
	defer mu.Unlock()

	// Begin a transaction
	tx := repository.DB.Begin()

	// Perform database operations within the transaction (use 'tx' from this point)
	if err := tx.Create(warehouseStock).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating order detail, %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction, %v", err)
	}

	return nil
}
