package adapter

import (
	"time"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type ProductEntity struct {
	ID          string    `gorm:"primaryKey;type:text;default:(lower(hex(randomblob(16))))" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}

func (pe *ProductEntity) ToDomain() (*domain.Product, error) {
	return &domain.Product{
		ID:          domain.ProductID(pe.ID),
		Name:        pe.Name,
		Description: pe.Description,
		Price:       pe.Price,
		CreatedAt:   pe.CreatedAt,
		UpdatedAt:   pe.UpdatedAt,
	}, nil
}

func NewProductEntityFromDomain(product *domain.Product) (*ProductEntity, error) {
	if product == nil {
		return nil, domain.ErrInvalidProductID
	}
	return &ProductEntity{
		ID:          string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
