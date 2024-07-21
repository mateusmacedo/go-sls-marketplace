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
	_adapter "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http/adapter"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestNetHTTPGetProductAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	tests := []struct {
		name           string
		method         string
		productID      string
		mockProduct    *application.GetProductOutput
		mockError      error
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:      "Successful product retrieval",
			method:    http.MethodGet,
			productID: "123",
			mockProduct: &application.GetProductOutput{
				ID:          "123",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   fixedTime,
				UpdatedAt:   fixedTime,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: &GetProductResponse{
				ID:          "123",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   fixedTime,
				UpdatedAt:   fixedTime,
			},
		},
		{
			name:           "Product not found",
			method:         http.MethodGet,
			productID:      "123",
			mockProduct:    nil,
			mockError:      domain.ErrNotFoundProduct,
			expectedStatus: http.StatusNotFound,
			expectedBody:   map[string]interface{}{"error": domain.ErrNotFoundProduct.Error()},
		},
		{
			name:           "Method not allowed",
			method:         http.MethodPost,
			productID:      "123",
			mockProduct:    nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]interface{}{"error": _adapter.ErrHttpMethodNotAllowed.Error()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUseCase := mocks.NewMockGetProductUseCase(mockCtrl)
			if tt.method == http.MethodGet {
				mockUseCase.EXPECT().Execute(application.GetProductInput{ID: tt.productID}).Return(tt.mockProduct, tt.mockError)
			}

			adapter := NewNetHTTPGetProductAdapter(mockUseCase)

			req, _ := http.NewRequest(tt.method, "/products/"+tt.productID, nil)
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				var response GetProductResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, &response)
			} else {
				var actualBody map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &actualBody)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, actualBody)
			}
		})
	}
}
