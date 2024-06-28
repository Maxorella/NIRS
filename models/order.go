package models

import "time"

type Order struct {
	OrderID   int64     `json:"order_id"`
	UserID    int64     `json:"user_id"`
	Address   string    `json:"address"`
	OrderDate time.Time `json:"order_date"`
	Total     int64     `json:"total"`
}
