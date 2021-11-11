package repository_impl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/repository"
)

func NewCartRepository(database *mongo.Database, configuration config.Config) repository.CartRepository {
	return &cartRepositoryImpl{
		Collection: database.Collection("cart"),
		Configuration: configuration,
	}
}

type cartRepositoryImpl struct {
	Collection *mongo.Collection
	Configuration config.Config
}

func (repository *cartRepositoryImpl) FindByUserId(userId string) (carts []entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	filter := bson.M{"user_id": userId}
	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)
	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)
	for _, document := range documents {
		carts = append(carts, entity.Cart{
			Id:        document["_id"].(string),
			UserId:    document["user_id"].(string),
			ProductId: document["product_id"].(string),
			Quantity:  document["quantity"].(int64),
			CreatedAt: document["created_at"].(int64),
			UpdatedAt: document["updated_at"].(int64),
		})
	}
	return carts
}

func (repository *cartRepositoryImpl) Insert(cart entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err :=repository.Collection.InsertOne(ctx, bson.D{
		{"_id", cart.Id},
		{"user_id", cart.UserId},
		{"product_id", cart.ProductId},
		{"quantity", cart.Quantity},
		{"created_at", cart.CreatedAt},
		{"updated_at", cart.UpdatedAt},
	})
	exception.PanicIfNeeded(err)
}

func (repository *cartRepositoryImpl) Update(cart entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id":cart.Id}
	option := bson.M{"$set": bson.M{
		"quantity":         	cart.Quantity,
		"updated_at":         	cart.UpdatedAt,
	}}
	_, err := repository.Collection.UpdateOne(ctx, filter, option)
	exception.PanicIfNeeded(err)
}

func (repository *cartRepositoryImpl) Delete(id string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx,bson.M{"_id":id})
	exception.PanicIfNeeded(err)
}



