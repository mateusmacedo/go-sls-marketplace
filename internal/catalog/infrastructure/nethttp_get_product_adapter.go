package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
)

type GetProductRequest struct {
	ID string `json:"id"`
}

type GetProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type NetHTTPGetProductAdapter struct {
	useCase application.GetProductsUseCase
}

func NewNetHTTPGetProductAdapter(useCase application.GetProductsUseCase) *NetHTTPGetProductAdapter {
	return &NetHTTPGetProductAdapter{
		useCase: useCase,
	}
}

func (a *NetHTTPGetProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrHttpMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Query().Get("id")
	input := application.GetProductInput{
		ID: productID,
	}

	product, err := a.useCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), HttpError[err])
		return
	}

	res := GetProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
