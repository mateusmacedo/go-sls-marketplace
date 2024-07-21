package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductAdder(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: nil,
		},
		{
			name: "Missing ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing ProductSaveRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": nil,
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductSaveRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
				"ProductSaveRepository": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adder, err := CreateProductAdder(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, adder)
			}
		})
	}
}

// Similar tests for other creation functions
func TestCreateProductDeleter(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockDeleteRepo := new(MockProductDeleteRepository)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFindRepository":   mockFindRepo,
				"ProductDeleteRepository": mockDeleteRepo,
			},
			expectedErr: nil,
		},
		{
			name: "Missing ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductDeleteRepository": mockDeleteRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing ProductDeleteRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository":   nil,
				"ProductDeleteRepository": mockDeleteRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductDeleteRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository":   mockFindRepo,
				"ProductDeleteRepository": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleter, err := CreateProductDeleter(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, deleter)
			}
		})
	}
}

func TestCreateProductFinder(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductFindRepository",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finder, err := CreateProductFinder(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, finder)
			}
		})
	}
}

func TestCreateAllProductFinder(t *testing.T) {
	mockFindAllRepo := new(MockProductFindAllRepository)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFindAllRepository": mockFindAllRepo,
			},
			expectedErr: nil,
		},
		{
			name:         "Missing ProductFindAllRepository",
			dependencies: map[string]interface{}{},
			expectedErr:  assert.AnError,
		},
		{
			name: "Nil ProductFindAllRepository",
			dependencies: map[string]interface{}{
				"ProductFindAllRepository": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allFinder, err := CreateAllProductFinder(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, allFinder)
			}
		})
	}
}

func TestCreateProductUpdater(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: nil,
		},
		{
			name: "Missing ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing ProductSaveRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductFindRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": nil,
				"ProductSaveRepository": mockSaveRepo,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil ProductSaveRepository",
			dependencies: map[string]interface{}{
				"ProductFindRepository": mockFindRepo,
				"ProductSaveRepository": nil,
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updater, err := CreateProductUpdater(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updater)
			}
		})
	}
}
