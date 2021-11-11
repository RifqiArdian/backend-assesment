package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"net/http"
	"online-store/config"
	"online-store/exception"
	"online-store/model"
	"online-store/repository"
	"online-store/repository/impl"
)

func GetProductRepository()(repository repository.ProductRepository){
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)
	return repository_impl.NewProductRepository(database,configuration)
}

func CheckProductId(value interface{}) error {
	repo := GetProductRepository()
	id, _ := value.(string)
	_, err := repo.FindById(id)
	if err != nil {
		return errors.New("Invalid product id")
	}
	return nil
}

func ValidateCreateCart(request model.InsertCartRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ProductId, validation.Required, validation.By(CheckProductId)),
		validation.Field(&request.Quantity, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}
}

func ValidateUpdateCart(request model.UpdateCartRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Quantity, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}
}

