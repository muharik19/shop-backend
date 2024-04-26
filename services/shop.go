package services

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/muharik19/shop-backend/constant"
	"github.com/muharik19/shop-backend/pkg/util"
	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
	"github.com/muharik19/shop-backend/models"
	"github.com/muharik19/shop-backend/repositories"
)

type ShopService struct {
	ShopRepository  repositories.ShopRepository
	OrderRepository repositories.OrderRepository
}

func NewShopService(shopRepository repositories.ShopRepository, orderRepository repositories.OrderRepository) *ShopService {
	return &ShopService{
		ShopRepository:  shopRepository,
		OrderRepository: orderRepository,
	}
}

func (service *ShopService) CreateShop(shop *models.Shop) (res *models.Response, err error) {
	_, err = service.ShopRepository.GetShopByEmail(shop.Email)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Email already exist",
		}, nil
	}

	_, err = service.ShopRepository.GetShopByPhoneNumber(shop.PhoneNumber)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Phone Number already exist",
		}, nil
	}

	shop.ID = uuid.New().String()
	err = service.ShopRepository.AddShop(shop)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Shop created successfully",
	}, nil
}

func (service *ShopService) CreateWarehouse(warehouse *models.Warehouse) (res *models.Response, err error) {
	_, err = service.ShopRepository.GetShopById(warehouse.ShopID)
	if err != nil {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Shop Not Found",
		}, nil
	}

	_, err = service.ShopRepository.GetWarehouseByShopId(warehouse.ShopID, warehouse.Name)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Name already exist",
		}, nil

	}

	warehouse.Code = util.GenerateWarehouse()
	err = service.ShopRepository.AddWarehouse(warehouse)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Warehouse created successfully",
	}, nil
}

func (service *ShopService) ActiveWarehouse(warehouse *models.Warehouse) (res *models.Response, err error) {
	data, err := service.ShopRepository.GetWarehouseByCode(warehouse.Code)
	if err != nil {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	if *data.Active && *warehouse.Active {
		return &models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Sorry, the warehouse is already enable",
		}, nil
	}

	if !*data.Active && !*warehouse.Active {
		return &models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Sorry, the warehouse is already disable",
		}, nil
	}

	updatedBy := "system"
	warehouse = &models.Warehouse{
		Code:      data.Code,
		Active:    warehouse.Active,
		UpdatedBy: &updatedBy,
	}
	err = service.ShopRepository.UpdateWarehouse(warehouse)
	if err != nil {
		return nil, err
	}

	message := "Warehouse disable successfully"
	if *warehouse.Active {
		message = "Warehouse enable successfully"
	}

	return &models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: message,
	}, nil
}

func (service *ShopService) CreateWarehouseStock(warehouseStock *models.WarehouseStock) (res *models.Response, err error) {
	// Simulate concurrent order insertions
	var wg sync.WaitGroup
	var mu sync.Mutex

	errs, _ := errgroup.WithContext(context.Background())

	data, err := service.ShopRepository.GetWarehouseByCode(warehouseStock.Code)
	if err != nil {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Code Not Found",
		}, nil
	}

	if !*data.Active {
		return &models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Sorry, the warehouse is already disable",
		}, nil
	}

	invoice, err := service.OrderRepository.GetOrderByInvoice(warehouseStock.Invoice)
	if err != nil {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Invoice Not Found",
		}, nil
	}

	if !*invoice.Payment {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Your not yet payment",
		}, nil
	}

	invoiceDetail, count, err := service.OrderRepository.GetOrderDetailByInvoice(invoice.Invoice)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Detail Order Not Found",
		}, nil
	}

	var arrWarehouseStock []models.WarehouseStock
	for _, v := range invoiceDetail {
		detail := models.WarehouseStock{
			Code:      data.Code,
			Invoice:   v.Invoice,
			ProductID: v.ProductID,
			Qty:       v.Qty,
			Price:     v.Price,
			Amount:    v.Amount,
			CreatedBy: "system",
		}
		arrWarehouseStock = append(arrWarehouseStock, detail)
	}

	wg.Add(1)
	errs.Go(func() error {
		errGo := service.ShopRepository.AddWarehouseStock(&arrWarehouseStock, &wg, &mu)
		if errGo != nil {
			return fmt.Errorf("error in go routine, AddWarehouseStock, %+v", errGo)
		}
		return nil
	})

	// Wait for completion and return the first error (if any)
	if err := errs.Wait(); err != nil {
		return nil, err
	}

	wg.Wait()

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Warehouse Stock created successfully",
	}, nil
}
