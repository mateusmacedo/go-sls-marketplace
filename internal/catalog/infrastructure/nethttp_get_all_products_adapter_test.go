package infrastructure

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockGetAllProductsUseCase struct {
	mock.Mock
}

func (m *MockGetAllProductsUseCase) Execute() ([]*domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func TestNetHTTPGetAllProductsAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name           string
		method         string
		mockProducts   []*domain.Product
		mockError      error
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:   "Successful retrieval of products",
			method: http.MethodGet,
			mockProducts: []*domain.Product{
				{
					ID:          "1",
					Name:        "Test Product 1",
					Description: "A test product",
					Price:       9.99,
					CreatedAt:   fixedTime,
					UpdatedAt:   fixedTime,
				},
				{
					ID:          "2",
					Name:        "Test Product 2",
					Description: "Another test product",
					Price:       19.99,
					CreatedAt:   fixedTime,
					UpdatedAt:   fixedTime,
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: []ProductResponse{
				{
					ID:          "1",
					Name:        "Test Product 1",
					Description: "A test product",
					Price:       9.99,
					CreatedAt:   fixedTime,
					UpdatedAt:   fixedTime,
				},
				{
					ID:          "2",
					Name:        "Test Product 2",
					Description: "Another test product",
					Price:       19.99,
					CreatedAt:   fixedTime,
					UpdatedAt:   fixedTime,
				},
			},
		},
		{
			name:           "Empty product list",
			method:         http.MethodGet,
			mockProducts:   []*domain.Product{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   []ProductResponse{},
		},
		{
			name:           "Internal server error",
			method:         http.MethodGet,
			mockProducts:   nil,
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal server error\n",
		},
		{
			name:           "Method not allowed",
			method:         http.MethodPost,
			mockProducts:   nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(MockGetAllProductsUseCase)
			if tt.method == http.MethodGet {
				mockUseCase.On("Execute").Return(tt.mockProducts, tt.mockError)
			}

			adapter := NewNetHTTPGetAllProductsAdapter(mockUseCase)

			req, _ := http.NewRequest(tt.method, "/products", nil)
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []ProductResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.method == http.MethodGet {
				mockUseCase.AssertCalled(t, "Execute")
			} else {
				mockUseCase.AssertNotCalled(t, "Execute")
			}
		})
	}
}
