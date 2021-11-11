package service

import "online-store/model"

type TransactionService interface {
	Get(userId string)(responses []model.GetTransactionResponse)
	Insert(request model.InsertTransactionRequest)()
}
