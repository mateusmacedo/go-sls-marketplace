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
	pkgapplication "github.com/mateusmacedo/go-sls-marketplace/pkg/application"
	"github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/log"
)

func main() {
	logger, _ := log.NewZapLogger()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		logger.Error("Error loading AWS config", err)
		return
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("PRODUCTS_TABLE")
	if tableName == "" {
		logger.Error("PRODUCTS_TABLE environment variable is not set", nil)
		return
	}

	serviceLocator, err := initializeServiceLocator(dynamoClient, tableName)
	if err != nil {
		logger.Error("Error initializing Service Locator", err)
		return
	}

	factory := initializeFactory(serviceLocator)

	addProductUseCase, err := factory.Create("AddProductUseCase")
	if err != nil {
		logger.Error("Error creating AddProductUseCase", err)
		return
	}

	addProductHandler := awsadapter.NewLambdaAddProductAdapter(addProductUseCase.(application.AddProductUseCase))
	lambda.Start(addProductHandler.Handle)
}

func initializeServiceLocator(dynamoClient *dynamodb.Client, tableName string) (pkgapplication.ServiceLocator, error) {
	serviceLocator := pkgapplication.NewSimpleServiceLocator()
	serviceLocator.Register("dynamoDBAPI", dynamoClient)
	serviceLocator.Register("dynamoTableName", tableName)

	// Repositories
	repositories := map[string]func(map[string]interface{}) (interface{}, error){
		"ProductSaveRepository": dynamodbadapter.CreateProductSaveRepository,
		"ProductFindRepository": dynamodbadapter.CreateProductFindRepository,
	}

	for name, factoryFunc := range repositories {
		repo, err := createRepository(factoryFunc, serviceLocator)
		if err != nil {
			return nil, err
		}
		serviceLocator.Register(name, repo)
	}

	return serviceLocator, nil
}

func createRepository(factoryFunc func(map[string]interface{}) (interface{}, error), serviceLocator pkgapplication.ServiceLocator) (interface{}, error) {
	dynamoDBAPI, err := serviceLocator.Resolve("dynamoDBAPI")
	if err != nil {
		return nil, err
	}

	dynamoTableName, err := serviceLocator.Resolve("dynamoTableName")
	if err != nil {
		return nil, err
	}

	dependencies := map[string]interface{}{
		"dynamoDBAPI":     dynamoDBAPI,
		"dynamoTableName": dynamoTableName,
	}

	return factoryFunc(dependencies)
}

func initializeFactory(serviceLocator pkgapplication.ServiceLocator) pkgapplication.Factory {
	factory := pkgapplication.NewFactory(serviceLocator)

	factory.RegisterRecipe("ProductAdder", pkgapplication.Recipe{
		Dependencies: []string{"ProductFindRepository", "ProductSaveRepository"},
		Factory:      domain.CreateProductAdder,
	})

	factory.RegisterRecipe("AddProductUseCase", pkgapplication.Recipe{
		Dependencies: []string{"ProductAdder"},
		Factory:      application.CreateAddProductUseCase,
	})

	createAndRegisterService(factory, serviceLocator, "ProductAdder")

	return factory
}

func createAndRegisterService(factory pkgapplication.Factory, serviceLocator pkgapplication.ServiceLocator, serviceName string) {
	service, err := factory.Create(serviceName)
	if err != nil {
		panic(err)
	}
	serviceLocator.Register(serviceName, service)
}
