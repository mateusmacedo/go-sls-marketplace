package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type ProductAddInput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductAddOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdateAt    string  `json:"updated_at"`
}

type AddProductUseCase struct {
	productAdder domain.ProductAdder
}

func NewProductAddUseCase(ProductAdder domain.ProductAdder) *AddProductUseCase {
	return &AddProductUseCase{
		productAdder: ProductAdder,
	}
}

func (u *AddProductUseCase) Execute(input ProductAddInput) (*ProductAddOutput, error) {
	id := domain.ProductID(input.ID)
	name := input.Name
	description := input.Description
	price := input.Price

	product, err := u.productAdder.AddProduct(id, name, description, price)
	if err != nil {
		return nil, err
	}

	return &ProductAddOutput{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
