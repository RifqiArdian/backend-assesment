package service_impl

import (
	"github.com/google/uuid"
	"online-store/config"
	"online-store/entity"
	"online-store/logger"
	"online-store/model"
	"online-store/repository"
	"online-store/service"
	"online-store/validation"
	"time"
)

func NewCartService(cartRepository *repository.CartRepository, productRepository *repository.ProductRepository, configuration config.Config) service.CartService {
	return &cartServiceImpl{
		CartRepository: *cartRepository,
		Configuration: configuration,
		ProductRepository: *productRepository,
	}
}

type cartServiceImpl struct {
	CartRepository repository.CartRepository
	ProductRepository repository.ProductRepository
	Configuration config.Config
}

func (service *cartServiceImpl) Get(userId string) (responses []model.GetCartResponse) {
	carts := service.CartRepository.FindByUserId(userId)
	for _, cart := range carts{
		product,_ := service.ProductRepository.FindById(cart.ProductId)
		responses = append(responses,model.GetCartResponse{
			Id:	cart.Id,
			Product:   model.GetProductResponse{
				Id:        product.Id,
				Name:      product.Name,
				Category:  product.Category,
				Image:     product.Image,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			},
			Quantity:  cart.Quantity,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
		})
	}

	logger.Info().Interface("Cart: ", nil).Msg("Get cart success")
	return responses
}

func (service *cartServiceImpl) Insert(request model.InsertCartRequest) {
	validation.ValidateCreateCart(request)
	cart := entity.Cart{
		Id:        uuid.New().String(),
		UserId:    request.UserId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	service.CartRepository.Insert(cart)

	logger.Info().Interface("Cart: ", request).Msg("Add cart success")
}

func (service *cartServiceImpl) Update(request model.UpdateCartRequest) {
	validation.ValidateUpdateCart(request)
	cart := entity.Cart{
		Id:        request.Id,
		Quantity:  request.Quantity,
		UpdatedAt: time.Now().UnixNano(),
	}
	service.CartRepository.Update(cart)

	logger.Info().Interface("Cart: ", request).Msg("Update cart success")
}

func (service *cartServiceImpl) Delete(id string) {
	service.CartRepository.Delete(id)

	logger.Info().Interface("Cart: ", id).Msg("Delete cart success")
}



