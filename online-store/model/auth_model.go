package model

type LoginRequest struct {
	Email		string `json:"email"`
	Password	string `json:"password"`
}

type RegisterRequest struct {
	Name 		string `json:"name"`
	Email		string `json:"email"`
	Password	string `json:"password"`
	Address		string `json:"address"`
}

type GetUserResponse struct {
	Id			string `json:"id"`
	Name		string `json:"name"`
	Email		string `json:"email"`
	Balance		int64 `json:"balance"`
	Token		string `json:"token"`
	Address		string `json:"address"`
	CreatedAt	int64 `json:"created_at"`
	UpdatedAt	int64 `json:"updated_at"`
}
