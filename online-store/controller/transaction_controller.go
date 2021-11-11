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

type TransactionController struct {
	TransactionService service.TransactionService
	Configuration config.Config
	AuthMiddleware middleware.AuthMiddleware
}

func NewTransactionController(adminActivityService *service.TransactionService, configuration config.Config, authMiddleware *middleware.AuthMiddleware) TransactionController {
	return TransactionController{
		TransactionService: *adminActivityService,
		Configuration: configuration,
		AuthMiddleware: *authMiddleware,
	}
}

func (controller *TransactionController) Route(app *fiber.App) {
	app.Get("transaction", controller.AuthMiddleware.CheckToken, controller.Index)
	app.Post("transaction", controller.AuthMiddleware.CheckToken, controller.Create)
}

func (controller *TransactionController) Index(c *fiber.Ctx) error {
	helper.LogRequest(c)
	userId := c.Cookies("user-id")
	responses := controller.TransactionService.Get(userId)
	return helper.Ok(c,responses)
}

func (controller *TransactionController)Create(c *fiber.Ctx) error {
	helper.LogRequest(c)
	var request model.InsertTransactionRequest
	request.UserId = c.Cookies("user-id")
	err := c.BodyParser(&request)
	if err!=nil{
		panic(exception.ValidationError{
			Status: http.StatusBadRequest,
			Message:  "Invalid data input",
		})
	}
	request.UserId = c.Cookies("user-id")
	controller.TransactionService.Insert(request)
	return helper.Ok(c,nil)
}