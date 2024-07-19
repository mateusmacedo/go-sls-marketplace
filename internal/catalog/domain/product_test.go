package domain

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name        string
		id          ProductID
		productName string
		description string
		price       float64
		wantErr     bool
	}{
		{
			name:        "Valid Product",
			id:          ProductID("123"),
			productName: "Test Product",
			description: "This is a test product",
			price:       9.99,
			wantErr:     false,
		},
		{
			name:        "Invalid Product - Empty ID",
			id:          ProductID(""),
			productName: "Test Product",
			description: "This is a test product",
			price:       9.99,
			wantErr:     true,
		},
		{
			name:        "Invalid Product - Empty Name",
			id:          ProductID("123"),
			productName: "",
			description: "This is a test product",
			price:       9.99,
			wantErr:     true,
		},
		{
			name:        "Invalid Product - Empty Description",
			id:          ProductID("123"),
			productName: "Test Product",
			description: "",
			price:       9.99,
			wantErr:     true,
		},
		{
			name:        "Invalid Product - Zero Price",
			id:          ProductID("123"),
			productName: "Test Product",
			description: "This is a test product",
			price:       0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(tt.id, tt.productName, tt.description, tt.price)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if product == nil {
					t.Error("Expected product to be created, but got nil")
				} else {
					if product.ID != tt.id {
						t.Errorf("Expected product ID to be %s, but got %s", tt.id, product.ID)
					}
					if product.Name != tt.productName {
						t.Errorf("Expected product name to be %s, but got %s", tt.productName, product.Name)
					}
					if product.Description != tt.description {
						t.Errorf("Expected product description to be %s, but got %s", tt.description, product.Description)
					}
					if product.Price != tt.price {
						t.Errorf("Expected product price to be %f, but got %f", tt.price, product.Price)
					}
				}
			}
		})
	}
}

func TestProductMethods(t *testing.T) {
	validProduct, _ := NewProduct(ProductID("123"), "Test Product", "This is a test product", 9.99)

	tests := []struct {
		name        string
		method      func() error
		wantErr     bool
		checkResult func() bool
	}{
		{
			name:    "Change Name - Valid",
			method:  func() error { return validProduct.ChangeName("Updated Name") },
			wantErr: false,
			checkResult: func() bool {
				return validProduct.Name == "Updated Name"
			},
		},
		{
			name:    "Change Name - Invalid",
			method:  func() error { return validProduct.ChangeName("") },
			wantErr: true,
			checkResult: func() bool {
				return validProduct.Name == "Updated Name"
			},
		},
		{
			name:    "Change Description - Valid",
			method:  func() error { return validProduct.ChangeDescription("Updated description") },
			wantErr: false,
			checkResult: func() bool {
				return validProduct.Description == "Updated description"
			},
		},
		{
			name:    "Change Description - Invalid",
			method:  func() error { return validProduct.ChangeDescription("") },
			wantErr: true,
			checkResult: func() bool {
				return validProduct.Description == "Updated description"
			},
		},
		{
			name:    "Change Price - Valid",
			method:  func() error { return validProduct.ChangePrice(19.99) },
			wantErr: false,
			checkResult: func() bool {
				return validProduct.Price == 19.99
			},
		},
		{
			name:    "Change Price - Invalid",
			method:  func() error { return validProduct.ChangePrice(0) },
			wantErr: true,
			checkResult: func() bool {
				return validProduct.Price == 19.99
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.method()

			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}

			if !tt.checkResult() {
				t.Errorf("%s failed to update or incorrectly updated the product", tt.name)
			}
		})
	}
}
