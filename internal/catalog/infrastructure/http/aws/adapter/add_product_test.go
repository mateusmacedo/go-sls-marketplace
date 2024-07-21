package adapter

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestLambdaAddProductAdapter_Handle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockAddProductUseCase(mockCtrl)

	tests := []struct {
		name               string
		httpMethod         string
		requestBody        string
		mockServiceResult  *application.AddProductOutput
		mockServiceError   error
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "Method Not Allowed",
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   `{"error": "method not allowed"}`,
		},
		{
			name:               "Bad Request - Invalid JSON",
			httpMethod:         http.MethodPost,
			requestBody:        "{invalid-json}",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error": "invalid character 'i' looking for beginning of object key string"}`,
		},
		{
			name:               "Service Error",
			httpMethod:         http.MethodPost,
			requestBody:        `{"id":"1","name":"Product","description":"Description","price":10.0}`,
			mockServiceError:   errors.New("some service error"),
			expectedStatusCode: infrahttp.HttpError[errors.New("some service error")],
			expectedResponse:   `{"error": "some service error"}`,
		},
		{
			name:        "Success",
			httpMethod:  http.MethodPost,
			requestBody: `{"id":"1","name":"Product","description":"Description","price":10.0}`,
			mockServiceResult: &application.AddProductOutput{
				ID:          "1",
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
				CreatedAt:   "2021-01-01T00:00:00Z",
				UpdatedAt:   "2021-01-02T00:00:00Z",
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"id":"1","name":"Product","description":"Description","price":10.0,"created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-02T00:00:00Z"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockServiceResult != nil || tt.mockServiceError != nil {
				mockService.EXPECT().
					Execute(gomock.Any()).
					Return(tt.mockServiceResult, tt.mockServiceError)
			}

			adapter := NewLambdaAddProductAdapter(mockService)
			req := events.APIGatewayProxyRequest{
				HTTPMethod: tt.httpMethod,
				Body:       tt.requestBody,
			}
			resp, err := adapter.Handle(context.Background(), req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
			assert.JSONEq(t, tt.expectedResponse, resp.Body)
		})
	}
}
