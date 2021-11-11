package repository

import "online-store/entity"

type CartRepository interface {
	FindByUserId(userId string)(carts []entity.Cart)
	Insert(cart entity.Cart)()
	Update(cart entity.Cart)()
	Delete(id string)
}
