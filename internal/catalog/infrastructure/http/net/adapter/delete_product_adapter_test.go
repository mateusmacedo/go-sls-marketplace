package adapter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type MockDeleteProductUseCase struct {
	mock.Mock
}

func (m *MockDeleteProductUseCase) Execute(input application.DeleteProductInput) error {
	args := m.Called(input)
	return args.Error(0)
}

func TestNetHTTPDeleteProductAdapter_Handle(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		mockError      error
		expectedStatus int
		expectedBody   string
		expectExecute  bool
		method         string
	}{
		{
			name:           "Successful product deletion",
			input:          "1",
			mockError:      nil,
			expectedStatus: http.StatusNoContent,
			expectedBody:   "",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name:           "Product not found",
			input:          "2",
			mockError:      domain.ErrNotFoundProduct,
			expectedStatus: http.StatusNotFound,
			expectedBody:   domain.ErrNotFoundProduct.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name:           "Invalid product ID",
			input:          "",
			mockError:      domain.ErrInvalidProductID,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   domain.ErrInvalidProductID.Error() + "\n",
			expectExecute:  false,
			method:         http.MethodDelete,
		},
		{
			name:           "Internal server error",
			input:          "3",
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   domain.ErrRepositoryProduct.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name:           "Method not allowed",
			input:          "1",
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   ErrHttpMethodNotAllowed.Error() + "\n",
			expectExecute:  false,
			method:         http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(MockDeleteProductUseCase)
			if tt.expectExecute {
				mockUseCase.On("Execute", mock.Anything).Return(tt.mockError)
			}

			adapter := NewNetHTTPDeleteProductAdapter(mockUseCase)

			req, _ := http.NewRequest(tt.method, "/products/"+tt.input, nil)
			rr := httptest.NewRecorder()

			adapter.Handle(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())

			if tt.expectExecute {
				mockUseCase.AssertCalled(t, "Execute", mock.Anything)
			} else {
				mockUseCase.AssertNotCalled(t, "Execute", mock.Anything)
			}
		})
	}
}