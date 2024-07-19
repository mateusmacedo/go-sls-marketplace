package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockAllProductFinder struct {
	mock.Mock
}

func (m *MockAllProductFinder) GetAllProducts() ([]*domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func TestGetAllProductsUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name             string
		mockBehavior     func(*MockAllProductFinder)
		expectedProducts []*GetAllProductsOutput
		expectedError    error
	}{
		{
			name: "Successful retrieval of products",
			mockBehavior: func(m *MockAllProductFinder) {
				createdAt := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
				updatedAt := time.Date(2023, 5, 2, 11, 0, 0, 0, time.UTC)
				m.On("GetAllProducts").Return([]*domain.Product{
					{
						ID:          domain.ProductID("1"),
						Name:        "Product 1",
						Description: "Description 1",
						Price:       10.0,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					},
					{
						ID:          "2",
						Name:        "Product 2",
						Description: "Description 2",
						Price:       20.0,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					},
				}, nil)
			},
			expectedProducts: []*GetAllProductsOutput{
				{
					ID:          "1",
					Name:        "Product 1",
					Description: "Description 1",
					Price:       10.0,
					CreatedAt:   "2023-05-01 10:00:00",
					UpdatedAt:   "2023-05-02 11:00:00",
				},
				{
					ID:          "2",
					Name:        "Product 2",
					Description: "Description 2",
					Price:       20.0,
					CreatedAt:   "2023-05-01 10:00:00",
					UpdatedAt:   "2023-05-02 11:00:00",
				},
			},
			expectedError: nil,
		},
		{
			name: "Error retrieving products",
			mockBehavior: func(m *MockAllProductFinder) {
				m.On("GetAllProducts").Return([]*domain.Product{}, domain.ErrRepositoryProduct)
			},
			expectedProducts: []*GetAllProductsOutput{},
			expectedError:    domain.ErrRepositoryProduct,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockFinder := new(MockAllProductFinder)
			useCase := NewGetAllProductsUseCase(mockFinder)

			tc.mockBehavior(mockFinder)

			products, err := useCase.Execute()

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

			mockFinder.AssertExpectations(t)
		})
	}
}
