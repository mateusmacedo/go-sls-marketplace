package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
)

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type NetHTTPDeleteProductAdapter struct {
	useCase application.DeleteProductUseCase
}

func NewNetHTTPDeleteProductAdapter(useCase application.DeleteProductUseCase) *NetHTTPDeleteProductAdapter {
	return &NetHTTPDeleteProductAdapter{
		useCase: useCase,
	}
}

func (a *NetHTTPDeleteProductAdapter) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := map[string]string{"error": pkghttp.ErrHttpMethodNotAllowed.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		response := map[string]string{"error": domain.ErrInvalidProductID.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(infrahttp.HttpError[domain.ErrInvalidProductID])
		json.NewEncoder(w).Encode(response)
		return
	}
	err := a.useCase.Execute(application.DeleteProductInput{
		ID: productID,
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

	w.WriteHeader(http.StatusNoContent)
}
