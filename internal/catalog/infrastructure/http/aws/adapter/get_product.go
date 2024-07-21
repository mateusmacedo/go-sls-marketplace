package adapter

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	_http "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
)

type GetProductUseCaseResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type LambdaGetProductUseCaseAdapter struct {
	service application.GetProductUseCase
	db      *dynamodb.DynamoDB
	table   string
}

func NewLambdaGetProductUseCaseAdapter(service application.GetProductUseCase) *LambdaGetProductUseCaseAdapter {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	return &LambdaGetProductUseCaseAdapter{
		service: service,
		db:      db,
		table:   os.Getenv("PRODUCTS_TABLE"),
	}
}

func (a *LambdaGetProductUseCaseAdapter) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != http.MethodGet {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       `{"error": "method not allowed"}`,
		}, nil
	}

	id := request.PathParameters["id"]

	product, err := a.service.Execute(application.GetProductInput{
		ID: id,
	})
	if err != nil {
		statusCode, ok := _http.HttpError[err]
		if !ok {
			statusCode = http.StatusInternalServerError
		}
		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Body:       `{"error": "` + err.Error() + `"}`,
		}, nil
	}

	res := GetProductUseCaseResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	responseBody, err := json.Marshal(res) // TODO: Test error
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "` + err.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}
