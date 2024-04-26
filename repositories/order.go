package repositories

import (
	"fmt"
	"sync"

	"github.com/muharik19/shop-backend/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (repository *OrderRepository) TransactionOrder(order *models.Order, orderDetail *[]models.OrderDetail, wg *sync.WaitGroup, mu *sync.Mutex) error {
	defer wg.Done()

	// Acquire a lock to ensure safe concurrent access
	mu.Lock()
	defer mu.Unlock()

	// Begin a transaction
	tx := repository.DB.Begin()

	// Perform database operations within the transaction (use 'tx' from this point)
	if err := tx.Create(order).Error; err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback()
		return fmt.Errorf("error creating order, %v", err)
	}

	if err := tx.Create(orderDetail).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error creating order detail, %v", err)
	}

	for _, v := range *orderDetail {
		if err := tx.Model(&models.Product{}).Where(&models.Product{ID: v.ProductID}).Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", v.Qty)}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error updating stock product, %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction, %v", err)
	}

	return nil
}

func (repository *OrderRepository) GetOrderByInvoice(invoice string) (models.Order, error) {
	var order models.Order
	if err := repository.DB.Where(&models.Order{Invoice: invoice}).Where("deleted_at isnull").First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (repository *OrderRepository) GetOrderDetailByInvoice(invoice string) ([]models.OrderDetail, int64, error) {
	var count int64
	var err error
	var orderDetail []models.OrderDetail

	result := repository.DB.Model(&models.OrderDetail{}).Where("deleted_at isnull").Where(&models.OrderDetail{Invoice: invoice}).Find(&orderDetail)
	if result.Error != nil {
		return nil, count, result.Error
	}

	err = result.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	return orderDetail, count, nil
}

func (repository *OrderRepository) GetOrderByFalse(payment bool) ([]models.Order, int64, error) {
	var count int64
	var err error
	var order []models.Order

	result := repository.DB.Where("deleted_at isnull").Where(models.Order{Payment: &payment}).Find(&order)
	if result.Error != nil {
		return nil, count, result.Error
	}

	err = result.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	return order, count, nil
}
