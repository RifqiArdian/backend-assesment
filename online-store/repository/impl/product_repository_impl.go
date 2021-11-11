package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/repository"
)

func NewProductRepository(database *mongo.Database, configuration config.Config) repository.ProductRepository {
	return &productRepositoryImpl{
		Collection: database.Collection("product"),
		Configuration: configuration,
		Database: *database,
		DatabaseContext:    nil,
		DatabaseCancelFunc: nil,
	}
}

type productRepositoryImpl struct {
	Collection *mongo.Collection
	Configuration config.Config
	Database           mongo.Database
	DatabaseContext    context.Context
	DatabaseCancelFunc context.CancelFunc
}

func (repository *productRepositoryImpl) SessionTransaction() (*mongo.Client, *options.TransactionOptions) {
	dbClient := repository.Database.Client()
	return dbClient, config.MongoTransOption()
}

func (repository *productRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err :=repository.Collection.InsertOne(ctx, bson.D{
		{"_id", product.Id},
		{"name", product.Name},
		{"category", product.Category},
		{"image", product.Image},
		{"price", product.Price},
		{"stock", product.Stock},
		{"created_at", product.CreatedAt},
		{"updated_at", product.UpdatedAt},
	})
	exception.PanicIfNeeded(err)
}

func (repository *productRepositoryImpl) UpdateStock(product entity.Product, sessionContext mongo.SessionContext)(err error) {
	filter := bson.M{"_id":product.Id}
	option := bson.M{"$set": bson.M{
		"stock":	         	product.Stock,
		"updated_at":         	product.UpdatedAt,
	}}
	_, err = repository.Collection.UpdateOne(sessionContext, filter, option)

	return err
}

func (repository *productRepositoryImpl) FindAll() (products []entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		products = append(products, entity.Product{
			Id:       document["_id"].(string),
			Name:     document["name"].(string),
			Image:    document["image"].(string),
			Category: document["category"].(string),
			Price:    document["price"].(int64),
			Stock: 	  document["stock"].(int64),
			CreatedAt:document["created_at"].(int64),
			UpdatedAt:document["updated_at"].(int64),
		})
	}
	return products
}

func (repository *productRepositoryImpl) FindById(id string) (product entity.Product,err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	filter := bson.M{"_id":id}
	err = repository.Collection.FindOne(ctx, filter).Decode(&product)
	return product,err
}



