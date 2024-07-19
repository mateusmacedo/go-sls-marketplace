package adapter

import (
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type DynamoDbProductEntity struct {
	ID          string    `json:"id" dynamodbav:"id"`
	Name        string    `json:"name" dynamodbav:"name"`
	Description string    `json:"description" dynamodbav:"description"`
	Price       float64   `json:"price" dynamodbav:"price"`
	CreatedAt   time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" dynamodbav:"updated_at"`
}

func (pe *DynamoDbProductEntity) ToDomain() (*domain.Product, error) {
	return &domain.Product{
		ID:          domain.ProductID(pe.ID),
		Name:        pe.Name,
		Description: pe.Description,
		Price:       pe.Price,
		CreatedAt:   pe.CreatedAt,
		UpdatedAt:   pe.UpdatedAt,
	}, nil
}

func NewProductEntityFromDomain(product *domain.Product) (*DynamoDbProductEntity, error) {
	if product == nil {
		return nil, domain.ErrInvalidProductID
	}
	return &DynamoDbProductEntity{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
