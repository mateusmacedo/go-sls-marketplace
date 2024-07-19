package domain

import (
	"time"
)

type ProductID string

type Product struct {
	ID          ProductID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(id ProductID, name, description string, price float64) (*Product, error) {
	if id == "" || name == "" || description == "" || price <= 0 {
		return nil, ErrInvalidProductID
	}

	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (p *Product) ChangeName(newName string) error {
	if newName == "" {
		return ErrInvalidProductName
	}
	p.Name = newName
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) ChangeDescription(newDescription string) error {
	if newDescription == "" {
		return ErrInvalidProductDescription
	}
	p.Description = newDescription
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) ChangePrice(newPrice float64) error {
	if newPrice <= 0 {
		return ErrInvalidProductPrice
	}
	p.Price = newPrice
	p.UpdatedAt = time.Now()
	return nil
}
