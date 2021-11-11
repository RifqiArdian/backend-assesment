package model

type GetTransactionResponse struct {
	Id				string `json:"id"`
	UserId			string `json:"user_id"`
	ProductId		GetProductResponse `json:"product_id"`
	Quantity		int64 `json:"quantity"`
	Price			int64 `json:"price"`
	TotalPrice		int64 `json:"total_price"`
	Address			string `json:"address"`
	CreatedAt		int64 `json:"created_at"`
	UpdatedAt		int64 `json:"updated_at"`
}

type InsertTransactionRequest struct {
	UserId			string `json:"user_id"`
	ProductId		string `json:"product_id"`
	Quantity		int64 `json:"quantity"`
	Address			string `json:"address"`
}
