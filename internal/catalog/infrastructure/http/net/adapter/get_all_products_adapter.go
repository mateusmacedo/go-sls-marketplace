package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
)

type GetAllProductsResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type NetHTTPGetAllProductsAdapter struct {
	useCase application.GetAllProductsUseCase
}

func NewNetHTTPGetAllProductsAdapter(useCase application.GetAllProductsUseCase) *NetHTTPGetAllProductsAdapter {
	return &NetHTTPGetAllProductsAdapter{
		useCase: useCase,
	}
}

func (a *NetHTTPGetAllProductsAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	products, err := a.useCase.Execute()
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, infrahttp.HttpError[err])
		return
	}

	response := make([]GetAllProductsResponse, len(products))
	for i, product := range products {
		response[i] = GetAllProductsResponse{
			ID:          string(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
