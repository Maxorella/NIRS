package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	intern "github.com/Maxorella/NIRS/internal"
	delivery "github.com/Maxorella/NIRS/internal/delivery/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
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
	Myhandler := delivery.NewClientAuthHandler(logger, repo)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(CORSMiddleware)
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)

	r.HandleFunc("/order_bad", Myhandler.GetOrderByIdInefficient).Methods(http.MethodGet)
	r.HandleFunc("/order_good", Myhandler.GetOrderById).Methods(http.MethodGet)
	r.HandleFunc("/user", Myhandler.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/product", Myhandler.GetProductByName).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:              ":8100",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info(fmt.Sprintf("Start server on %s\n", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	sig := <-signalCh
	logger.Info(fmt.Sprintf("Received signal: %v\n", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server shutdown failed: %s\n", err))
	}
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}
