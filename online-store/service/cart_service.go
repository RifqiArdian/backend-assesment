package service

import (
	"online-store/model"
)

type CartService interface {
	Get(userId string)(responses []model.GetCartResponse)
	Insert(request model.InsertCartRequest)()
	Update(request model.UpdateCartRequest)()
	Delete(id string)()
}
