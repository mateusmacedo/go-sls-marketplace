package http

import (
	"bytes"
	"encoding/json"
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
		input          interface{}
		mockError      error
		expectedStatus int
		expectedBody   string
		expectExecute  bool
		method         string
	}{
		{
			name: "Successful product deletion",
			input: DeleteProductRequest{
				ID: "1",
			},
			mockError:      nil,
			expectedStatus: http.StatusNoContent,
			expectedBody:   "",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name: "Product not found",
			input: DeleteProductRequest{
				ID: "2",
			},
			mockError:      domain.ErrNotFoundProduct,
			expectedStatus: http.StatusNotFound,
			expectedBody:   domain.ErrNotFoundProduct.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name: "Invalid product ID",
			input: DeleteProductRequest{
				ID: "",
			},
			mockError:      domain.ErrInvalidProductID,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   domain.ErrInvalidProductID.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name: "Internal server error",
			input: DeleteProductRequest{
				ID: "3",
			},
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   domain.ErrRepositoryProduct.Error() + "\n",
			expectExecute:  true,
			method:         http.MethodDelete,
		},
		{
			name:           "Invalid JSON input",
			input:          "invalid json",
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid character 'i' looking for beginning of value\n",
			expectExecute:  false,
			method:         http.MethodDelete,
		},
		{
			name: "Method not allowed",
			input: DeleteProductRequest{
				ID: "1",
			},
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

			var body []byte
			var err error
			if str, ok := tt.input.(string); ok {
				body = []byte(str)
			} else {
				body, err = json.Marshal(tt.input)
				assert.NoError(t, err)
			}

			req, _ := http.NewRequest(tt.method, "/products", bytes.NewBuffer(body))
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
