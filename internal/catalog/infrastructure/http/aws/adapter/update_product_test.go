package adapter

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestLambdaUpdateProductUseCaseAdapter_Handle(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		productID      string
		requestBody    string
		mockOutput     *application.UpdateProductOutput
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful product update",
			method:         http.MethodPut,
			productID:      "1",
			requestBody:    `{"name":"Updated Product","description":"An updated product","price":19.99}`,
			mockOutput:     &application.UpdateProductOutput{ID: "1", Name: "Updated Product", Description: "An updated product", Price: 19.99, CreatedAt: "2021-01-01T00:00:00Z", UpdatedAt: "2021-01-02T00:00:00Z"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":"1","name":"Updated Product","description":"An updated product","price":19.99,"created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-02T00:00:00Z"}`,
		},
		{
			name:           "Failed product update",
			method:         http.MethodPut,
			productID:      "2",
			requestBody:    `{"name":"Failed Update","description":"A product that fails to be updated","price":29.99}`,
			mockOutput:     nil,
			mockError:      domain.ErrRepositoryProduct,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"error in repository"}`,
		},
		{
			name:           "Invalid input - empty product ID",
			method:         http.MethodPut,
			productID:      "",
			requestBody:    `{"name":"Invalid product","description":"An invalid product","price":39.99}`,
			mockOutput:     nil,
			mockError:      domain.ErrInvalidProductID,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid product ID"}`,
		},
		{
			name:           "Invalid input - empty name",
			method:         http.MethodPut,
			productID:      "3",
			requestBody:    `{"name":"","description":"Invalid product","price":39.99}`,
			mockOutput:     nil,
			mockError:      domain.ErrInvalidProductName,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid product name"}`,
		},
		{
			name:           "Invalid JSON input",
			method:         http.MethodPut,
			productID:      "4",
			requestBody:    `{invalid json}`,
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "invalid character 'i' looking for beginning of object key string"}`,
		},
		{
			name:           "Method not allowed",
			method:         http.MethodGet,
			productID:      "5",
			requestBody:    `{"name":"Updated Product","description":"An updated product","price":19.99}`,
			mockOutput:     nil,
			mockError:      nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   `{"error":"method not allowed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUseCase := mocks.NewMockUpdateProductUseCase(mockCtrl)
			adapter := NewLambdaUpdateProductUseCaseAdapter(mockUseCase)

			if tt.method == http.MethodPut && tt.productID != "" {
				if tt.mockError != nil {
					mockUseCase.EXPECT().
						Execute(gomock.Any()).
						Return(nil, tt.mockError).Times(1)
				} else if tt.mockOutput != nil {
					mockUseCase.EXPECT().
						Execute(gomock.Any()).
						Return(tt.mockOutput, nil).Times(1)
				}
			}

			request := events.APIGatewayProxyRequest{
				HTTPMethod: tt.method,
				PathParameters: map[string]string{
					"id": tt.productID,
				},
				Body: tt.requestBody,
			}

			response, err := adapter.Handle(context.Background(), request)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, response.StatusCode)
			assert.JSONEq(t, tt.expectedBody, response.Body)
		})
	}
}
