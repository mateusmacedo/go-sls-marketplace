package adapter

import (
	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

// Função para criar o product save repository.
func CreateProductSaveRepository(dependencies map[string]interface{}) (domain.ProductSaveRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductSaveRepository(db), nil
}

// Função para criar o product find repository.
func CreateProductFindRepository(dependencies map[string]interface{}) (domain.ProductFindRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductFindRepository(db), nil
}

// Função para criar o product find all repository.
func CreateProductFindAllRepository(dependencies map[string]interface{}) (domain.ProductFindAllRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductFindAllRepository(db), nil
}

// Função para criar o product delete repository.
func CreateProductDeleteRepository(dependencies map[string]interface{}) (domain.ProductDeleteRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductDeleteRepository(db), nil
}
