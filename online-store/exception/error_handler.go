package exception

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"online-store/helper"
	"online-store/logger"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	status, _ := err.(ValidationError)
	logger.Error().Msg(err.Error())

	if status.Status == http.StatusBadRequest {
		return helper.BadRequest(ctx,err)
	}

	if status.Status == http.StatusUnauthorized {
		return helper.Unauthorized(ctx,err)
	}

	if e, ok := err.(*fiber.Error); ok {
		if e.Code == http.StatusMethodNotAllowed {
			return helper.MethodNotAllowed(ctx,err)
		}
	}

	return helper.InternalServerError(ctx,err)
}
