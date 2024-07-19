package infrastructure

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
)

type ProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := a.useCase.Execute()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := make([]ProductResponse, len(products))
	for i, product := range products {
		response[i] = ProductResponse{
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
