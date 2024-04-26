package repositories

import (
	"fmt"
	"sync"

	"github.com/muharik19/shop-backend/models"
	"github.com/muharik19/shop-backend/pkg/util"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repository *ProductRepository) AddProduct(product *models.Product) error {
	return repository.DB.Create(product).Error
}

func (repository *ProductRepository) GetProducts(pagination util.Pagination, where map[string]string) ([]models.Product, int64, error) {
	var count int64
	var err error
	var field, sort string
	var users []models.Product

	queryBuilder := repository.DB.Model(&models.Product{}).Where("deleted_at isnull")

	if where["id"] != "" {
		queryBuilder = queryBuilder.Where("id = ?", where["id"])
	}

	if where["name"] != "" {
		name := fmt.Sprintf("%%%s%%", where["name"])
		queryBuilder = queryBuilder.Where(`"name" ILIKE ?`, name)
	}

	if where["price"] != "" {
		queryBuilder = queryBuilder.Where("email = ?", where["email"])
	}

	if where["stock"] != "" {
		queryBuilder = queryBuilder.Where("partner_id = ?", where["partnerId"])
	}

	if where["shopId"] != "" {
		queryBuilder = queryBuilder.Where("shop_id = ?", where["shopId"])
	}

	if pagination.Field != "" {
		if pagination.Field == "id" {
			field = "INITCAP(id)"
		} else if pagination.Field == "name" {
			field = `INITCAP("name")`
		} else if pagination.Field == "shopId" {
			field = `INITCAP("shop_id")`
		} else if pagination.Field == "price" {
			field = "price"
		} else if pagination.Field == "stock" {
			field = "stock"
		} else {
			field = "created_at"
		}
	} else {
		field = "created_at"
	}

	if pagination.Sort != "" {
		sort = pagination.Sort
	} else {
		sort = "DESC"
	}

	err = queryBuilder.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	orderBy := fmt.Sprintf("%s %s", field, sort)
	limitBuilder := queryBuilder.Limit(int(pagination.Limit)).Offset(int(offset)).Order(orderBy)

	result := limitBuilder.Find(&users)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return users, count, nil
}

func (repository *ProductRepository) GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := repository.DB.Where(&models.Product{ID: id}).Where("deleted_at isnull").First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (repository *ProductRepository) UpdateReleaseStock(product *models.Product, order *models.Order, orderDetail *models.OrderDetail, wg *sync.WaitGroup, mu *sync.Mutex) error {
	defer wg.Done()

	// Acquire a lock to ensure safe concurrent access
	mu.Lock()
	defer mu.Unlock()

	// Begin a transaction
	tx := repository.DB.Begin()

	// Perform database operations within the transaction (use 'tx' from this point)
	if err := tx.Model(&models.Product{}).Where(&models.Product{ID: product.ID}).Updates(map[string]interface{}{"stock": gorm.Expr("stock + ?", product.Stock)}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating stock product, %v", err)
	}

	if err := tx.Model(&models.Order{}).Where(&models.Order{Invoice: order.Invoice}).Updates(models.Order{DeletedAt: order.DeletedAt, DeletedBy: order.DeletedBy}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete invoice order, %v", err)
	}

	if err := tx.Model(&models.OrderDetail{}).Where(&models.OrderDetail{ProductID: orderDetail.ProductID}).Updates(models.OrderDetail{DeletedAt: orderDetail.DeletedAt, DeletedBy: orderDetail.DeletedBy}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete invoice order, %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction, %v", err)
	}

	return nil
}
