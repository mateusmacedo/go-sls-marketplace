package http

import (
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

var HttpError = map[error]int{
	domain.ErrInvalidProductID:          http.StatusBadRequest,
	domain.ErrInvalidProductName:        http.StatusBadRequest,
	domain.ErrInvalidProductDescription: http.StatusBadRequest,
	domain.ErrInvalidProductPrice:       http.StatusBadRequest,
	domain.ErrAlreadyExistsProduct:      http.StatusConflict,
	domain.ErrNotFoundProduct:           http.StatusNotFound,
	domain.ErrRepositoryProduct:         http.StatusInternalServerError,
}
