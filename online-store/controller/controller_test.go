package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"online-store/config"
	"online-store/middleware"
	"online-store/repository/impl"
	"online-store/service/impl"
)

func createTestApp() *fiber.App{
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	cartController.Route(app)
	productController.Route(app)
	authController.Route(app)
	transactionController.Route(app)
	return app
}

var configuration = config.New("../.env")
var database = config.NewMongoDatabase(configuration)

var cartRepository = repository_impl.NewCartRepository(database, configuration)
var productRepository = repository_impl.NewProductRepository(database, configuration)
var transactionRepository = repository_impl.NewTransactionRepository(database, configuration)
var userRepository = repository_impl.NewUserRepository(database, configuration)

var cartService = service_impl.NewCartService(&cartRepository,&productRepository,configuration)
var productService = service_impl.NewProductService(&productRepository,configuration)
var transactionService = service_impl.NewTransactionService(&transactionRepository,&productRepository, &userRepository,configuration)
var authService = service_impl.NewAuthService(&userRepository,configuration)

var authMiddleware = middleware.NewAuthMiddleware(&userRepository)

var cartController = NewCartController(&cartService,configuration,&authMiddleware)
var productController = NewProductController(&productService,configuration,&authMiddleware)
var transactionController = NewTransactionController(&transactionService,configuration,&authMiddleware)
var authController = NewAuthController(&authService,configuration,&authMiddleware)

var app = createTestApp()