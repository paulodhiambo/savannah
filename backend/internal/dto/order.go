package dto

type CreateOrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	UserId    int `json:"user_id"`
}
