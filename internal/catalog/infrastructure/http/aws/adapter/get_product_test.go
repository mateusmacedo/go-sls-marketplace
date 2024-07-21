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
	httperror "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http/error"
	httpadapter "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http/adapter"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestLambdaGetProductUseCaseAdapter_Handle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockGetProductUseCase(mockCtrl)

	tests := []struct {
		name               string
		httpMethod         string
		pathParameters     map[string]string
		mockServiceResult  *application.GetProductOutput
		mockServiceError   error
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "Method Not Allowed",
			httpMethod:         http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   `{"error": "method not allowed"}`,
		},
		{
			name:               "Product Not Found",
			httpMethod:         http.MethodGet,
			pathParameters:     map[string]string{"id": "999"},
			mockServiceError:   domain.ErrNotFoundProduct,
			expectedStatusCode: httperror.HttpError[domain.ErrNotFoundProduct],
			expectedResponse:   `{"error": "product not found"}`,
		},
		{
			name:               "Service Error",
			httpMethod:         http.MethodGet,
			pathParameters:     map[string]string{"id": "1"},
			mockServiceError:   httpadapter.ErrServiceError,
			expectedStatusCode: httperror.HttpError[httpadapter.ErrServiceError],
			expectedResponse:   `{"error": "some service error"}`,
		},
		{
			name:           "Success",
			httpMethod:     http.MethodGet,
			pathParameters: map[string]string{"id": "1"},
			mockServiceResult: &application.GetProductOutput{
				ID:          "1",
				Name:        "Product",
				Description: "Description",
				Price:       10.0,
				CreatedAt:   "2021-01-01T00:00:00Z",
				UpdatedAt:   "2021-01-02T00:00:00Z",
			},
			expectedStatusCode: http.StatusOK,
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

			adapter := NewLambdaGetProductUseCaseAdapter(mockService)
			req := events.APIGatewayProxyRequest{
				HTTPMethod:     tt.httpMethod,
				PathParameters: tt.pathParameters,
			}
			resp, err := adapter.Handle(context.Background(), req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
			if tt.expectedResponse != "" {
				assert.JSONEq(t, tt.expectedResponse, resp.Body)
			}
		})
	}
}
