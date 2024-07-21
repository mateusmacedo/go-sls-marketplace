package application

import (
	"fmt"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

func CreateAddProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service, ok := dependencies["ProductAdder"].(domain.ProductAdder)
	if !ok || service == nil {
		return nil, fmt.Errorf("missing or invalid ProductAdder dependency")
	}
	return NewAddProductUseCase(service), nil
}

func CreateDeleteProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service, ok := dependencies["ProductDeleter"].(domain.ProductDeleter)
	if !ok || service == nil {
		return nil, fmt.Errorf("missing or invalid ProductDeleter dependency")
	}
	return NewDeleteProductUseCase(service), nil
}

func CreateGetAllProductsUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder, ok := dependencies["AllProductFinder"].(domain.AllProductFinder)
	if !ok || finder == nil {
		return nil, fmt.Errorf("missing or invalid AllProductFinder dependency")
	}
	return NewGetAllProductsUseCase(finder), nil
}

func CreateGetProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder, ok := dependencies["ProductFinder"].(domain.ProductFinder)
	if !ok || finder == nil {
		return nil, fmt.Errorf("missing or invalid ProductFinder dependency")
	}
	return NewGetProductUseCase(finder), nil
}

func CreateUpdateProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service, ok := dependencies["ProductUpdater"].(domain.ProductUpdater)
	if !ok || service == nil {
		return nil, fmt.Errorf("missing or invalid ProductUpdater dependency")
	}
	return NewUpdateProductUseCase(service), nil
}
