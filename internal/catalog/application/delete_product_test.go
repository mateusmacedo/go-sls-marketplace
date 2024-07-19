package application

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockProductDeleter struct {
	mock.Mock
}

func (m *MockProductDeleter) DeleteProduct(id domain.ProductID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestDeleteProductUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name          string
		input         DeleteProductInput
		mockBehavior  func(*MockProductDeleter, domain.ProductID)
		expectedError error
	}{
		{
			name:  "Successful deletion of product",
			input: DeleteProductInput{ID: "1"},
			mockBehavior: func(m *MockProductDeleter, id domain.ProductID) {
				m.On("DeleteProduct", id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:  "Product not found",
			input: DeleteProductInput{ID: "999"},
			mockBehavior: func(m *MockProductDeleter, id domain.ProductID) {
				m.On("DeleteProduct", id).Return(errors.New("product not found"))
			},
			expectedError: errors.New("product not found"),
		},
		{
			name:  "Database error",
			input: DeleteProductInput{ID: "2"},
			mockBehavior: func(m *MockProductDeleter, id domain.ProductID) {
				m.On("DeleteProduct", id).Return(errors.New("database error"))
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			mockDeleter := new(MockProductDeleter)
			useCase := NewDeleteProductUseCase(mockDeleter)

			tc.mockBehavior(mockDeleter, domain.ProductID(tc.input.ID))

			err := useCase.Execute(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockDeleter.AssertExpectations(t)
		})
	}
}
