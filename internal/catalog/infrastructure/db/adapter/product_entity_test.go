package adapter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

func TestProductEntity_ToDomain(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		pe      ProductEntity
		want    *domain.Product
		wantErr error
	}{
		{
			name: "Valid ProductEntity",
			pe: ProductEntity{
				ID:          "test-id",
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       9.99,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			want: &domain.Product{
				ID:          domain.ProductID("test-id"),
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       9.99,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pe.ToDomain()
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewProductEntityFromDomain(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		product *domain.Product
		want    *ProductEntity
		wantErr error
	}{
		{
			name: "Valid Product",
			product: &domain.Product{
				ID:          domain.ProductID("test-id"),
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       9.99,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			want: &ProductEntity{
				ID:          "test-id",
				Name:        "Test Product",
				Description: "This is a test product",
				Price:       9.99,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: nil,
		},
		{
			name:    "Nil Product",
			product: nil,
			want:    nil,
			wantErr: domain.ErrInvalidProductID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProductEntityFromDomain(tt.product)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
