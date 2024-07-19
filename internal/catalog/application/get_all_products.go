package application

import (
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type GetAllProductsOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetAllProductsUseCase interface {
	Execute() ([]*GetAllProductsOutput, error)
}

type getAllProductsUseCase struct {
	productFinder domain.AllProductFinder
}

func NewGetAllProductsUseCase(productFinder domain.AllProductFinder) GetAllProductsUseCase {
	return &getAllProductsUseCase{
		productFinder: productFinder,
	}
}

func (u *getAllProductsUseCase) Execute() ([]*GetAllProductsOutput, error) {
	productsOutput := []*GetAllProductsOutput{}
	products, err := u.productFinder.GetAllProducts()
	if err != nil {
		return productsOutput, err
	}
	for _, product := range products {
		productsOutput = append(productsOutput, &GetAllProductsOutput{
			ID:          string(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
		})
	}

	return productsOutput, nil
}
