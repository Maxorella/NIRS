package delivery

import (
	"context"
	"encoding/json"
	"github.com/Maxorella/NIRS/internal"
	"github.com/Maxorella/NIRS/internal/config"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

type AuthClientHandler struct {
	logger *zap.Logger
	repo   *internal.Repo
}

func NewClientAuthHandler(logger *zap.Logger, repo *internal.Repo) *AuthClientHandler {
	return &AuthClientHandler{logger: logger, repo: repo}
}

func (h *AuthClientHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	user, err := h.repo.GetUserByID(context.Background(), rand.Intn(config.NumBase))
	if err != nil {
		http.Error(w, "Error retrieving user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthClientHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	order, err := h.repo.GetOrderByID(context.Background(), rand.Intn(config.NumBase))
	if err != nil {
		http.Error(w, "Error retrieving order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthClientHandler) GetOrderByIdInefficient(w http.ResponseWriter, r *http.Request) {
	order, err := h.repo.GetOrderByIDCount(context.Background(), rand.Intn(config.NumBase))
	if err != nil {
		http.Error(w, "Error retrieving order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthClientHandler) GetProductByName(w http.ResponseWriter, r *http.Request) {

	product, err := h.repo.GetProductByName(context.Background())
	if err != nil {
		http.Error(w, "Error retrieving product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthClientHandler) GetOrderDetailsByOrderID(w http.ResponseWriter, r *http.Request) {

	orderID := rand.Intn(config.NumBase)

	orderDetails, err := h.repo.GetOrderDetailsByOrderID(context.Background(), orderID)
	if err != nil {
		http.Error(w, "Error retrieving order details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orderDetails); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
