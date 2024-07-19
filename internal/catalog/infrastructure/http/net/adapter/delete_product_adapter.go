package adapter

import (
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
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
		http.Error(w, ErrHttpMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		http.Error(w, domain.ErrInvalidProductID.Error(), http.StatusBadRequest)
		return
	}
	err := a.useCase.Execute(application.DeleteProductInput{
		ID: productID,
	})
	if err != nil {
		http.Error(w, err.Error(), HttpError[err])
		return
	}

	w.WriteHeader(http.StatusNoContent)
}