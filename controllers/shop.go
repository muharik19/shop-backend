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

type ShopController struct {
	ShopService services.ShopService
}

func NewShopController(shopService services.ShopService) *ShopController {
	return &ShopController{
		ShopService: shopService,
	}
}

// CreateShop godoc
// @Summary Create a shop
// @Description Create a shop
// @Tags shops
// @Accept  json
// @Produce  json
// @Param shop body models.ShopRegister true "Shop"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 302 {object} models.Response
// @Router /shops [post]
func (controller *ShopController) CreateShop(c *gin.Context) {
	var shopRegister models.ShopRegister
	if err := c.ShouldBindJSON(&shopRegister); err != nil {
		middleware.Response(c, shopRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	shop := models.Shop{
		Email:       shopRegister.Email,
		Name:        shopRegister.Name,
		PhoneNumber: shopRegister.PhoneNumber,
		Address:     shopRegister.Address,
		CreatedBy:   shopRegister.Name,
	}

	response, err := controller.ShopService.CreateShop(&shop)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, shopRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, shopRegister, *response)
}

// CreateWarehouse godoc
// @Summary Create a warehouse
// @Description Create a warehouse
// @Tags warehouses
// @Accept  json
// @Produce  json
// @Param warehouse body models.WarehouseRegister true "Warehouse"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 404 {object} models.Response
// @Response 302 {object} models.Response
// @Router /warehouses [post]
func (controller *ShopController) CreateWarehouse(c *gin.Context) {
	var warehouseRegister models.WarehouseRegister
	if err := c.ShouldBindJSON(&warehouseRegister); err != nil {
		middleware.Response(c, warehouseRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	warehouse := models.Warehouse{
		ShopID:  warehouseRegister.ShopID,
		Name:    warehouseRegister.Name,
		Address: warehouseRegister.Address,
	}

	response, err := controller.ShopService.CreateWarehouse(&warehouse)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, warehouseRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, warehouseRegister, *response)
}

// ActiveWarehouse godoc
// @Summary Active a warehouse
// @Description Active a warehouse
// @Tags warehouses
// @Accept  json
// @Produce  json
// @Param warehouse body models.WarehouseActiveByCode true "Warehouse"
// @Response 200 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 404 {object} models.Response
// @Router /warehouses/:code [patch]
func (controller *ShopController) ActiveWarehouse(c *gin.Context) {
	var active models.WarehouseActiveByCode
	if err := c.ShouldBindJSON(&active); err != nil {
		middleware.Response(c, active, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	warehouse := models.Warehouse{
		Code:   c.Param("code"),
		Active: active.Active,
	}

	response, err := controller.ShopService.ActiveWarehouse(&warehouse)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, active, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, active, *response)
}

// CreateWarehouseStock godoc
// @Summary Create a warehouseStock
// @Description Create a warehouseStock
// @Tags warehouses
// @Accept  json
// @Produce  json
// @Param warehouse body models.WarehouseStockRegister true "Warehouse"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 404 {object} models.Response
// @Router /warehouses/stock [post]
func (controller *ShopController) CreateWarehouseStock(c *gin.Context) {
	var warehouseStockRegister models.WarehouseStockRegister
	if err := c.ShouldBindJSON(&warehouseStockRegister); err != nil {
		middleware.Response(c, warehouseStockRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	warehouse := models.WarehouseStock{
		Code:    warehouseStockRegister.Code,
		Invoice: warehouseStockRegister.Invoice,
	}

	response, err := controller.ShopService.CreateWarehouseStock(&warehouse)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, warehouseStockRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, warehouseStockRegister, *response)
}
