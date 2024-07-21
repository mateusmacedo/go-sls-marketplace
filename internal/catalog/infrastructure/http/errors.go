package http

import (
	"net/http"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	pkginfrahttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
)

var HttpError = map[error]int{
	domain.ErrInvalidProductID:          http.StatusBadRequest,
	domain.ErrInvalidProductName:        http.StatusBadRequest,
	domain.ErrInvalidProductDescription: http.StatusBadRequest,
	domain.ErrInvalidProductPrice:       http.StatusBadRequest,
	domain.ErrAlreadyExistsProduct:      http.StatusConflict,
	domain.ErrNotFoundProduct:           http.StatusNotFound,
	domain.ErrRepositoryProduct:         http.StatusInternalServerError,
	pkginfrahttp.ErrHttpInvalidJSON:     http.StatusBadRequest,
	pkginfrahttp.ErrServiceError:        http.StatusInternalServerError,
}
