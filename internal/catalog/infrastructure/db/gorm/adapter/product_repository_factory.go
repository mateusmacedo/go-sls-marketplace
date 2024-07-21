package adapter

import (
	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

func CreateProductSaveRepository(dependencies map[string]interface{}) (domain.ProductSaveRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductSaveRepository(db), nil
}

func CreateProductFindRepository(dependencies map[string]interface{}) (domain.ProductFindRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductFindRepository(db), nil
}

func CreateProductFindAllRepository(dependencies map[string]interface{}) (domain.ProductFindAllRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductFindAllRepository(db), nil
}

func CreateProductDeleteRepository(dependencies map[string]interface{}) (domain.ProductDeleteRepository, error) {
	db := dependencies["db"].(*gorm.DB)

	return NewGormProductDeleteRepository(db), nil
}
