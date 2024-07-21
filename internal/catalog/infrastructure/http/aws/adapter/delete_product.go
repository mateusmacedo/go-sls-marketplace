package adapter

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	infrahttp "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http"
)

type LambdaDeleteProductAdapter struct {
	service application.DeleteProductUseCase
	db      *dynamodb.DynamoDB
	table   string
}

func NewLambdaDeleteProductAdapter(service application.DeleteProductUseCase) *LambdaDeleteProductAdapter {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	return &LambdaDeleteProductAdapter{
		service: service,
		db:      db,
		table:   os.Getenv("PRODUCTS_TABLE"),
	}
}

func (a *LambdaDeleteProductAdapter) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != http.MethodDelete {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       `{"error": "method not allowed"}`,
		}, nil
	}

	id := request.PathParameters["id"]
	err := a.service.Execute(application.DeleteProductInput{
		ID: id,
	})

	if err != nil {
		statusCode, ok := infrahttp.HttpError[err]
		if !ok {
			statusCode = http.StatusInternalServerError
		}
		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Body:       `{"error": "` + err.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}
