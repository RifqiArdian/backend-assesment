package model

type GetCartResponse struct {
	Id			string `json:"id"`
	Product		GetProductResponse `json:"product_id"`
	Quantity	int64 `json:"quantity"`
	CreatedAt	int64 `json:"created_at"`
	UpdatedAt	int64 `json:"updated_at"`
}

type InsertCartRequest struct {
	UserId		string `json:"user_id"`
	ProductId	string `json:"product_id"`
	Quantity	int64 `json:"quantity"`
}

type UpdateCartRequest struct {
	Id			string `json:"id"`
	Quantity	int64 `json:"quantity"`
}
