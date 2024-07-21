package application

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/domain/mocks"
)

func TestAddProductUseCase_Execute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductAdder := mocks.NewMockProductAdder(mockCtrl)

	tests := []struct {
		name           string
		input          AddProductInput
		mockOutput     *domain.Product
		mockError      error
		expectedOutput *AddProductOutput
		expectedError  error
	}{
		{
			name: "Success",
			input: AddProductInput{
				ID:          "1",
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
			},
			mockOutput: &domain.Product{
				ID:          domain.ProductID("1"),
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
				CreatedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			expectedOutput: &AddProductOutput{
				ID:          "1",
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
				CreatedAt:   "2021-01-01T00:00:00Z",
				UpdatedAt:   "2021-01-02T00:00:00Z",
			},
		},
		{
			name: "Service Error",
			input: AddProductInput{
				ID:          "1",
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
			},
			mockError:     errors.New("some service error"),
			expectedError: errors.New("some service error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockOutput != nil || tt.mockError != nil {
				mockProductAdder.EXPECT().
					AddProduct(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(tt.mockOutput, tt.mockError)
			}

			useCase := NewAddProductUseCase(mockProductAdder)
			output, err := useCase.Execute(tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, output)
			}
		})
	}
}
