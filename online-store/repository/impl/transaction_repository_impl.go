package repository_impl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/repository"
)

func NewTransactionRepository(database *mongo.Database, configuration config.Config) repository.TransactionRepository {
	return &transactionRepositoryImpl{
		Collection: database.Collection("transaction"),
		Configuration: configuration,
	}
}

type transactionRepositoryImpl struct {
	Collection *mongo.Collection
	Configuration config.Config
}

func (repository *transactionRepositoryImpl) Insert(transaction entity.Transaction) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err :=repository.Collection.InsertOne(ctx, bson.D{
		{"_id", transaction.Id},
		{"user_id", transaction.UserId},
		{"product_id", transaction.ProductId},
		{"quantity", transaction.Quantity},
		{"price", transaction.Price},
		{"total_price", transaction.TotalPrice},
		{"address", transaction.Address},
		{"created_at", transaction.CreatedAt},
		{"updated_at", transaction.UpdatedAt},
	})
	exception.PanicIfNeeded(err)
}

func (repository *transactionRepositoryImpl) FindByUserId(userId string) (transactions []entity.Transaction) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"user_id": userId}
	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		transactions = append(transactions, entity.Transaction{
			Id:         document["_id"].(string),
			UserId:     document["user_id"].(string),
			ProductId:  document["product_id"].(string),
			Quantity:   document["quantity"].(int64),
			Price:      document["price"].(int64),
			TotalPrice: document["total_price"].(int64),
			Address:    document["address"].(string),
			CreatedAt:  document["created_at"].(int64),
			UpdatedAt:  document["updated_at"].(int64),
		})
	}
	return transactions
}



