package helper

import (
	"github.com/gofiber/fiber/v2"
	"online-store/logger"
)

func LogRequest(c *fiber.Ctx) {
	Ip := c.Get("X-Real-IP")
	if Ip == "" {
		Ip = c.IP()
	}

	logger.Info().Msg("method=" + c.Method() + " endpoint=" + c.OriginalURL() + " ip=" + Ip)
}
