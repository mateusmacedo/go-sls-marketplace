package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
}

type UpdateProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type NetHTTPUpdateProductAdapter struct {
	useCase application.UpdateProductUseCase
}

func NewNetHTTPUpdateProductAdapter(useCase application.UpdateProductUseCase) *NetHTTPUpdateProductAdapter {
	return &NetHTTPUpdateProductAdapter{useCase: useCase}
}

func (a *NetHTTPUpdateProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, ErrHttpMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		http.Error(w, domain.ErrInvalidProductID.Error(), http.StatusBadRequest)
		return
	}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input := application.UpdateProductInput{
		ID:          productID,
		Name:        *req.Name,
		Description: *req.Description,
		Price:       *req.Price,
	}

	product, err := a.useCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), HttpError[err])
		return
	}

	res := UpdateProductResponse{
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
