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
	"time"
)

type AuthController struct {
	AuthService service.AuthService
	Configuration config.Config
	AuthMiddleware middleware.AuthMiddleware
}

func NewAuthController(adminActivityService *service.AuthService, configuration config.Config, authMiddleware *middleware.AuthMiddleware) AuthController {
	return AuthController{
		AuthService: *adminActivityService,
		Configuration: configuration,
		AuthMiddleware: *authMiddleware,
	}
}

func (controller *AuthController) Route(app *fiber.App) {
	app.Post("auth/register", controller.Register)
	app.Post("auth/login", controller.Login)
	app.Get("auth/profile", controller.AuthMiddleware.CheckToken, controller.Profile)
	app.Post("auth/logout", controller.AuthMiddleware.CheckToken, controller.Logout)
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	helper.LogRequest(c)
	var request model.RegisterRequest

	//parsing request to model
	err := c.BodyParser(&request)
	if err!=nil{
		panic(exception.ValidationError{
			Status: http.StatusBadRequest,
			Message:  "Invalid data input",
		})
	}

	//call register service
	response := controller.AuthService.Register(request)

	//create cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "user-id",
		Value:    response.Id,
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	return helper.Ok(c,response)
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	helper.LogRequest(c)
	var request model.LoginRequest

	//parsing request to model
	err := c.BodyParser(&request)
	if err!=nil{
		panic(exception.ValidationError{
			Status: http.StatusBadRequest,
			Message:  "Invalid data input",
		})
	}

	//call login service
	response := controller.AuthService.Login(request)

	//create cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "user-id",
		Value:    response.Id,
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	return helper.Ok(c,response)
}

func (controller *AuthController) Logout(c *fiber.Ctx) error {
	helper.LogRequest(c)
	token := c.Cookies("token")
	controller.AuthService.Logout(token)

	//delete cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "LOGOUT",
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "user-id",
		Value:    "LOGOUT",
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "strict",
	})
	return helper.Ok(c,nil)
}

func (controller *AuthController) Profile(c *fiber.Ctx) error {
	helper.LogRequest(c)
	token := c.Cookies("token")
	response := controller.AuthService.GetProfile(token)
	return helper.Ok(c,response)
}
