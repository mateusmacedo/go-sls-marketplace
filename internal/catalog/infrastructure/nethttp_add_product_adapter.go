package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type AddProductRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type AddProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

var HttpError = map[error]int{
	domain.ErrInvalidProductID:          http.StatusBadRequest,
	domain.ErrInvalidProductName:        http.StatusBadRequest,
	domain.ErrInvalidProductDescription: http.StatusBadRequest,
	domain.ErrInvalidProductPrice:       http.StatusBadRequest,
	domain.ErrAlreadyExistsProduct:      http.StatusConflict,
	domain.ErrNotFoundProduct:           http.StatusNotFound,
	domain.ErrRepositoryProduct:         http.StatusInternalServerError,
}

type NetHTTPAddProductAdapter struct {
	service application.AddProductUseCase
}

func NewNetHTTPAddProductAdapter(service application.AddProductUseCase) *NetHTTPAddProductAdapter {
	return &NetHTTPAddProductAdapter{service: service}
}

func (a *NetHTTPAddProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	var req AddProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := a.service.Execute(application.AddProductInput{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		http.Error(w, err.Error(), HttpError[err])
		return
	}

	res := AddProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
