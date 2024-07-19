package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type GetAllProductsUseCase interface {
	Execute() ([]*domain.Product, error)
}

type getAllProductsUseCase struct {
	productFinder domain.AllProductFinder
}

func NewGetAllProductsUseCase(productFinder domain.AllProductFinder) GetAllProductsUseCase {
	return &getAllProductsUseCase{
		productFinder: productFinder,
	}
}

func (u *getAllProductsUseCase) Execute() ([]*domain.Product, error) {
	products, err := u.productFinder.GetAllProducts()
	if err != nil {
		return products, err
	}
	return products, nil
}
