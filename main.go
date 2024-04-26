package main

import (
	"context"
	"fmt"
	"log"

	"4d63.com/tz"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/muharik19/shop-backend/controllers"
	"gorm.io/gorm"

	_ "github.com/muharik19/shop-backend/docs"
	"github.com/muharik19/shop-backend/pkg/database"
	"github.com/muharik19/shop-backend/pkg/util"
	"github.com/muharik19/shop-backend/repositories"
	"github.com/muharik19/shop-backend/routes"
	"github.com/muharik19/shop-backend/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Shop API
// @description This is a service reporting shopping.
// @version 1
// @host localhost:9090
// @BasePath /v1
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
// @contact.name Muharik
// @contact.email ahmadmuharik@gmail.com
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := database.InitGorm(ctx)
	if err != nil {
		panic(err)
	}

	// Cron job release stock if not yet payment
	err = InitCron(db)
	if err != nil {
		panic(err)
	}

	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)

	// Repositories
	userRepository := repositories.NewUserRepository(db)
	shopRepository := repositories.NewShopRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	// Services
	userService := services.NewUserService(*userRepository)
	shopService := services.NewShopService(*shopRepository, *orderRepository)
	productService := services.NewProductService(*productRepository, *orderRepository)
	authService := services.NewAuthService(*userRepository)
	orderService := services.NewOrderService(*orderRepository, *productRepository, *shopRepository)

	// Controllers
	userController := controllers.NewUserController(*userService)
	shopController := controllers.NewShopController(*shopService)
	productController := controllers.NewProductController(*productService)
	authController := controllers.NewAuthController(*authService)
	orderController := controllers.NewOrderController(*orderService)

	router := routes.NewRouter(*userController, *shopController, *productController, *authController, *orderController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%s", util.Getenv("HTTP_PORT"))
	router.Run(port)
}

func InitCron(db *gorm.DB) error {
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)
	productService := services.NewProductService(*productRepository, *orderRepository)

	var loc = "Asia/Jakarta"
	var wib, _ = tz.LoadLocation(loc)
	scheduler := gocron.NewScheduler(wib)

	_, err := scheduler.Cron("*/01 * * * *").Do(productService.ReleaseStockNotPayment)
	if err != nil {
		return err
	}

	go scheduler.StartAsync()

	return nil
}
