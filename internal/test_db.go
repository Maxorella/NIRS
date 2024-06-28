package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Maxorella/NIRS/models"
)

func (r *Repo) GetUserByID(ctx context.Context, userID int) (*models.Users, error) {
	query := `SELECT user_id, user_name, email, date_birth, registration_date
              FROM "user"
              WHERE user_id = $1`

	var user models.Users
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID, &user.UserName, &user.Email, &user.DateBirth, &user.RegDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with id %d", userID)
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return &user, nil
}

func (r *Repo) GetOrderByID(ctx context.Context, orderID int) (*models.Order, error) {
	query := `SELECT order_id, user_id, order_date, address, total
              FROM "order"
              WHERE order_id = $1`

	var order models.Order
	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&order.OrderID, &order.UserID, &order.OrderDate, &order.Address, &order.Total)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order found with id %d", orderID)
		}
		return nil, fmt.Errorf("error querying order: %w", err)
	}

	return &order, nil
}

func (r *Repo) GetProductByName(ctx context.Context) (*models.Product, error) {
	query := `SELECT product_id, product_name, price, stock
              FROM product
              WHERE product_name = $1`

	var product models.Product
	name := generateRandomString(20)
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&product.ProductID, &product.ProductName, &product.Price, &product.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no product found with name %d", name)
		}
		return nil, fmt.Errorf("error querying product: %w", err)
	}

	return &product, nil
}

func (r *Repo) GetOrderDetailsByOrderID(ctx context.Context, orderID int) ([]models.OrderDetail, error) {
	query := `SELECT order_detail_id, order_id, product_id, quantity, price
              FROM "order_detail"
              WHERE order_id = $1`

	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("error querying order details: %w", err)
	}
	defer rows.Close()

	var orderDetails []models.OrderDetail
	for rows.Next() {
		var orderDetail models.OrderDetail
		if err := rows.Scan(&orderDetail.OrderDetailID, &orderDetail.OrderID, &orderDetail.ProductID, &orderDetail.Quantity, &orderDetail.Price); err != nil {
			return nil, fmt.Errorf("error scanning order detail: %w", err)
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return orderDetails, nil
}

func (r *Repo) GetOrderByIDCount(ctx context.Context, orderID int) (*models.Order, error) {
	query := `SELECT order_id, user_id, order_date, address
              FROM "order"
              WHERE order_id = $1`

	var order models.Order
	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&order.OrderID, &order.UserID, &order.OrderDate, &order.Address)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order found with id %d", orderID)
		}
		return nil, fmt.Errorf("error querying order: %w", err)
	}

	orderDetails, err := r.GetOrderDetailsByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("error querying order details: %w", err)
	}

	total := int64(0)
	for _, detail := range orderDetails {
		total += detail.Quantity * detail.Price
	}
	order.Total = total

	return &order, nil
}
