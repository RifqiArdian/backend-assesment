package service_impl

import (
	"online-store/config"
	"online-store/logger"
	"online-store/model"
	"online-store/repository"
	"online-store/service"
)

func NewProductService(productRepository *repository.ProductRepository, configuration config.Config) service.ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
		Configuration: configuration,
	}
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
	Configuration config.Config
}

func (service *productServiceImpl) Get() (responses []model.GetProductResponse) {
	products := service.ProductRepository.FindAll()
	for _, product := range products{
		responses = append(responses,model.GetProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			Category:  product.Category,
			Image:     product.Image,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	logger.Info().Interface("Product: ", nil).Msg("Get product success")
	return responses
}
