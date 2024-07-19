package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type DeleteProductInput struct {
	ID string `json:"id"`
}

type DeleteProductUseCase interface {
	Execute(input DeleteProductInput) error
}

type deleteProductUseCase struct {
	productDeleter domain.ProductDeleter
}

func NewDeleteProductUseCase(productDeleter domain.ProductDeleter) *deleteProductUseCase {
	return &deleteProductUseCase{
		productDeleter: productDeleter,
	}
}

func (u *deleteProductUseCase) Execute(input DeleteProductInput) error {
	id := domain.ProductID(input.ID)

	err := u.productDeleter.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
