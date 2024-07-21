package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddProductUseCase(t *testing.T) {
	mockProductAdder := new(MockProductAdder)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductAdder": mockProductAdder,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductAdder",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductAdder",
			dependencies: map[string]interface{}{
				"ProductAdder": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase, err := CreateAddProductUseCase(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, useCase)
			}
		})
	}
}

// Similar tests for other creation functions
func TestCreateDeleteProductUseCase(t *testing.T) {
	mockProductDeleter := new(MockProductDeleter)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductDeleter": mockProductDeleter,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductDeleter",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductDeleter",
			dependencies: map[string]interface{}{
				"ProductDeleter": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase, err := CreateDeleteProductUseCase(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, useCase)
			}
		})
	}
}

func TestCreateGetAllProductsUseCase(t *testing.T) {
	mockAllProductFinder := new(MockAllProductFinder)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"AllProductFinder": mockAllProductFinder,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing AllProductFinder",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil AllProductFinder",
			dependencies: map[string]interface{}{
				"AllProductFinder": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase, err := CreateGetAllProductsUseCase(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, useCase)
			}
		})
	}
}

func TestCreateGetProductUseCase(t *testing.T) {
	mockProductFinder := new(MockProductFinder)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFinder": mockProductFinder,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductFinder",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductFinder",
			dependencies: map[string]interface{}{
				"ProductFinder": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase, err := CreateGetProductUseCase(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, useCase)
			}
		})
	}
}

func TestCreateUpdateProductUseCase(t *testing.T) {
	mockProductUpdater := new(MockProductUpdater)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductUpdater": mockProductUpdater,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductUpdater",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductUpdater",
			dependencies: map[string]interface{}{
				"ProductUpdater": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase, err := CreateUpdateProductUseCase(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, useCase)
			}
		})
	}
}
