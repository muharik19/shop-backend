package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muharik19/shop-backend/controllers"
	"github.com/muharik19/shop-backend/middleware"
)

func NewRouter(userController controllers.UserController, shopController controllers.ShopController, productController controllers.ProductController, authController controllers.AuthController, orderController controllers.OrderController) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	baseRouter := router.Group("/v1")

	/* ======= SHOP AUTH PRODUCT ROUTER ======= */
	shopAuthProduct := baseRouter.Group("/products")
	shopAuthProduct.Use(middleware.AuthMiddleware())
	shopAuthProduct.GET("", productController.GetProducts)

	/* ======= SHOP AUTH ORDER ROUTER ======= */
	shopAuthOrder := baseRouter.Group("/orders")
	shopAuthOrder.Use(middleware.AuthMiddleware())
	shopAuthOrder.POST("", orderController.CreateOrder)

	/* ======= SHOP ROUTER USERS ======= */
	users := baseRouter.Group("/users")
	users.POST("", userController.CreateUser)

	/* ======= SHOP ROUTER USERS ======= */
	shops := baseRouter.Group("/shops")
	shops.POST("", shopController.CreateShop)

	/* ======= WAREHOUSE ROUTER USERS ======= */
	warehouses := baseRouter.Group("/warehouses")
	warehouses.POST("", shopController.CreateWarehouse)
	warehouses.PATCH("/:code", shopController.ActiveWarehouse)
	warehouses.POST("/stock", shopController.CreateWarehouseStock)

	/* ======= SHOP ROUTER PRODUCTS ======= */
	products := baseRouter.Group("/products")
	products.POST("", productController.CreateProduct)

	/* ======= SHOP ROUTER AUTH ======= */
	auth := baseRouter.Group("/auth")
	auth.POST("/login", authController.Login)

	return router
}
