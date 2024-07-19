package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

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

type GetProductsUseCase interface {
	Execute(input GetProductInput) (*GetProductOutput, error)
}

type getProductsUseCase struct {
	productFinder domain.ProductFinder
}

func NewGetProductsUseCase(productFinder domain.ProductFinder) GetProductsUseCase {
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
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
