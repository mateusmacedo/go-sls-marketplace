package application

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockProductUpdater struct {
	mock.Mock
}

func (m *MockProductUpdater) UpdateProduct(id domain.ProductID, name, description string, price float64) (*domain.Product, error) {
	args := m.Called(id, name, description, price)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestUpdateProductUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name           string
		input          UpdateProductInput
		mockBehavior   func(*MockProductUpdater, domain.ProductID, string, string, float64)
		expectedOutput *UpdateProductOutput
		expectedError  error
	}{
		{
			name: "Successful update of product",
			input: UpdateProductInput{
				ID:          "1",
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       15.99,
			},
			mockBehavior: func(m *MockProductUpdater, id domain.ProductID, name, description string, price float64) {
				createdAt := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
				updatedAt := time.Now()
				m.On("UpdateProduct", id, name, description, price).Return(&domain.Product{
					ID:          id,
					Name:        name,
					Description: description,
					Price:       price,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
				}, nil)
			},
			expectedOutput: &UpdateProductOutput{
				ID:          "1",
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       15.99,
				CreatedAt:   "2023-05-01T10:00:00Z",
			},
			expectedError: nil,
		},
		{
			name: "Product not found",
			input: UpdateProductInput{
				ID:          "999",
				Name:        "Non-existent Product",
				Description: "This product doesn't exist",
				Price:       9.99,
			},
			mockBehavior: func(m *MockProductUpdater, id domain.ProductID, name, description string, price float64) {
				m.On("UpdateProduct", id, name, description, price).Return(nil, errors.New("product not found"))
			},
			expectedOutput: nil,
			expectedError:  errors.New("product not found"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockUpdater := new(MockProductUpdater)
			useCase := NewUpdateProductUseCase(mockUpdater)

			tc.mockBehavior(mockUpdater, domain.ProductID(tc.input.ID), tc.input.Name, tc.input.Description, tc.input.Price)

			output, err := useCase.Execute(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput.ID, output.ID)
				assert.Equal(t, tc.expectedOutput.Name, output.Name)
				assert.Equal(t, tc.expectedOutput.Description, output.Description)
				assert.Equal(t, tc.expectedOutput.Price, output.Price)
				assert.Equal(t, tc.expectedOutput.CreatedAt, output.CreatedAt)

				_, parseErr := time.Parse("2006-01-02T15:04:05Z", output.CreatedAt)
				assert.NoError(t, parseErr)
			}

			mockUpdater.AssertExpectations(t)
		})
	}
}
