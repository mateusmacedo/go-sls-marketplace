package application

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockProductFinder struct {
	mock.Mock
}

func (m *MockProductFinder) GetProduct(id domain.ProductID) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestGetProductsUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name           string
		input          GetProductInput
		mockBehavior   func(*MockProductFinder, domain.ProductID)
		expectedOutput *GetProductOutput
		expectedError  error
	}{
		{
			name:  "Successful retrieval of product",
			input: GetProductInput{ID: "1"},
			mockBehavior: func(m *MockProductFinder, id domain.ProductID) {
				createdAt := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
				updatedAt := time.Date(2023, 5, 2, 11, 0, 0, 0, time.UTC)
				m.On("GetProduct", id).Return(&domain.Product{
					ID:          id,
					Name:        "Test Product",
					Description: "Test Description",
					Price:       10.99,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				}, nil)
			},
			expectedOutput: &GetProductOutput{
				ID:          "1",
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.99,
				CreatedAt:   "2023-05-01 10:00:00",
				UpdateAt:    "2023-05-02 11:00:00",
			},
			expectedError: nil,
		},
		{
			name:  "Product not found",
			input: GetProductInput{ID: "999"},
			mockBehavior: func(m *MockProductFinder, id domain.ProductID) {
				m.On("GetProduct", id).Return(nil, errors.New("product not found"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("product not found"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockFinder := new(MockProductFinder)
			useCase := NewGetProductsUseCase(mockFinder)

			tc.mockBehavior(mockFinder, domain.ProductID(tc.input.ID))

			output, err := useCase.Execute(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, output)
			}

			mockFinder.AssertExpectations(t)
		})
	}
}