package repository_impl

import (
	"github.com/google/uuid"
	"online-store/config"
	"online-store/entity"
	"testing"
	"time"
)

func TestProduct_Add(t *testing.T) {
	configuration := config.New("../../.env")
	database := config.NewMongoDatabase(configuration)
	productRepository := NewProductRepository(database,configuration)

	product := entity.Product{
		Id:        uuid.New().String(),
		Name:      "Teh 1 Kg",
		Category:  "Minuman",
		Image:     "https://asset.kompas.com/crops/49IuzTlG_FJi0-smKo0BWkTSLFY=/50x34:450x300/750x500/data/photo/2019/12/30/5e099171e7576.jpg",
		Price:     10000,
		Stock:     100,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	productRepository.Insert(product)
}
