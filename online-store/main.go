package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"online-store/config"
	"online-store/controller"
	"online-store/exception"
	"online-store/logger"
	"online-store/middleware"
	"online-store/repository/impl"
	"online-store/service/impl"
	"os"
	"strconv"
	"time"
)

func main()  {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	cartRepository := repository_impl.NewCartRepository(database, configuration)
	productRepository := repository_impl.NewProductRepository(database, configuration)
	transactionRepository := repository_impl.NewTransactionRepository(database, configuration)
	userRepository := repository_impl.NewUserRepository(database, configuration)

	cartService := service_impl.NewCartService(&cartRepository,&productRepository,configuration)
	productService := service_impl.NewProductService(&productRepository,configuration)
	transactionService := service_impl.NewTransactionService(&transactionRepository,&productRepository, &userRepository,configuration)
	authService := service_impl.NewAuthService(&userRepository,configuration)

	authMiddleware := middleware.NewAuthMiddleware(&userRepository)

	cartController := controller.NewCartController(&cartService,configuration,&authMiddleware)
	productController := controller.NewProductController(&productService,configuration,&authMiddleware)
	transactionController := controller.NewTransactionController(&transactionService,configuration,&authMiddleware)
	authController := controller.NewAuthController(&authService,configuration,&authMiddleware)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		Header: "X-Trace-Id",
		Generator: func() string {
			logger.TraceId = strconv.FormatInt(time.Now().UnixNano(), 32)
			return logger.TraceId
		},
	}))

	cartController.Route(app)
	productController.Route(app)
	transactionController.Route(app)
	authController.Route(app)

	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}