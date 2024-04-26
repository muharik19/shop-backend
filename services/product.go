package services

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/muharik19/shop-backend/constant"
	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
	"github.com/muharik19/shop-backend/models"
	logger "github.com/muharik19/shop-backend/pkg/logging"
	"github.com/muharik19/shop-backend/pkg/util"
	"github.com/muharik19/shop-backend/repositories"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
	OrderRepository   repositories.OrderRepository
}

func NewProductService(productRepository repositories.ProductRepository, orderRepository repositories.OrderRepository) *ProductService {
	return &ProductService{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
	}
}

func (service *ProductService) CreateProduct(product *models.Product) (res *models.Response, err error) {
	product.ID = uuid.New().String()
	product.CreatedBy = "system"
	err = service.ProductRepository.AddProduct(product)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Product created successfully",
	}, nil
}

func (service *ProductService) GetProducts(filter map[string][]string) (res *models.Response, err error) {
	limit := constant.LIMIT_DEFAULT
	if len(filter["limit"]) > 0 {
		// string to int
		limit, err = strconv.Atoi(filter["limit"][0])
		if err != nil {
			return nil, err
		}
	}

	page := constant.PAGE_DEFAULT
	if len(filter["page"]) > 0 {
		// string to int
		page, err = strconv.Atoi(filter["page"][0])
		if err != nil {
			return nil, err
		}
	}

	var field string
	if len(filter["field"]) > 0 {
		field = filter["field"][0]
	}

	var sort string
	if len(filter["sort"]) > 0 {
		sort = filter["sort"][0]
	}

	m := make(map[string]string)
	if len(filter["search"]) > 0 {
		entries := strings.Split(strings.TrimSpace(filter["search"][0]), ",")
		for _, e := range entries {
			parts := strings.Split(e, "=")
			m[parts[0]] = parts[1]
		}
	}

	pagination := util.GeneratePaginationFromRequest(limit, page, field, sort)
	products, count, err := service.ProductRepository.GetProducts(pagination, m)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	data := models.ListProduct{
		Page:      page,
		Limit:     limit,
		Total:     int(count),
		TotalPage: int(math.Ceil(float64(count) / float64(limit))),
		Products:  products,
	}

	return &models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Product list successfully",
		ResponseData: data,
	}, nil
}

func (service *ProductService) ReleaseStockNotPayment() {
	logger.Info("================== Scheduler Release Stock Product Started ====================")

	// Simulate concurrent order insertions
	var wg sync.WaitGroup
	var mu sync.Mutex

	errs, _ := errgroup.WithContext(context.Background())

	order, count, err := service.OrderRepository.GetOrderByFalse(false)
	if err != nil {
		logger.Errf("ERR SCH RELEASE STOCK ORDER %+v", err)
	}

	if count == 0 {
		logger.Infof("COUNT SCH RELEASE STOCK ORDER %v", count)
	}

	deletedAt := time.Now()
	deletedBy := "system"

	for _, v := range order {
		orderDetail, count, err := service.OrderRepository.GetOrderDetailByInvoice(v.Invoice)
		if err != nil {
			logger.Errf("ERR SCH RELEASE STOCK ORDER DETAIL BY INVOICE %+v", err)
		}

		if count == 0 {
			logger.Infof("COUNT SCH RELEASE STOCK ORDER DETAIL BY INVOICE %v", count)
		}

		removeOrder := &models.Order{
			Invoice:   v.Invoice,
			DeletedAt: &deletedAt,
			DeletedBy: &deletedBy,
		}

		for _, o := range orderDetail {
			wg.Add(1)
			product := &models.Product{
				ID:    o.ProductID,
				Stock: o.Qty,
			}

			removeOrderDetail := &models.OrderDetail{
				ProductID: o.ProductID,
				DeletedAt: &deletedAt,
				DeletedBy: &deletedBy,
			}

			errs.Go(func() error {
				errGo := service.ProductRepository.UpdateReleaseStock(product, removeOrder, removeOrderDetail, &wg, &mu)
				if errGo != nil {
					logger.Errf("ERR SCH RELEASE STOCK UPDATE PRODUCT IN GO ROUTINE %+v", err)
					return fmt.Errorf("error in go routine, TransactionOrder, %+v", errGo)
				}
				return nil
			})
		}
		wg.Wait()
	}
	logger.Infof("SUCCESS SCH RELEASE STOCK UPDATE PRODUCT TOTAL PAYMENT FALSE %d", count)

	logger.Info("================== Scheduler Release Stock Product Ended ====================")
}
