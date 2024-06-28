package main

import (
	"context"
	"database/sql"
	"fmt"
	intern "github.com/Maxorella/NIRS/internal"
	"github.com/Maxorella/NIRS/internal/config"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	_ = godotenv.Load()
	logger := zap.Must(zap.NewDevelopment())
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))

	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	repo := intern.NewRepository(db, logger)

	for i := 0; i < config.NumBase; i++ {
		err := repo.CreateUser(context.Background())
		if err != nil {
			logger.Info("error create user " + string(i))
			continue
		}
	}

	for i := 0; i < config.NumBase; i++ {
		err := repo.CreateProduct(context.Background())
		if err != nil {
			logger.Info("error create product " + string(i))
			continue
		}
	}

	for i := 0; i < config.NumBase; i++ {
		err := repo.CreateOrder(context.Background(), config.NumBase)
		if err != nil {
			logger.Info("error create order " + string(i))
			continue
		}
	}

	for i := 0; i < config.NumBase*2; i++ {
		err := repo.CreateOrderDetail(context.Background(), config.NumBase)
		if err != nil {
			logger.Info("error create order_detail " + string(i))
			continue
		}
	}
}
