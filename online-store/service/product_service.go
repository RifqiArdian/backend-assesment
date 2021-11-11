package service

import "online-store/model"

type ProductService interface {
	Get()(responses []model.GetProductResponse)
}
