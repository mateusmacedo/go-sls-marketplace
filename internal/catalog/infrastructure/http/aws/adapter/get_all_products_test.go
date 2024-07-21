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

func TestLambdaGetAllProductsAdapter_Handle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mocks.NewMockGetAllProductsUseCase(mockCtrl)

	tests := []struct {
		name               string
		httpMethod         string
		mockServiceResult  []*application.GetAllProductsOutput
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
			name:               "Service Error",
			httpMethod:         http.MethodGet,
			mockServiceError:   errors.New("some service error"),
			expectedStatusCode: infrahttp.HttpError[errors.New("some service error")],
			expectedResponse:   `{"error": "some service error"}`,
		},
		{
			name:       "Success",
			httpMethod: http.MethodGet,
			mockServiceResult: []*application.GetAllProductsOutput{
				{
					ID:          "1",
					Name:        "Product 1",
					Description: "Description 1",
					Price:       10.0,
					CreatedAt:   "2021-01-01T00:00:00Z",
					UpdatedAt:   "2021-01-02T00:00:00Z",
				},
				{
					ID:          "2",
					Name:        "Product 2",
					Description: "Description 2",
					Price:       20.0,
					CreatedAt:   "2021-02-01T00:00:00Z",
					UpdatedAt:   "2021-02-02T00:00:00Z",
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"id":"1","name":"Product 1","description":"Description 1","price":10.0,"created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-02T00:00:00Z"},{"id":"2","name":"Product 2","description":"Description 2","price":20.0,"created_at":"2021-02-01T00:00:00Z","updated_at":"2021-02-02T00:00:00Z"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockServiceResult != nil || tt.mockServiceError != nil {
				mockService.EXPECT().
					Execute().
					Return(tt.mockServiceResult, tt.mockServiceError)
			}

			adapter := NewLambdaGetAllProductsAdapter(mockService)
			req := events.APIGatewayProxyRequest{
				HTTPMethod: tt.httpMethod,
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