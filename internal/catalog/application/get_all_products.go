package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type GetAllProductsUseCase struct {
	productFinder domain.ProductFinder
}

func NewGetAllProductsUseCase(productFinder domain.ProductFinder) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		productFinder: productFinder,
	}
}

func (u *GetAllProductsUseCase) Execute() ([]*domain.Product, error) {
	products, err := u.productFinder.GetAllProducts()
	if err != nil {
		return products, err
	}
	return products, nil
}
