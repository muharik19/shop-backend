package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muharik19/shop-backend/constant"
	"github.com/muharik19/shop-backend/middleware"
	"github.com/muharik19/shop-backend/models"
	logger "github.com/muharik19/shop-backend/pkg/logging"
	"github.com/muharik19/shop-backend/services"
)

type OrderController struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create a order
// @Description Create a order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.OrderRegister true "Order"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 404 {object} models.Response
// @Router /orders [post]
func (controller *OrderController) CreateOrder(c *gin.Context) {
	v, ok := c.Get(constant.GIN_KEY)
	if !ok {
		c.JSON(401, models.Response{
			Code:         http.StatusUnauthorized,
			ResponseCode: constant.FAILED_AUTHORIZED,
			ResponseDesc: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	var orderRegister models.OrderRegister
	if err := c.ShouldBindJSON(&orderRegister); err != nil {
		middleware.Response(c, orderRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	order := models.Order{
		ShopID:    orderRegister.ShopID,
		Payment:   orderRegister.Payment,
		UserID:    v.(*models.UserClaims).ID,
		CreatedBy: v.(*models.UserClaims).Name,
	}

	var orderDetail []models.OrderDetail
	for _, e := range orderRegister.Products {
		detail := models.OrderDetail{
			ProductID: e.ProductID,
			Qty:       e.Qty,
		}
		orderDetail = append(orderDetail, detail)
	}

	response, err := controller.OrderService.CreateOrder(&order, &orderDetail)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, orderRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, orderRegister, *response)
}
