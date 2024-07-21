package application

import (
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

func CreateAddProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service := dependencies["ProductAdder"].(domain.ProductAdder)
	return NewAddProductUseCase(service), nil
}

func CreateDeleteProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service := dependencies["ProductDeleter"].(domain.ProductDeleter)
	return NewDeleteProductUseCase(service), nil
}

func CreateGetAllProductsUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder := dependencies["AllProductFinder"].(domain.AllProductFinder)
	return NewGetAllProductsUseCase(finder), nil
}

func CreateGetProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder := dependencies["ProductFinder"].(domain.ProductFinder)
	return NewGetProductUseCase(finder), nil
}

func CreateUpdateProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service := dependencies["ProductUpdater"].(domain.ProductUpdater)
	return NewUpdateProductUseCase(service), nil
}
