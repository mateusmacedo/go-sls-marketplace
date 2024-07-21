package adapter

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestNetHTTPGetAllProductsAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	tests := []struct {
		name           string
		method         string
		mockProducts   []*application.GetAllProductsOutput
		mockError      error
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:   "Successful retrieval of products",
			method: http.MethodGet,
			mockProducts: []*application.GetAllProductsOutput{
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
			expectedBody: []*GetAllProductsResponse{
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
			mockProducts:   []*application.GetAllProductsOutput{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   []*GetAllProductsResponse{},
		},
		{
			name:           "Internal server error",
			method:         http.MethodGet,
			mockProducts:   nil,
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]interface{}{"error": domain.ErrRepositoryProduct.Error()},
		},
		{
			name:           "Method not allowed",
			method:         http.MethodPost,
			mockProducts:   nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]interface{}{"error": "method not allowed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUseCase := mocks.NewMockGetAllProductsUseCase(mockCtrl)
			adapter := NewNetHTTPGetAllProductsAdapter(mockUseCase)

			if tt.method == http.MethodGet {
				mockUseCase.EXPECT().Execute().Return(tt.mockProducts, tt.mockError)
			}

			req, _ := http.NewRequest(tt.method, "/products", nil)
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				var response []*GetAllProductsResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			} else {
				var actualBody map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &actualBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, actualBody)
			}
		})
	}
}
