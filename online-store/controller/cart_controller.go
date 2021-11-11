package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"online-store/config"
	"online-store/exception"
	"online-store/helper"
	"online-store/middleware"
	"online-store/model"
	"online-store/service"
)

type CartController struct {
	CartService service.CartService
	Configuration config.Config
	AuthMiddleware middleware.AuthMiddleware
}

func NewCartController(adminActivityService *service.CartService, configuration config.Config, authMiddleware *middleware.AuthMiddleware) CartController {
	return CartController{
		CartService: *adminActivityService,
		Configuration: configuration,
		AuthMiddleware: *authMiddleware,
	}
}

func (controller *CartController) Route(app *fiber.App) {
	app.Get("cart", controller.AuthMiddleware.CheckToken, controller.Index)
	app.Post("cart", controller.AuthMiddleware.CheckToken, controller.Create)
	app.Put("cart/:id", controller.AuthMiddleware.CheckToken, controller.Update)
	app.Delete("cart/:id", controller.AuthMiddleware.CheckToken, controller.Delete)
}

func (controller *CartController) Index(c *fiber.Ctx) error {
	helper.LogRequest(c)
	userId := c.Cookies("user-id")
	carts := controller.CartService.Get(userId)
	return helper.Ok(c,carts)
}

func (controller *CartController) Create(c *fiber.Ctx) error {
	helper.LogRequest(c)
	var request model.InsertCartRequest
	request.UserId = c.Cookies("user-id")
	err := c.BodyParser(&request)
	if err!=nil{
		panic(exception.ValidationError{
			Status: http.StatusBadRequest,
			Message:  "Invalid data input",
		})
	}
	controller.CartService.Insert(request)
	return helper.Ok(c,nil)
}

func (controller *CartController) Update(c *fiber.Ctx) error {
	helper.LogRequest(c)
	var request model.UpdateCartRequest
	request.Id=c.Params("id")
	err := c.BodyParser(&request)
	if err!=nil{
		panic(exception.ValidationError{
			Status: http.StatusBadRequest,
			Message:  "Invalid data input",
		})
	}
	controller.CartService.Update(request)
	return helper.Ok(c,nil)
}

func (controller *CartController) Delete(c *fiber.Ctx) error {
	helper.LogRequest(c)
	id := c.Params("id")
	controller.CartService.Delete(id)
	return helper.Ok(c,nil)
}