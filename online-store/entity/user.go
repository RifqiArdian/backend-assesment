package entity

type User struct {
	Id			string `bson:"_id"`
	Name		string `bson:"name"`
	Email		string `bson:"email"`
	Password	string `bson:"password"`
	Balance		int64 `bson:"balance"`
	Token		string `bson:"token"`
	Address		string `bson:"address"`
	CreatedAt	int64 `bson:"created_at"`
	UpdatedAt	int64 `bson:"updated_at"`
}
