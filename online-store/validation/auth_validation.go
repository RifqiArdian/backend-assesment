package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"net/http"
	"online-store/exception"
	"online-store/model"
)

func ValidateLogin(request model.LoginRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}
}

func ValidateRegister(request model.RegisterRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Address, validation.Required),

	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}
}
