package adapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	internalhttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestNetHTTPAddProductAdapter_Handle(t *testing.T) {
	tests := []struct {
		name               string
		httpMethod         string
		requestBody        string
		mockOutput         *application.AddProductOutput
		mockError          error
		expectedStatusCode int
		expectedResponse   map[string]interface{}
	}{
		{
			name:               "Method Not Allowed",
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   map[string]interface{}{"error": pkghttp.ErrHttpMethodNotAllowed.Error()},
		},
		{
			name:               "Invalid JSON",
			httpMethod:         http.MethodPost,
			requestBody:        "invalid json",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   map[string]interface{}{"error": "invalid JSON"},
		},
		{
			name:               "Service Error",
			httpMethod:         http.MethodPost,
			requestBody:        `{"id":"1","name":"Product","description":"Description","price":10.0}`,
			mockError:          pkghttp.ErrServiceError,
			expectedStatusCode: internalhttp.HttpError[pkghttp.ErrServiceError],
			expectedResponse:   map[string]interface{}{"error": pkghttp.ErrServiceError.Error()},
		},
		{
			name:               "Success",
			httpMethod:         http.MethodPost,
			requestBody:        `{"id":"1","name":"Product","description":"Description","price":10.0}`,
			mockOutput:         &application.AddProductOutput{ID: "1", Name: "Product", Description: "Description", Price: 10.0, CreatedAt: "2021-01-01T00:00:00Z", UpdatedAt: "2021-01-02T00:00:00Z"},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   map[string]interface{}{"id": "1", "name": "Product", "description": "Description", "price": 10.0, "created_at": "2021-01-01T00:00:00Z", "updated_at": "2021-01-02T00:00:00Z"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockAddProductUseCase(mockCtrl)
			methodGuard := pkghttp.NewHttpMethodGuard([]string{http.MethodPost})

			adapter := NewNetHTTPAddProductAdapter(
				WithService(mockService),
				WithMethodGuard(methodGuard),
			)

			if tt.httpMethod == http.MethodPost && tt.requestBody != "" {
				if tt.mockError != nil {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(nil, tt.mockError).Times(1)
				} else if tt.mockOutput != nil {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(tt.mockOutput, nil).Times(1)
				}
			}

			var req *http.Request
			if tt.requestBody != "" {
				req = httptest.NewRequest(tt.httpMethod, "/products", bytes.NewBufferString(tt.requestBody))
			} else {
				req = httptest.NewRequest(tt.httpMethod, "/products", nil)
			}
			rec := httptest.NewRecorder()

			adapter.Handle(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)

			var actualBody map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&actualBody); err == nil {
				assert.Equal(t, tt.expectedResponse, actualBody)
			} else {
				assert.Equal(t, 0, len(tt.expectedResponse))
			}
		})
	}
}
