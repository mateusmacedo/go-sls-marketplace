package application

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

type DeleteProductInput struct {
	ID string `json:"id"`
}

type DeleteProductUseCase struct {
	productDeleter domain.ProductDeleter
}

func NewDeleteProductUseCase(productDeleter domain.ProductDeleter) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productDeleter: productDeleter,
	}
}

func (u *DeleteProductUseCase) Execute(input DeleteProductInput) error {
	id := domain.ProductID(input.ID)

	err := u.productDeleter.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
