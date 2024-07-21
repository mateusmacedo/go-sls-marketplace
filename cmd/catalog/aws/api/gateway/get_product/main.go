package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	dynamodbadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/db/dynamodb/adapter"
	awsadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http/aws/adapter"
	"github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/log"
)

func main() {
	logger, _ := log.NewZapLogger()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		logger.Error("Error loading AWS config", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("PRODUCTS_TABLE")
	if tableName == "" {
		logger.Error("PRODUCTS_TABLE environment variable is not set", nil)
	}
	repo := dynamodbadapter.NewDynamoDbProductRepository(dynamoClient, tableName)
	service := domain.NewProductFinder(repo)
	usecase := application.NewGetProductUseCase(service)
	handler := awsadapter.NewLambdaGetProductUseCaseAdapter(usecase)

	lambda.Start(handler.Handle)
}
