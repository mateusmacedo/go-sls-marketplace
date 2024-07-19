package domain

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductSaveRepository struct {
	mock.Mock
}

func (m *MockProductSaveRepository) Save(product *Product) error {
	args := m.Called(product)
	return args.Error(0)
}

type MockProductFindRepository struct {
	mock.Mock
}

func (m *MockProductFindRepository) Find(id ProductID) (*Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*Product), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockProductFindAllRepository struct {
	mock.Mock
}

func (m *MockProductFindAllRepository) FindAll() ([]*Product, error) {
	args := m.Called()
	return args.Get(0).([]*Product), args.Error(1)
}

type MockProductDeleteRepository struct {
	mock.Mock
}

func (m *MockProductDeleteRepository) Delete(id ProductID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestAddProduct(t *testing.T) {
	t.Run("Successful addition", func(t *testing.T) {
		mockSaveRepo := new(MockProductSaveRepository)
		mockFindRepo := new(MockProductFindRepository)
		mockFindAllRepo := new(MockProductFindAllRepository)
		mockDeleteRepo := new(MockProductDeleteRepository)

		mockFindRepo.On("Find", mock.Anything).Return(nil, nil)
		mockSaveRepo.On("Save", mock.Anything).Return(nil)

		service := NewProductService(mockSaveRepo, mockFindRepo, mockFindAllRepo, mockDeleteRepo)

		product, err := service.AddProduct("1", "Produto Teste", "Descrição Teste", 10.0)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		mockFindRepo.AssertCalled(t, "Find", mock.Anything)
		mockSaveRepo.AssertCalled(t, "Save", mock.Anything)
	})

	t.Run("Find repository returns error", func(t *testing.T) {
		mockSaveRepo := new(MockProductSaveRepository)
		mockFindRepo := new(MockProductFindRepository)
		mockFindAllRepo := new(MockProductFindAllRepository)
		mockDeleteRepo := new(MockProductDeleteRepository)

		mockFindRepo.On("Find", mock.Anything).Return(nil, errors.New("error"))

		service := NewProductService(mockSaveRepo, mockFindRepo, mockFindAllRepo, mockDeleteRepo)

		_, err := service.AddProduct("1", "Produto Teste", "Descrição Teste", 10.0)

		assert.NotNil(t, err)
		mockFindRepo.AssertCalled(t, "Find", mock.Anything)
		mockSaveRepo.AssertNotCalled(t, "Save", mock.Anything)
	})

	t.Run("Product already exists", func(t *testing.T) {
		mockSaveRepo := new(MockProductSaveRepository)
		mockFindRepo := new(MockProductFindRepository)
		mockFindAllRepo := new(MockProductFindAllRepository)
		mockDeleteRepo := new(MockProductDeleteRepository)

		mockFindRepo.On("Find", mock.Anything).Return(&Product{}, nil)

		service := NewProductService(mockSaveRepo, mockFindRepo, mockFindAllRepo, mockDeleteRepo)

		_, err := service.AddProduct("1", "Produto Teste", "Descrição Teste", 10.0)

		assert.NotNil(t, err)
		mockFindRepo.AssertCalled(t, "Find", mock.Anything)
		mockSaveRepo.AssertNotCalled(t, "Save", mock.Anything)
	})

	t.Run("Invalid product data", func(t *testing.T) {
		mockSaveRepo := new(MockProductSaveRepository)
		mockFindRepo := new(MockProductFindRepository)
		mockFindAllRepo := new(MockProductFindAllRepository)
		mockDeleteRepo := new(MockProductDeleteRepository)

		mockFindRepo.On("Find", mock.Anything).Return(nil, nil)

		service := NewProductService(mockSaveRepo, mockFindRepo, mockFindAllRepo, mockDeleteRepo)

		_, err := service.AddProduct("", "", "Descrição Teste", -10.0)

		assert.NotNil(t, err)
		mockFindRepo.AssertCalled(t, "Find", mock.Anything)
		mockSaveRepo.AssertNotCalled(t, "Save", mock.Anything)
	})

	t.Run("Save repository returns error", func(t *testing.T) {
		mockSaveRepo := new(MockProductSaveRepository)
		mockFindRepo := new(MockProductFindRepository)
		mockFindAllRepo := new(MockProductFindAllRepository)
		mockDeleteRepo := new(MockProductDeleteRepository)

		mockFindRepo.On("Find", mock.Anything).Return(nil, nil)
		mockSaveRepo.On("Save", mock.Anything).Return(errors.New("error"))

		service := NewProductService(mockSaveRepo, mockFindRepo, mockFindAllRepo, mockDeleteRepo)

		_, err := service.AddProduct("1", "Produto Teste", "Descrição Teste", 10.0)

		assert.NotNil(t, err)
		mockFindRepo.AssertCalled(t, "Find", mock.Anything)
		mockSaveRepo.AssertCalled(t, "Save", mock.Anything)
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("Successful retrieval", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		service := NewProductService(nil, mockFindRepo, nil, nil)

		expectedProduct := &Product{
			ID:          ProductID("1"),
			Name:        "Produto Teste",
			Description: "Descrição Teste",
			Price:       10.0,
		}

		mockFindRepo.On("Find", ProductID("1")).Return(expectedProduct, nil)

		product, err := service.GetProduct(ProductID("1"))

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, product)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
	})

	t.Run("Product not found", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		service := NewProductService(nil, mockFindRepo, nil, nil)

		mockFindRepo.On("Find", ProductID("1")).Return(nil, errors.New("product not found"))

		product, err := service.GetProduct(ProductID("1"))

		assert.Nil(t, product)
		assert.Equal(t, errors.New("product not found"), err)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
	})

	t.Run("Repository error", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		service := NewProductService(nil, mockFindRepo, nil, nil)

		mockFindRepo.On("Find", ProductID("1")).Return(nil, errors.New("database error"))

		product, err := service.GetProduct(ProductID("1"))

		assert.Nil(t, product)
		assert.Equal(t, errors.New("database error"), err)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
	})

	t.Run("Invalid product ID", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		service := NewProductService(nil, mockFindRepo, nil, nil)

		product, err := service.GetProduct("")

		assert.Nil(t, product)
		assert.Equal(t, errors.New("invalid product ID"), err)
		mockFindRepo.AssertNotCalled(t, "Find")
	})
}

func TestProductService_GetAllProducts(t *testing.T) {
	testCases := []struct {
		name           string
		setupMock      func(*MockProductFindAllRepository)
		expectedResult []*Product
		expectedError  string
	}{
		{
			name: "Successful retrieval of products",
			setupMock: func(mockRepo *MockProductFindAllRepository) {
				mockRepo.On("FindAll").Return([]*Product{
					{ID: ProductID("1"), Name: "Product 1", Description: "Description 1", Price: 10.0},
					{ID: ProductID("2"), Name: "Product 2", Description: "Description 2", Price: 20.0},
				}, nil)
			},
			expectedResult: []*Product{
				{ID: ProductID("1"), Name: "Product 1", Description: "Description 1", Price: 10.0},
				{ID: ProductID("2"), Name: "Product 2", Description: "Description 2", Price: 20.0},
			},
			expectedError: "",
		},
		{
			name: "Empty product list",
			setupMock: func(mockRepo *MockProductFindAllRepository) {
				mockRepo.On("FindAll").Return([]*Product{}, nil)
			},
			expectedResult: []*Product{},
			expectedError:  "",
		},
		{
			name: "Database error",
			setupMock: func(mockRepo *MockProductFindAllRepository) {
				mockRepo.On("FindAll").Return([]*Product{}, errors.New("database connection error"))
			},
			expectedResult: nil,
			expectedError:  "database connection error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFindRepo := new(MockProductFindAllRepository)
			tc.setupMock(mockFindRepo)

			service := NewProductService(nil, nil, mockFindRepo, nil)

			result, err := service.GetAllProducts()

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}

			mockFindRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Successful update", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		mockSaveRepo := new(MockProductSaveRepository)
		service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

		existingProduct := &Product{
			ID:          ProductID("1"),
			Name:        "Produto Original",
			Description: "Descrição Original",
			Price:       10.0,
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now().Add(-24 * time.Hour),
		}

		mockFindRepo.On("Find", ProductID("1")).Return(existingProduct, nil)
		mockSaveRepo.On("Save", mock.Anything).Return(nil)

		result, err := service.UpdateProduct(ProductID("1"), "Produto Atualizado", "Descrição Atualizada", 15.0)

		assert.Nil(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, ProductID("1"), result.ID)
		assert.Equal(t, "Produto Atualizado", result.Name)
		assert.Equal(t, "Descrição Atualizada", result.Description)
		assert.Equal(t, 15.0, result.Price)

		assert.Equal(t, result.CreatedAt, existingProduct.CreatedAt)

		assert.Equal(t, result.UpdatedAt, existingProduct.UpdatedAt)

		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
		mockSaveRepo.AssertCalled(t, "Save", mock.Anything)
	})

	t.Run("Product not found", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		mockSaveRepo := new(MockProductSaveRepository)
		service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

		mockFindRepo.On("Find", ProductID("1")).Return(nil, nil)

		result, err := service.UpdateProduct(ProductID("1"), "Produto Atualizado", "Descrição Atualizada", 15.0)

		assert.Nil(t, result)
		assert.Equal(t, errors.New("product not found"), err)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
		mockSaveRepo.AssertNotCalled(t, "Save")
	})

	t.Run("Find repository error", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		mockSaveRepo := new(MockProductSaveRepository)
		service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

		mockFindRepo.On("Find", ProductID("1")).Return(nil, errors.New("database error"))

		result, err := service.UpdateProduct(ProductID("1"), "Produto Atualizado", "Descrição Atualizada", 15.0)

		assert.Nil(t, result)
		assert.Equal(t, errors.New("database error"), err)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
		mockSaveRepo.AssertNotCalled(t, "Save")
	})

	t.Run("Save repository error", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		mockSaveRepo := new(MockProductSaveRepository)
		service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

		existingProduct := &Product{
			ID:          ProductID("1"),
			Name:        "Produto Original",
			Description: "Descrição Original",
			Price:       10.0,
		}

		mockFindRepo.On("Find", ProductID("1")).Return(existingProduct, nil)
		mockSaveRepo.On("Save", mock.Anything).Return(errors.New("save error"))

		result, err := service.UpdateProduct(ProductID("1"), "Produto Atualizado", "Descrição Atualizada", 15.0)

		assert.Nil(t, result)
		assert.Equal(t, errors.New("save error"), err)
		mockFindRepo.AssertCalled(t, "Find", ProductID("1"))
		mockSaveRepo.AssertCalled(t, "Save", mock.Anything)
	})

	t.Run("Invalid product data", func(t *testing.T) {
		mockFindRepo := new(MockProductFindRepository)
		existingProduct := &Product{
			ID:          ProductID("1"),
			Name:        "Produto Original",
			Description: "Descrição Original",
			Price:       10.0,
		}
		mockFindRepo.On("Find", ProductID("1")).Return(existingProduct, nil)

		mockSaveRepo := new(MockProductSaveRepository)

		service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

		testCases := []struct {
			name        string
			id          ProductID
			newName     string
			newDesc     string
			newPrice    float64
			expectedErr string
		}{
			{
				name:        "Empty name",
				id:          ProductID("1"),
				newName:     "",
				newDesc:     "Valid description",
				newPrice:    10.0,
				expectedErr: "invalid product name",
			},
			{
				name:        "Empty description",
				id:          ProductID("1"),
				newName:     "Valid name",
				newDesc:     "",
				newPrice:    10.0,
				expectedErr: "invalid product description",
			},
			{
				name:        "Negative price",
				id:          ProductID("1"),
				newName:     "Valid name",
				newDesc:     "Valid description",
				newPrice:    -10.0,
				expectedErr: "invalid product price",
			},
			{
				name:        "Empty ID",
				id:          ProductID(""),
				newName:     "Valid name",
				newDesc:     "Valid description",
				newPrice:    10.0,
				expectedErr: "invalid product ID",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := service.UpdateProduct(tc.id, tc.newName, tc.newDesc, tc.newPrice)

				assert.Nil(t, result)
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
				mockFindRepo.AssertNotCalled(t, "Find")
				mockSaveRepo.AssertNotCalled(t, "Save")
			})
		}
	})
}

func TestProductService_DeleteProduct(t *testing.T) {
	testCases := []struct {
		name          string
		productID     ProductID
		setupMock     func(*MockProductFindRepository, *MockProductDeleteRepository)
		expectedError string
	}{
		{
			name:      "Successful product deletion",
			productID: ProductID("1"),
			setupMock: func(mockFindRepo *MockProductFindRepository, mockDeleteRepo *MockProductDeleteRepository) {
				mockFindRepo.On("Find", ProductID("1")).Return(&Product{ID: ProductID("1")}, nil)
				mockDeleteRepo.On("Delete", ProductID("1")).Return(nil)
			},
			expectedError: "",
		},
		{
			name:      "Product not found",
			productID: ProductID("2"),
			setupMock: func(mockFindRepo *MockProductFindRepository, mockDeleteRepo *MockProductDeleteRepository) {
				mockFindRepo.On("Find", ProductID("2")).Return(nil, nil)
			},
			expectedError: "product not found",
		},
		{
			name:      "Database error during find",
			productID: ProductID("3"),
			setupMock: func(mockFindRepo *MockProductFindRepository, mockDeleteRepo *MockProductDeleteRepository) {
				mockFindRepo.On("Find", ProductID("3")).Return(nil, errors.New("database error"))
			},
			expectedError: "database error",
		},
		{
			name:      "Database error during delete",
			productID: ProductID("4"),
			setupMock: func(mockFindRepo *MockProductFindRepository, mockDeleteRepo *MockProductDeleteRepository) {
				mockFindRepo.On("Find", ProductID("4")).Return(&Product{ID: ProductID("4")}, nil)
				mockDeleteRepo.On("Delete", ProductID("4")).Return(errors.New("delete error"))
			},
			expectedError: "delete error",
		},
		{
			name:      "Invalid product ID",
			productID: ProductID(""),
			setupMock: func(mockFindRepo *MockProductFindRepository, mockDeleteRepo *MockProductDeleteRepository) {
				// No setup needed for this case
			},
			expectedError: "invalid product ID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockFindRepo := new(MockProductFindRepository)
			mockDeleteRepo := new(MockProductDeleteRepository)
			tc.setupMock(mockFindRepo, mockDeleteRepo)

			service := NewProductService(nil, mockFindRepo, nil, mockDeleteRepo)

			err := service.DeleteProduct(tc.productID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockFindRepo.AssertExpectations(t)
			mockDeleteRepo.AssertExpectations(t)
		})
	}
}
