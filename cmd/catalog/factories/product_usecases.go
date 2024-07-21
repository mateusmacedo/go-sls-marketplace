package factories

import (
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

// Função para criar o caso de uso de adição de produto.
func CreateAddProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service := dependencies["ProductAdder"].(domain.ProductAdder)
	return application.NewAddProductUseCase(service), nil
}

// Função para criar o caso de uso de deleção de produto.
func CreateDeleteProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	service := dependencies["ProductDeleter"].(domain.ProductDeleter)
	return application.NewDeleteProductUseCase(service), nil
}

// Função para criar o caso de uso de listagem de produtos.
func CreateGetAllProductsUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder := dependencies["AllProductFinder"].(domain.AllProductFinder)
	return application.NewGetAllProductsUseCase(finder), nil
}

// Função para criar o caso de uso de listagem de produtos por ID.
func CreateGetProductUseCase(dependencies map[string]interface{}) (interface{}, error) {
	finder := dependencies["ProductFinder"].(domain.ProductFinder)
	return application.NewGetProductUseCase(finder), nil
}
