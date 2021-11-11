package entity

type Transaction struct {
	Id				string `bson:"_id"`
	UserId			string `bson:"user_id"`
	ProductId		string `bson:"product_id"`
	Quantity		int64 `bson:"quantity"`
	Price			int64 `bson:"price"`
	TotalPrice		int64 `bson:"total_price"`
	Address			string `bson:"address"`
	CreatedAt		int64 `bson:"created_at"`
	UpdatedAt		int64 `bson:"updated_at"`
}
