package repository

import "online-store/entity"

type TransactionRepository interface {
	FindByUserId(userId string)(transactions []entity.Transaction)
	Insert(transaction entity.Transaction)()
}