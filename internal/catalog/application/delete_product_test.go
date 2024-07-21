package application

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/domain/mocks"
)

func TestDeleteProductUseCase_Execute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockProductDeleter(mockCtrl)
	useCase := NewDeleteProductUseCase(mockService)

	testCases := []struct {
		name          string
		input         DeleteProductInput
		mockBehavior  func(*mocks.MockProductDeleter, domain.ProductID)
		expectedError error
	}{
		{
			name:  "Successful deletion of product",
			input: DeleteProductInput{ID: "1"},
			mockBehavior: func(m *mocks.MockProductDeleter, id domain.ProductID) {
				m.EXPECT().DeleteProduct(id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:  "Product not found",
			input: DeleteProductInput{ID: "999"},
			mockBehavior: func(m *mocks.MockProductDeleter, id domain.ProductID) {
				m.EXPECT().DeleteProduct(id).Return(errors.New("product not found"))
			},
			expectedError: errors.New("product not found"),
		},
		{
			name:  "Database error",
			input: DeleteProductInput{ID: "2"},
			mockBehavior: func(m *mocks.MockProductDeleter, id domain.ProductID) {
				m.EXPECT().DeleteProduct(id).Return(errors.New("database error"))
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(mockService, domain.ProductID(tc.input.ID))

			err := useCase.Execute(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
