package application

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

// MockProductFinder é um mock da interface ProductFinder
type MockProductFinder struct {
	mock.Mock
}

func (m *MockProductFinder) GetAllProducts() ([]*domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductFinder) GetProduct(id domain.ProductID) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestGetAllProductsUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name             string
		mockBehavior     func(*MockProductFinder)
		expectedProducts []*domain.Product
		expectedError    error
	}{
		{
			name: "Successful retrieval of products",
			mockBehavior: func(m *MockProductFinder) {
				m.On("GetAllProducts").Return([]*domain.Product{
					{
						ID:          "1",
						Name:        "Product 1",
						Description: "Description 1",
						Price:       10.0,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
					{
						ID:          "2",
						Name:        "Product 2",
						Description: "Description 2",
						Price:       20.0,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				}, nil)
			},
			expectedProducts: []*domain.Product{
				{
					ID:          "1",
					Name:        "Product 1",
					Description: "Description 1",
					Price:       10.0,
				},
				{
					ID:          "2",
					Name:        "Product 2",
					Description: "Description 2",
					Price:       20.0,
				},
			},
			expectedError: nil,
		},
		{
			name: "Error retrieving products",
			mockBehavior: func(m *MockProductFinder) {
				m.On("GetAllProducts").Return([]*domain.Product(nil), errors.New("database error"))
			},
			expectedProducts: []*domain.Product{},
			expectedError:    errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock and use case for each test
			mockFinder := new(MockProductFinder)
			useCase := NewGetAllProductsUseCase(mockFinder)

			// Set up mock behavior
			tc.mockBehavior(mockFinder)

			// Execute the use case
			products, err := useCase.Execute()

			// Check the results
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				assert.Equal(t, len(tc.expectedProducts), len(products))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expectedProducts), len(products))

				for i, expectedProduct := range tc.expectedProducts {
					assert.Equal(t, expectedProduct.ID, products[i].ID)
					assert.Equal(t, expectedProduct.Name, products[i].Name)
					assert.Equal(t, expectedProduct.Description, products[i].Description)
					assert.Equal(t, expectedProduct.Price, products[i].Price)
					assert.NotZero(t, products[i].CreatedAt)
					assert.NotZero(t, products[i].UpdatedAt)
				}
			}

			// Verify that all expected mock calls were made
			mockFinder.AssertExpectations(t)
		})
	}
}
