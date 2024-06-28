package models

type OrderDetail struct {
	OrderDetailID int64 `json:"order_detail_id"`
	OrderID       int64 `json:"order_id"`
	ProductID     int64 `json:"product_id"`
	Quantity      int64 `json:"quantity"`
	Price         int64 `json:"price"`
}
