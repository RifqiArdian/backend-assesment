package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"online-store/exception"
	"online-store/repository"
)

type AuthMiddleware interface {
	CheckToken(c *fiber.Ctx)error
}

func NewAuthMiddleware(userRepository *repository.UserRepository) AuthMiddleware {
	return &authMiddlewareImpl{
		UserRepository: *userRepository,
	}
}

type authMiddlewareImpl struct {
	UserRepository repository.UserRepository
}

func (middleware *authMiddlewareImpl) CheckToken(c *fiber.Ctx) error {
	token := c.Cookies("token")
	_,err := middleware.UserRepository.FindByToken(token)
	if err != nil{
		panic(exception.ValidationError{
			Message: "invalid token",
			Status:  http.StatusUnauthorized,
		})
	}
	return c.Next()
}


