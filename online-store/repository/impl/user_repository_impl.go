package repository_impl

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"online-store/config"
	"online-store/entity"
	"online-store/exception"
	"online-store/repository"
)

func NewUserRepository(database *mongo.Database, configuration config.Config) repository.UserRepository {
	return &userRepositoryImpl{
		Collection: database.Collection("user"),
		Configuration: configuration,
	}
}

type userRepositoryImpl struct {
	Collection *mongo.Collection
	Configuration config.Config
}

func (repository *userRepositoryImpl) UpdateBalance(user entity.User, sessionContext mongo.SessionContext)(err error) {
	filter := bson.M{"_id":user.Id}
	option := bson.M{"$set": bson.M{
		"updated_at":	user.UpdatedAt,
		"balance":	user.Balance,
	}}
	_, err = repository.Collection.UpdateOne(sessionContext, filter, option)

	return err
}

func (repository *userRepositoryImpl) UpdateToken(userId string,token string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id":userId}
	option := bson.M{"$set": bson.M{
		"token":	token,
	}}
	_, err := repository.Collection.UpdateOne(ctx, filter, option)
	exception.PanicIfNeeded(err)
}

func (repository *userRepositoryImpl) FindById(id string) (user entity.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	filter := bson.M{"_id":id}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	return user,err
}

func (repository *userRepositoryImpl) DeleteToken(token string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"token":token}
	option := bson.M{"$set": bson.M{
		"token":         		"",
	}}
	_, err := repository.Collection.UpdateOne(ctx, filter, option)
	exception.PanicIfNeeded(err)
}

func (repository *userRepositoryImpl) FindByToken(token string) (user entity.User,err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	filter := bson.M{"token":token}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	return user,err
}

func (repository *userRepositoryImpl) FindByEmail(email string) (user entity.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	filter := bson.M{"email":email}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	return user,err
}

func (repository *userRepositoryImpl) Insert(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err :=repository.Collection.InsertOne(ctx, bson.D{
		{"_id", user.Id},
		{"name", user.Name},
		{"email", user.Email},
		{"password", user.Password},
		{"balance", user.Balance},
		{"token", user.Token},
		{"address", user.Address},
		{"created_at", user.CreatedAt},
		{"updated_at", user.UpdatedAt},
	})
	exception.PanicIfNeeded(err)
}

