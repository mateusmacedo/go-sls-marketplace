package application

import (
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type AddProductInput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type AddProductOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type AddProductUseCase interface {
	Execute(input AddProductInput) (*AddProductOutput, error)
}

type addProductUseCase struct {
	productAdder domain.ProductAdder
}

func NewProductAddUseCase(ProductAdder domain.ProductAdder) AddProductUseCase {
	return &addProductUseCase{
		productAdder: ProductAdder,
	}
}

func (u *addProductUseCase) Execute(input AddProductInput) (*AddProductOutput, error) {
	id := domain.ProductID(input.ID)
	name := input.Name
	description := input.Description
	price := input.Price

	product, err := u.productAdder.AddProduct(id, name, description, price)
	if err != nil {
		return nil, err
	}

	return &AddProductOutput{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}, nil
}
