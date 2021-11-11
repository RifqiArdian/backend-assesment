package helper

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"online-store/model"
)

func Ok (c *fiber.Ctx,data interface{})error{
	return c.Status(http.StatusOK).JSON(model.HttpResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func BadRequest (c *fiber.Ctx,err error)error{
	return c.Status(http.StatusBadRequest).JSON(model.HttpResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD_REQUEST",
		Data:   err.Error(),
	})
}

func Unauthorized(c *fiber.Ctx,err error) error {
	return c.Status(http.StatusUnauthorized).JSON(model.HttpResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   err.Error(),
	})
}

func MethodNotAllowed(c *fiber.Ctx,err error) error {
	return c.Status(http.StatusMethodNotAllowed).JSON(model.HttpResponse{
		Code:   http.StatusMethodNotAllowed,
		Status: "METHOD_NOT_ALLOWED",
		Data:   err.Error(),
	})
}

func InternalServerError(c *fiber.Ctx,err error)error{
	return c.Status(http.StatusInternalServerError).JSON(model.HttpResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}