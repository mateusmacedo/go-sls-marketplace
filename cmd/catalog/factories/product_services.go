package factories

import (
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

// Função para criar o serviço de adição de produto.
func CreateProductAdder(dependencies map[string]interface{}) (domain.ProductAdder, error) {
	findRepository := dependencies["findRepository"].(domain.ProductFindRepository)
	saveRepository := dependencies["saveRepository"].(domain.ProductSaveRepository)

	return domain.NewProductAdder(findRepository, saveRepository), nil
}

// Função para criar o serviço de deleção de produto.
func CreateProductDeleter(dependencies map[string]interface{}) (domain.ProductDeleter, error) {
	findRepository := dependencies["findRepository"].(domain.ProductFindRepository)
	deleteRepository := dependencies["deleteRepository"].(domain.ProductDeleteRepository)

	return domain.NewProductDeleter(findRepository, deleteRepository), nil
}

// Função para criar o serviço de listagem de produtos.
func CreateProductFinder(dependencies map[string]interface{}) (domain.ProductFinder, error) {
	findRepository := dependencies["findRepository"].(domain.ProductFindRepository)

	return domain.NewProductFinder(findRepository), nil
}

// Função para criar o serviço de listagem de todos os produtos.
func CreateAllProductFinder(dependencies map[string]interface{}) (domain.AllProductFinder, error) {
	findAllRepository := dependencies["findAllRepository"].(domain.ProductFindAllRepository)

	return domain.NewAllProductFinder(findAllRepository), nil
}

// Funcão para criar o serviço de atualização de produto.
func CreateProductUpdater(dependencies map[string]interface{}) (domain.ProductUpdater, error) {
	findRepository := dependencies["findRepository"].(domain.ProductFindRepository)
	saveRepository := dependencies["saveRepository"].(domain.ProductSaveRepository)

	return domain.NewProductUpdater(findRepository, saveRepository), nil
}
