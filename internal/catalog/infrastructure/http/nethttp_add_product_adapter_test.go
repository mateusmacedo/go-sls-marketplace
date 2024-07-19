package http

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

type MockAddProductUseCase struct {
	mock.Mock
}

func (m *MockAddProductUseCase) Execute(input application.AddProductInput) (*application.AddProductOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*application.AddProductOutput), args.Error(1)
}

func TestNetHTTPAddProductAdapter_Handle(t *testing.T) {
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name           string
		input          interface{}
		mockOutput     *application.AddProductOutput
		mockError      error
		expectedStatus int
		expectedBody   interface{}
		expectExecute  bool
		method         string
	}{
		{
			name: "Successful product addition",
			input: AddProductRequest{
				ID:          "1",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
			},
			mockOutput: &application.AddProductOutput{
				ID:          "1",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   fixedTime.Format(time.RFC3339),
				UpdatedAt:   fixedTime.Format(time.RFC3339),
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody: AddProductResponse{
				ID:          "1",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   fixedTime.Format(time.RFC3339),
				UpdatedAt:   fixedTime.Format(time.RFC3339),
			},
			expectExecute: true,
			method:        http.MethodPost,
		},
		{
			name: "Failed product addition",
			input: AddProductRequest{
				ID:          "2",
				Name:        "Failed Product",
				Description: "A product that fails to be added",
				Price:       19.99,
			},
			mockOutput:     nil,
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   domain.ErrRepositoryProduct.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodPost,
		},
		{
			name: "Invalid input - empty name",
			input: AddProductRequest{
				ID:          "3",
				Name:        "",
				Description: "Invalid product",
				Price:       29.99,
			},
			mockOutput:     nil,
			mockError:      domain.ErrInvalidProductName,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   domain.ErrInvalidProductName.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodPost,
		},
		{
			name:           "Invalid JSON input",
			input:          "invalid json",
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid character 'i' looking for beginning of value\n",
			expectExecute:  false,
			method:         http.MethodPost,
		},
		{
			name:           "Method not allowed",
			input:          AddProductRequest{},
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   ErrHttpMethodNotAllowed.Error() + "\n",
			expectExecute:  false,
			method:         http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(MockAddProductUseCase)
			if tt.expectExecute {
				mockUseCase.On("Execute", mock.Anything).Return(tt.mockOutput, tt.mockError)
			}

			adapter := NewNetHTTPAddProductAdapter(mockUseCase)

			var body []byte
			var err error
			if str, ok := tt.input.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.input)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest(tt.method, "/products", bytes.NewReader(body))
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response AddProductResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).ID, response.ID)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).Name, response.Name)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).Description, response.Description)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).Price, response.Price)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).CreatedAt, response.CreatedAt)
				assert.Equal(t, tt.expectedBody.(AddProductResponse).UpdatedAt, response.UpdatedAt)
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
