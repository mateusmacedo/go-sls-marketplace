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
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
)

type UpdateProductUseCaseRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type UpdateProductUseCaseResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type LambdaUpdateProductUseCaseAdapter struct {
	service application.UpdateProductUseCase
	db      *dynamodb.DynamoDB
	table   string
}

func NewLambdaUpdateProductUseCaseAdapter(service application.UpdateProductUseCase) *LambdaUpdateProductUseCaseAdapter {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	return &LambdaUpdateProductUseCaseAdapter{
		service: service,
		db:      db,
		table:   os.Getenv("PRODUCTS_TABLE"),
	}
}

func (a *LambdaUpdateProductUseCaseAdapter) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != http.MethodPut && request.HTTPMethod != http.MethodPatch {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       pkghttp.ErrHttpMethodNotAllowed.Error(),
		}, nil
	}

	id := request.PathParameters["id"]
	var req UpdateProductUseCaseRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	product, err := a.service.Execute(application.UpdateProductInput{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: infrahttp.HttpError[err],
			Body:       err.Error(),
		}, nil
	}

	res := UpdateProductUseCaseResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	responseBody, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}
