package adapter

import (
	"bytes"
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

func TestNetHTTPUpdateProductAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	tests := []struct {
		name           string
		productID      string
		input          interface{}
		mockOutput     *application.UpdateProductOutput
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
		expectExecute  bool
		method         string
	}{
		{
			name:      "Successful product update",
			productID: "1",
			input: UpdateProductRequest{
				Name:        stringPtr("Updated Product"),
				Description: stringPtr("An updated product"),
				Price:       float64Ptr(19.99),
			},
			mockOutput: &application.UpdateProductOutput{
				ID:          "1",
				Name:        "Updated Product",
				Description: "An updated product",
				Price:       19.99,
				CreatedAt:   fixedTime,
				UpdatedAt:   fixedTime,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"id":          "1",
				"name":        "Updated Product",
				"description": "An updated product",
				"price":       19.99,
				"created_at":  fixedTime,
				"updated_at":  fixedTime,
			},
			expectExecute: true,
			method:        http.MethodPut,
		},
		{
			name:      "Failed product update",
			productID: "2",
			input: UpdateProductRequest{
				Name:        stringPtr("Failed Update"),
				Description: stringPtr("A product that fails to be updated"),
				Price:       float64Ptr(29.99),
			},
			mockOutput:     nil,
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   map[string]interface{}{"error": domain.ErrRepositoryProduct.Error()},
			expectExecute:  true,
			method:         http.MethodPut,
		},
		{
			name:      "Invalid input - empty product ID",
			productID: "",
			input: UpdateProductRequest{
				Name:        stringPtr("Invalid product"),
				Description: stringPtr("An invalid product"),
				Price:       float64Ptr(39.99),
			},
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": domain.ErrInvalidProductID.Error()},
			expectExecute:  false,
			method:         http.MethodPut,
		},
		{
			name:      "Invalid input - empty name",
			productID: "3",
			input: UpdateProductRequest{
				Name:        stringPtr(""),
				Description: stringPtr("Invalid product"),
				Price:       float64Ptr(39.99),
			},
			mockOutput:     nil,
			mockError:      domain.ErrInvalidProductName,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": domain.ErrInvalidProductName.Error()},
			expectExecute:  true,
			method:         http.MethodPut,
		},
		{
			name:           "Invalid JSON input",
			productID:      "4",
			input:          "invalid json",
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "invalid character 'i' looking for beginning of value"},
			expectExecute:  false,
			method:         http.MethodPut,
		},
		{
			name:           "Method not allowed",
			productID:      "5",
			input:          UpdateProductRequest{},
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]interface{}{"error": "method not allowed"},
			expectExecute:  false,
			method:         http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mocks.NewMockUpdateProductUseCase(ctrl)
			if tt.expectExecute {
				mockUseCase.EXPECT().Execute(gomock.Any()).Return(tt.mockOutput, tt.mockError)
			} else {
				mockUseCase.EXPECT().Execute(gomock.Any()).Times(0)
			}

			adapter := NewNetHTTPUpdateProductAdapter(mockUseCase)

			var body []byte
			var err error
			if str, ok := tt.input.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.input)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest(tt.method, "/products/"+tt.productID, bytes.NewReader(body))
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var responseBody map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedBody, responseBody)

		})
	}
}

// Helper functions to create pointers for UpdateProductRequest
func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}
