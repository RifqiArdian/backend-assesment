package model

type GetProductResponse struct {
	Id			string `json:"id"`
	Name		string `json:"name"`
	Category	string `json:"category"`
	Image		string `json:"image"`
	Price		int64 `json:"price"`
	Stock		int64 `json:"stock"`
	CreatedAt	int64 `json:"created_at"`
	UpdatedAt	int64 `json:"updated_at"`
}
