package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"online-store/exception"
	"online-store/model"
)

func ValidateCreateTransaction(request model.InsertTransactionRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ProductId, validation.Required,validation.By(CheckProductId)),
		validation.Field(&request.Quantity, validation.Required),
		validation.Field(&request.Address, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}
}
