package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type UpdateProductInput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type UpdateProductOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdateAt    string  `json:"updated_at"`
}

type UpdateProductUseCase struct {
	productUpdater domain.ProductUpdater
}

func NewUpdateProductUseCase(productUpdater domain.ProductUpdater) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productUpdater: productUpdater,
	}
}

func (u *UpdateProductUseCase) Execute(input UpdateProductInput) (*UpdateProductOutput, error) {
	id := domain.ProductID(input.ID)

	product, err := u.productUpdater.UpdateProduct(id, input.Name, input.Description, input.Price)
	if err != nil {
		return nil, err
	}

	return &UpdateProductOutput{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
