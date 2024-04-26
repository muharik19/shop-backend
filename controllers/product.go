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

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.ProductRegister true "Product"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Router /products [post]
func (controller *ProductController) CreateProduct(c *gin.Context) {
	var productRegister models.ProductRegister
	if err := c.ShouldBindJSON(&productRegister); err != nil {
		middleware.Response(c, productRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	user := models.Product{
		Name:  productRegister.Name,
		Price: productRegister.Price,
		Stock: productRegister.Stock,
	}

	response, err := controller.ProductService.CreateProduct(&user)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, productRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, productRegister, *response)
}

// ListProduct godoc
// @Summary List a product
// @Description List a product
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param collection query []string false "string collection" collectionFormat(multi)
// @Response 200 {object} models.Response
// @Response 404 {object} models.Response
// @Response 500 {object} models.Response
// @Router /products [get]
func (controller *ProductController) GetProducts(c *gin.Context) {
	filter := c.Request.URL.Query()
	response, err := controller.ProductService.GetProducts(filter)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, filter, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, filter, *response)
}
