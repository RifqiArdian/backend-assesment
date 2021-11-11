package entity

type Cart struct {
	Id			string `bson:"_id"`
	UserId		string `bson:"user_id"`
	ProductId	string `bson:"product_id"`
	Quantity	int64 `bson:"quantity"`
	CreatedAt	int64 `bson:"created_at"`
	UpdatedAt	int64 `bson:"updated_at"`
}
