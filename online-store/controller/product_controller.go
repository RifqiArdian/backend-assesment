package controller

import (
	"github.com/gofiber/fiber/v2"
	"online-store/config"
	"online-store/helper"
	"online-store/middleware"
	"online-store/service"
)

type ProductController struct {
	ProductService service.ProductService
	Configuration config.Config
	AuthMiddleware middleware.AuthMiddleware
}

func NewProductController(adminActivityService *service.ProductService, configuration config.Config, authMiddleware *middleware.AuthMiddleware) ProductController {
	return ProductController{
		ProductService: *adminActivityService,
		Configuration: configuration,
		AuthMiddleware: *authMiddleware,
	}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Get("product", controller.AuthMiddleware.CheckToken, controller.Index)
}

func (controller *ProductController) Index(c *fiber.Ctx) error {
	helper.LogRequest(c)
	responses := controller.ProductService.Get()
	return helper.Ok(c,responses)
}
