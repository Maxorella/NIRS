package internal

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type Repo struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewRepository(db *sql.DB, logger *zap.Logger) *Repo {
	return &Repo{db: db, logger: logger}
}

func (r *Repo) CreateUser(ctx context.Context) error {
	userName := generateRandomString(10)
	email := fmt.Sprintf("%s@example.com", generateRandomString(8))
	dateBirth := generateRandomDate()
	registrationDate := time.Now()

	query := `INSERT INTO "user" (user_name, email, date_birth, registration_date)
			  VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, userName, email, dateBirth, registrationDate)
	if err != nil {
		return err
	}

	return nil
}

func generateRandomString(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateRandomDate() time.Time {
	min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (r *Repo) CreateOrder(ctx context.Context, maxId int) error {
	_, err := r.db.Exec(`
		INSERT INTO "order" (user_id, order_date, address)
		VALUES ($1, $2, $3)
		RETURNING order_id
	`, rand.Intn(maxId), generateRandomDate(), generateRandomString(40))

	if err != nil {
		return fmt.Errorf("error inserting order: %w", err)
	}
	return nil
}

func (r *Repo) CreateProduct(ctx context.Context) error {
	_, err := r.db.Exec(`
		INSERT INTO product (product_name, price, stock)
		VALUES ($1, $2, $3)
		RETURNING product_id
	`, generateRandomString(20), rand.Intn(10000), rand.Intn(1000))

	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	return nil
}

func (r *Repo) CreateOrderDetail(ctx context.Context, maxId int) error {
	_, err := r.db.Exec(`
		INSERT INTO order_detail (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
		RETURNING order_detail_id
	`, rand.Intn(maxId), rand.Intn(maxId), rand.Intn(100), rand.Intn(10000))
	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	return nil
}
