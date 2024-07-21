package application

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/domain/mocks"
)

func TestGetAllProductsUseCase_Execute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFinder := mocks.NewMockAllProductFinder(mockCtrl)
	useCase := NewGetAllProductsUseCase(mockFinder)

	testCases := []struct {
		name             string
		mockBehavior     func(*mocks.MockAllProductFinder)
		expectedProducts []*GetAllProductsOutput
		expectedError    error
	}{
		{
			name: "Successful retrieval of products",
			mockBehavior: func(m *mocks.MockAllProductFinder) {
				createdAt := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
				updatedAt := time.Date(2023, 5, 2, 11, 0, 0, 0, time.UTC)
				m.EXPECT().GetAllProducts().Return([]*domain.Product{
					{
						ID:          domain.ProductID("1"),
						Name:        "Product 1",
						Description: "Description 1",
						Price:       10.0,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					},
					{
						ID:          domain.ProductID("2"),
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
					CreatedAt:   "2023-05-01T10:00:00Z",
					UpdatedAt:   "2023-05-02T11:00:00Z",
				},
				{
					ID:          "2",
					Name:        "Product 2",
					Description: "Description 2",
					Price:       20.0,
					CreatedAt:   "2023-05-01T10:00:00Z",
					UpdatedAt:   "2023-05-02T11:00:00Z",
				},
			},
			expectedError: nil,
		},
		{
			name: "Error retrieving products",
			mockBehavior: func(m *mocks.MockAllProductFinder) {
				m.EXPECT().GetAllProducts().Return(nil, domain.ErrRepositoryProduct)
			},
			expectedProducts: []*GetAllProductsOutput{},
			expectedError:    domain.ErrRepositoryProduct,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(mockFinder)

			products, err := useCase.Execute()

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				assert.Equal(t, tc.expectedProducts, products)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expectedProducts), len(products))

				for i, expectedProduct := range tc.expectedProducts {
					assert.Equal(t, expectedProduct.ID, products[i].ID)
					assert.Equal(t, expectedProduct.Name, products[i].Name)
					assert.Equal(t, expectedProduct.Description, products[i].Description)
					assert.Equal(t, expectedProduct.Price, products[i].Price)
					assert.Equal(t, expectedProduct.CreatedAt, products[i].CreatedAt)
					assert.Equal(t, expectedProduct.UpdatedAt, products[i].UpdatedAt)
				}
			}
		})
	}
}
