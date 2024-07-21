package adapter

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestNetHTTPDeleteProductAdapter_Handle(t *testing.T) {
	tests := []struct {
		name               string
		httpMethod         string
		productID          string
		mockServiceError   error
		expectedStatusCode int
		expectedResponse   interface{}
	}{
		{
			name:               "Method Not Allowed",
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   map[string]interface{}{"error": pkghttp.ErrHttpMethodNotAllowed.Error()},
		},
		{
			name:               "Invalid Product ID",
			httpMethod:         http.MethodDelete,
			productID:          "",
			expectedStatusCode: infrahttp.HttpError[domain.ErrInvalidProductID],
			expectedResponse:   map[string]interface{}{"error": domain.ErrInvalidProductID.Error()},
		},
		{
			name:               "Product Not Found",
			httpMethod:         http.MethodDelete,
			productID:          "123",
			mockServiceError:   domain.ErrNotFoundProduct,
			expectedStatusCode: infrahttp.HttpError[domain.ErrNotFoundProduct],
			expectedResponse:   map[string]interface{}{"error": domain.ErrNotFoundProduct.Error()},
		},
		{
			name:               "Service Error",
			httpMethod:         http.MethodDelete,
			productID:          "123",
			mockServiceError:   pkghttp.ErrServiceError,
			expectedStatusCode: infrahttp.HttpError[pkghttp.ErrServiceError],
			expectedResponse:   map[string]interface{}{"error": pkghttp.ErrServiceError.Error()},
		},
		{
			name:               "Success",
			httpMethod:         http.MethodDelete,
			productID:          "123",
			expectedStatusCode: http.StatusNoContent,
			expectedResponse:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockDeleteProductUseCase(mockCtrl)
			adapter := NewNetHTTPDeleteProductAdapter(mockService)

			if tt.httpMethod == http.MethodDelete && tt.productID != "" {
				if tt.mockServiceError != nil {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(tt.mockServiceError).Times(1)
				} else {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(nil).Times(1)
				}
			}

			req := httptest.NewRequest(tt.httpMethod, "/products/"+tt.productID, nil)
			rec := httptest.NewRecorder()

			adapter.Handle(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)

			if tt.expectedResponse != nil {
				var actualBody map[string]interface{}
				err := json.NewDecoder(res.Body).Decode(&actualBody)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				assert.Equal(t, tt.expectedResponse, actualBody)
			} else {
				body, _ := io.ReadAll(res.Body)
				assert.Equal(t, 0, len(body))
			}
		})
	}
}
