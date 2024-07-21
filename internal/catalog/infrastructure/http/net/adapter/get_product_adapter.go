package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
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
	useCase application.GetProductUseCase
}

func NewNetHTTPGetProductAdapter(useCase application.GetProductUseCase) *NetHTTPGetProductAdapter {
	return &NetHTTPGetProductAdapter{
		useCase: useCase,
	}
}

func (a *NetHTTPGetProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "`+pkghttp.ErrHttpMethodNotAllowed.Error()+`"}`, http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Path[len("/products/"):]
	input := application.GetProductInput{
		ID: productID,
	}

	product, err := a.useCase.Execute(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "`+err.Error()+`"}`, infrahttp.HttpError[err])
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
