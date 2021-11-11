package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"online-store/entity"
)

type ProductRepository interface {
	SessionTransaction()(*mongo.Client, *options.TransactionOptions)
	FindAll()(products []entity.Product)
	FindById(id string)(product entity.Product,err error)
	UpdateStock(product entity.Product, sessionContext mongo.SessionContext)(err error)
	Insert(product entity.Product)()
}
