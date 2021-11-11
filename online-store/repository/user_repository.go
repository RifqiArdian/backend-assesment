package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"online-store/entity"
)

type UserRepository interface {
	FindByEmail(email string)(user entity.User, err error)
	FindById(id string)(user entity.User, err error)
	FindByToken(token string)(user entity.User, err error)
	Insert(user entity.User)()
	DeleteToken(token string)
	UpdateBalance(user entity.User, sessionContext mongo.SessionContext)(err error)
	UpdateToken(userId string,token string)
}
