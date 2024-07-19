package domain

import (
	"errors"
	"testing"

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

func TestUpdateProduct(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)
	service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

	productToUpdate := &Product{
		ID:          ProductID(ProductID("1")),
		Name:        "Produto 1",
		Description: "Descrição 1",
		Price:       10.0,
	}

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(productToUpdate, nil)
	mockSaveRepo.On("Save", mock.MatchedBy(func(p *Product) bool {
		return p.ID == ProductID(ProductID("1")) && p.Name == "Produto Atualizado" && p.Description == "Descrição Atualizada" && p.Price == 20.0
	})).Return(nil)

	_, err := service.UpdateProduct(ProductID(ProductID("1")), "Produto Atualizado", "Descrição Atualizada", 20.0)

	assert.Nil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockSaveRepo.AssertCalled(t, "Save", mock.MatchedBy(func(p *Product) bool {
		return p.ID == ProductID(ProductID("1")) && p.Name == "Produto Atualizado" && p.Description == "Descrição Atualizada" && p.Price == 20.0
	}))
}

func TestUpdateProductWhenFindRepositoryReturnError(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)
	service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(nil, errors.New("error"))

	_, err := service.UpdateProduct(ProductID(ProductID("1")), "Produto Atualizado", "Descrição Atualizada", 20.0)

	assert.NotNil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockSaveRepo.AssertNotCalled(t, "Save", mock.Anything)
}

func TestUpdateProductWhenFindRepositoryReturnNil(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)
	service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(nil, nil)

	_, err := service.UpdateProduct(ProductID(ProductID("1")), "Produto Atualizado", "Descrição Atualizada", 20.0)

	assert.NotNil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockSaveRepo.AssertNotCalled(t, "Save", mock.Anything)
}

func TestUpdateProductWhenSaveRepositoryReturnErr(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockSaveRepo := new(MockProductSaveRepository)
	service := NewProductService(mockSaveRepo, mockFindRepo, nil, nil)

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(&Product{ID: ProductID("1")}, nil)
	mockSaveRepo.On("Save", mock.Anything).Return(errors.New("error"))

	_, err := service.UpdateProduct(ProductID(ProductID("1")), "Produto Atualizado", "Descrição Atualizada", 20.0)

	assert.NotNil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockSaveRepo.AssertCalled(t, "Save", mock.Anything)
}

func TestDeleteProduct(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockDeleteRepo := new(MockProductDeleteRepository)
	service := NewProductService(nil, mockFindRepo, nil, mockDeleteRepo)

	productID := ProductID(ProductID("1"))

	mockFindRepo.On("Find", productID).Return(&Product{ID: productID}, nil)
	mockDeleteRepo.On("Delete", productID).Return(nil)

	err := service.DeleteProduct(productID)
	assert.Nil(t, err)
	mockFindRepo.AssertCalled(t, "Find", productID)
	mockDeleteRepo.AssertCalled(t, "Delete", productID)

	// Corrected: Reset the mock expectations for the second test case
	mockFindRepo.ExpectedCalls = nil
	mockDeleteRepo.ExpectedCalls = nil

	mockFindRepo.On("Find", productID).Return(nil, nil)

	err = service.DeleteProduct(productID)
	assert.Equal(t, errors.New("product not found"), err)
	mockFindRepo.AssertCalled(t, "Find", productID)
}

func TestDeleteProductWhenFindRepositoryReturnError(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockDeleteRepo := new(MockProductDeleteRepository)
	service := NewProductService(nil, mockFindRepo, nil, mockDeleteRepo)

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(nil, errors.New("error"))

	err := service.DeleteProduct(ProductID(ProductID("1")))
	assert.NotNil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockDeleteRepo.AssertNotCalled(t, "Delete", mock.Anything)
}

func TestDeleteProductWhenDeleteRepositoryReturnError(t *testing.T) {
	mockFindRepo := new(MockProductFindRepository)
	mockDeleteRepo := new(MockProductDeleteRepository)
	service := NewProductService(nil, mockFindRepo, nil, mockDeleteRepo)

	mockFindRepo.On("Find", ProductID(ProductID("1"))).Return(&Product{ID: ProductID("1")}, nil)
	mockDeleteRepo.On("Delete", ProductID(ProductID("1"))).Return(errors.New("error"))

	err := service.DeleteProduct(ProductID(ProductID("1")))
	assert.NotNil(t, err)
	mockFindRepo.AssertCalled(t, "Find", ProductID(ProductID("1")))
	mockDeleteRepo.AssertCalled(t, "Delete", ProductID(ProductID("1")))
}
