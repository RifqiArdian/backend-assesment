package entity

type Product struct {
	Id			string `bson:"_id"`
	Name		string `bson:"name"`
	Category	string `bson:"category"`
	Image		string `bson:"image"`
	Price		int64 `bson:"price"`
	Stock		int64 `bson:"stock"`
	CreatedAt	int64 `bson:"created_at"`
	UpdatedAt	int64 `bson:"updated_at"`
}
