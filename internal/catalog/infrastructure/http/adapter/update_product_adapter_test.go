package adapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockUpdateProductUseCase struct {
	mock.Mock
}

func (m *MockUpdateProductUseCase) Execute(input application.UpdateProductInput) (*application.UpdateProductOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*application.UpdateProductOutput), args.Error(1)
}

func TestNetHTTPUpdateProductAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name           string
		productID      string
		input          interface{}
		mockOutput     *application.UpdateProductOutput
		mockError      error
		expectedStatus int
		expectedBody   interface{}
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
				CreatedAt:   fixedTime.Format(time.RFC3339),
				UpdatedAt:   fixedTime.Format(time.RFC3339),
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: UpdateProductResponse{
				ID:          "1",
				Name:        "Updated Product",
				Description: "An updated product",
				Price:       19.99,
				CreatedAt:   fixedTime.Format(time.RFC3339),
				UpdatedAt:   fixedTime.Format(time.RFC3339),
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
			expectedBody:   domain.ErrRepositoryProduct.Error() + "\n",
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
			expectedBody:   domain.ErrInvalidProductID.Error() + "\n",
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
			expectedBody:   domain.ErrInvalidProductName.Error() + "\n",
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
			expectedBody:   "invalid character 'i' looking for beginning of value\n",
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
			expectedBody:   "method not allowed\n",
			expectExecute:  false,
			method:         http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(MockUpdateProductUseCase)
			if tt.expectExecute {
				mockUseCase.On("Execute", mock.Anything).Return(tt.mockOutput, tt.mockError)
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

			if tt.expectedStatus == http.StatusOK {
				var response UpdateProductResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).ID, response.ID)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).Name, response.Name)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).Description, response.Description)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).Price, response.Price)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).CreatedAt, response.CreatedAt)
				assert.Equal(t, tt.expectedBody.(UpdateProductResponse).UpdatedAt, response.UpdatedAt)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.expectExecute {
				mockUseCase.AssertCalled(t, "Execute", mock.Anything)
			} else {
				mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
			}
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
