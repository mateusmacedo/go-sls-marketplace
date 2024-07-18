package domain

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if product == nil {
		t.Error("Expected product to be created, but got nil")
	}

	if product == nil || product.ID != id {
		t.Errorf("Expected product ID to be %s, but got %s", id, product.ID)
	}

	if product.Name != name {
		t.Errorf("Expected product name to be %s, but got %s", name, product.Name)
	}

	if product.Description != description {
		t.Errorf("Expected product description to be %s, but got %s", description, product.Description)
	}

	if product.Price != price {
		t.Errorf("Expected product price to be %f, but got %f", price, product.Price)
	}
}

func TestNewProductInvalidAttributes(t *testing.T) {
	id := ProductID("")
	name := ""
	description := ""
	price := 0.0

	product, err := NewProduct(id, name, description, price)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if product != nil {
		t.Errorf("Expected product to be nil, but got: %v", product)
	}
}

func TestChangeName(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newName := "Updated Name"
	err = product.ChangeName(newName)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if product.Name != newName {
		t.Errorf("Expected product name to be %s, but got %s", newName, product.Name)
	}
}

func TestChangeNameInvalid(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newName := ""
	err = product.ChangeName(newName)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if product.Name != name {
		t.Errorf("Expected product name to be %s, but got %s", name, product.Name)
	}
}

func TestChangeDescription(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newDescription := "Updated description"
	err = product.ChangeDescription(newDescription)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if product.Description != newDescription {
		t.Errorf("Expected product description to be %s, but got %s", newDescription, product.Description)
	}
}

func TestChangeDescriptionInvalid(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newDescription := ""
	err = product.ChangeDescription(newDescription)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if product.Description != description {
		t.Errorf("Expected product description to be %s, but got %s", description, product.Description)
	}
}

func TestChangePrice(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newPrice := 19.99
	err = product.ChangePrice(newPrice)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if product.Price != newPrice {
		t.Errorf("Expected product price to be %f, but got %f", newPrice, product.Price)
	}
}

func TestChangePriceInvalid(t *testing.T) {
	id := ProductID("123")
	name := "Test Product"
	description := "This is a test product"
	price := 9.99

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	newPrice := 0.0
	err = product.ChangePrice(newPrice)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if product.Price != price {
		t.Errorf("Expected product price to be %f, but got %f", price, product.Price)
	}
}
