package services

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/muharik19/shop-backend/constant"
	"github.com/muharik19/shop-backend/models"
	"github.com/muharik19/shop-backend/pkg/util"
	"github.com/muharik19/shop-backend/repositories"
	"golang.org/x/sync/errgroup"
)

type OrderService struct {
	OrderRepository   repositories.OrderRepository
	ProductRepository repositories.ProductRepository
	ShopRepository    repositories.ShopRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, productRepository repositories.ProductRepository, shopRepository repositories.ShopRepository) *OrderService {
	return &OrderService{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
		ShopRepository:    shopRepository,
	}
}

func (service *OrderService) CreateOrder(order *models.Order, orderDetail *[]models.OrderDetail) (res *models.Response, err error) {
	// Simulate concurrent order insertions
	var wg sync.WaitGroup
	var mu sync.Mutex

	errs, _ := errgroup.WithContext(context.Background())

	var arrOrderDetail []models.OrderDetail
	var amountOrder float64
	invoice := util.GenerateInvoice()

	_, errShop := service.ShopRepository.GetShopById(order.ShopID)
	if errShop != nil {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: "Shop Not Found",
		}, nil
	}

	for _, v := range *orderDetail {
		product, errProduct := service.ProductRepository.GetProductById(v.ProductID)
		if errProduct != nil {
			return &models.Response{
				Code:         http.StatusNotFound,
				ResponseCode: constant.FAILED_NOT_FOUND,
				ResponseDesc: "Product Not Found",
			}, nil
		}

		if product.Stock == 0 {
			continue
		}

		amountDetail := v.Qty * product.Price
		detail := models.OrderDetail{
			Invoice:   invoice,
			ProductID: v.ProductID,
			Qty:       v.Qty,
			Price:     product.Price,
			Amount:    amountDetail,
			CreatedBy: v.CreatedBy,
		}
		amountOrder += amountDetail
		arrOrderDetail = append(arrOrderDetail, detail)
	}

	order.Invoice = invoice
	order.Amount = amountOrder

	wg.Add(1)
	errs.Go(func() error {
		errGo := service.OrderRepository.TransactionOrder(order, &arrOrderDetail, &wg, &mu)
		if errGo != nil {
			return fmt.Errorf("error in go routine, TransactionOrder, %+v", errGo)
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
		ResponseDesc: "Order created successfully",
	}, nil
}
