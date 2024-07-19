package application

import (
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type GetProductInput struct {
	ID string `json:"id"`
}

type GetProductOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetProductUseCase interface {
	Execute(input GetProductInput) (*GetProductOutput, error)
}

type getProductsUseCase struct {
	productFinder domain.ProductFinder
}

func NewGetProductUseCase(productFinder domain.ProductFinder) GetProductUseCase {
	return &getProductsUseCase{
		productFinder: productFinder,
	}
}

func (u *getProductsUseCase) Execute(input GetProductInput) (*GetProductOutput, error) {
	id := domain.ProductID(input.ID)

	product, err := u.productFinder.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return &GetProductOutput{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}, nil
}
