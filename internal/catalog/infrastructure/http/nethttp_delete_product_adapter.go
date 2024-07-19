package http

import (
	"encoding/json"
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
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

	var req DeleteProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := a.useCase.Execute(application.DeleteProductInput{
		ID: req.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), HttpError[err])
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
