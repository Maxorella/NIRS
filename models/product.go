package models

type Product struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}
