package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/application/mocks"
)

func TestLambdaDeleteProductAdapter_Handle(t *testing.T) {
	tests := []struct {
		name               string
		httpMethod         string
		pathParameters     map[string]string
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
			name:               "Product Not Found",
			httpMethod:         http.MethodDelete,
			pathParameters:     map[string]string{"id": "999"},
			mockServiceError:   domain.ErrNotFoundProduct,
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   `{"error": "product not found"}`,
		},
		{
			name:               "Service Error",
			httpMethod:         http.MethodDelete,
			pathParameters:     map[string]string{"id": "1"},
			mockServiceError:   domain.ErrRepositoryProduct,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error": "error in repository"}`,
		},
		{
			name:               "Service unknown error",
			httpMethod:         http.MethodDelete,
			pathParameters:     map[string]string{"id": "1"},
			mockServiceError:   errors.New("some service error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error": "some service error"}`,
		},
		{
			name:               "Success",
			httpMethod:         http.MethodDelete,
			pathParameters:     map[string]string{"id": "1"},
			expectedStatusCode: http.StatusNoContent,
			expectedResponse:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockDeleteProductUseCase(mockCtrl)

			if tt.httpMethod == http.MethodDelete {
				if tt.mockServiceError != nil {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(tt.mockServiceError)
				} else {
					mockService.EXPECT().
						Execute(gomock.Any()).
						Return(nil)
				}
			}

			adapter := NewLambdaDeleteProductAdapter(mockService)
			req := events.APIGatewayProxyRequest{
				HTTPMethod:     tt.httpMethod,
				PathParameters: tt.pathParameters,
			}
			resp, err := adapter.Handle(context.Background(), req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
			if tt.expectedResponse != "" {
				var expectedBody map[string]interface{}
				var actualBody map[string]interface{}
				assert.NoError(t, json.Unmarshal([]byte(tt.expectedResponse), &expectedBody))
				assert.NoError(t, json.Unmarshal([]byte(resp.Body), &actualBody))
				assert.Equal(t, expectedBody, actualBody)
			} else {
				assert.Empty(t, resp.Body)
			}
		})
	}
}
