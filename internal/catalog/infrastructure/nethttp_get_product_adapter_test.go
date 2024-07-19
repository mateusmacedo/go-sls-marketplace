package infrastructure

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockGetProductsUseCase struct {
	mock.Mock
}

func (m *MockGetProductsUseCase) Execute(input application.GetProductInput) (*application.GetProductOutput, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*application.GetProductOutput), args.Error(1)
}

func TestNetHTTPGetProductAdapter_Handle(t *testing.T) {
	mockUseCase := new(MockGetProductsUseCase)
	adapter := NewNetHTTPGetProductAdapter(mockUseCase)

	tests := []struct {
		name           string
		method         string
		url            string
		mockInput      application.GetProductInput
		mockOutput     *application.GetProductOutput
		mockError      error
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:           "Successful product retrieval",
			method:         "GET",
			url:            "/product?id=123",
			mockInput:      application.GetProductInput{ID: "123"},
			mockOutput:     &application.GetProductOutput{ID: "123", Name: "Test Product", Description: "A test product", Price: 9.99, CreatedAt: "2023-01-01 00:00:00", UpdatedAt: "2023-01-01 00:00:00"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: GetProductResponse{
				ID:          "123",
				Name:        "Test Product",
				Description: "A test product",
				Price:       9.99,
				CreatedAt:   "2023-01-01 00:00:00",
				UpdatedAt:   "2023-01-01 00:00:00",
			},
		},
		{
			name:           "Product not found",
			method:         "GET",
			url:            "/product?id=456",
			mockInput:      application.GetProductInput{ID: "456"},
			mockOutput:     nil,
			mockError:      domain.ErrNotFoundProduct,
			expectedStatus: http.StatusNotFound,
			expectedBody:   domain.ErrNotFoundProduct.Error() + "\n",
		},
		{
			name:           "Invalid product ID",
			method:         "GET",
			url:            "/product?id=",
			mockInput:      application.GetProductInput{ID: ""},
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   domain.ErrInvalidProductID.Error() + "\n",
		},
		{
			name:           "Method not allowed",
			method:         "POST",
			url:            "/product?id=123",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   ErrHttpMethodNotAllowed.Error() + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockInput.ID != "" {
				mockUseCase.On("Execute", tt.mockInput).Return(tt.mockOutput, tt.mockError)
			}

			var body []byte
			req, err := http.NewRequest(tt.method, tt.url, bytes.NewBuffer(body))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(adapter.Handle)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())

			mockUseCase.AssertExpectations(t)
		})
	}
}
