package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
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

type NetHTTPAddProductAdapter struct {
	service     application.AddProductUseCase
	methodGuard pkghttp.HttpMethodGuard
}

type HTTPAddProductAdapterOption func(*NetHTTPAddProductAdapter) error

func WithService(service application.AddProductUseCase) HTTPAddProductAdapterOption {
	return func(a *NetHTTPAddProductAdapter) error {
		a.service = service
		return nil
	}
}

func WithMethodGuard(methodGuard pkghttp.HttpMethodGuard) HTTPAddProductAdapterOption {
	return func(a *NetHTTPAddProductAdapter) error {
		a.methodGuard = methodGuard
		return nil
	}
}

func NewNetHTTPAddProductAdapter(opts ...HTTPAddProductAdapterOption) *NetHTTPAddProductAdapter {
	adapter := &NetHTTPAddProductAdapter{}

	for _, opt := range opts {
		opt(adapter)
	}

	return adapter
}

func (a *NetHTTPAddProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if !a.methodGuard.IsMethodAllowed(r.Method) {
		response := map[string]string{"error": pkghttp.ErrHttpMethodNotAllowed.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	var req AddProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := map[string]string{"error": pkghttp.ErrHttpInvalidJSON.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(infrahttp.HttpError[pkghttp.ErrHttpInvalidJSON])
		json.NewEncoder(w).Encode(response)
		return
	}

	product, err := a.service.Execute(application.AddProductInput{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		statusCode, ok := infrahttp.HttpError[err]
		if !ok {
			err = pkghttp.ErrServiceError
			statusCode = infrahttp.HttpError[err]
		}
		response := map[string]string{"error": err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
