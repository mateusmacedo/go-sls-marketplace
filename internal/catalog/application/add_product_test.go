package application

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

// MockProductService é um mock para domain.ProductService
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) AddProduct(id domain.ProductID, name, description string, price float64) (*domain.Product, error) {
	args := m.Called(id, name, description, price)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestProductAddUseCase_Execute(t *testing.T) {
	mockService := new(MockProductService)
	useCase := NewProductAddUseCase(mockService)

	testCases := []struct {
		name           string
		input          AddProductInput
		mockBehavior   func()
		expectedError  error
		expectedOutput *AddProductOutput
	}{
		{
			name: "Successful product addition",
			input: AddProductInput{
				ID:          "123",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
			},
			mockBehavior: func() {
				createdAt := time.Now()
				updatedAt := createdAt
				mockService.On("AddProduct", domain.ProductID("123"), "Test Product", "A test product", 9.99).
					Return(&domain.Product{
						ID:          domain.ProductID("123"),
						Name:        "Test Product",
						Description: "A test product",
						Price:       9.99,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					}, nil)
			},
			expectedError: nil,
			expectedOutput: &AddProductOutput{
				ID:          "123",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			},
		},
		{
			name: "Error adding product",
			input: AddProductInput{
				ID:          "456",
				Name:        "Failed Product",
				Description: "This product will fail",
				Price:       19.99,
			},
			mockBehavior: func() {
				mockService.On("AddProduct", domain.ProductID("456"), "Failed Product", "This product will fail", 19.99).
					Return(nil, errors.New("failed to add product"))
			},
			expectedError:  errors.New("failed to add product"),
			expectedOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			output, err := useCase.Execute(tc.input)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, tc.expectedOutput.ID, output.ID)
				assert.Equal(t, tc.expectedOutput.Name, output.Name)
				assert.Equal(t, tc.expectedOutput.Description, output.Description)
				assert.Equal(t, tc.expectedOutput.Price, output.Price)

				// Verificar se as datas estão no formato correto
				_, err := time.Parse(time.RFC3339, output.CreatedAt)
				assert.NoError(t, err)
				_, err = time.Parse(time.RFC3339, output.UpdatedAt)
				assert.NoError(t, err)
			}

			mockService.AssertExpectations(t)
		})
	}
}
