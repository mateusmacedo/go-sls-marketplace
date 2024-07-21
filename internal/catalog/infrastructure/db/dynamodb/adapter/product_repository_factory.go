package adapter

import "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"

func CreateProductSaveRepository(dependencies map[string]interface{}) (domain.ProductSaveRepository, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductSaveRepository(db, tableName), nil
}

func CreateProductFindRepository(dependencies map[string]interface{}) (domain.ProductFindRepository, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductFindRepository(db, tableName), nil
}

func CreateProductFindAllRepository(dependencies map[string]interface{}) (domain.ProductFindAllRepository, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductFindAllRepository(db, tableName), nil
}

func CreateProductDeleteRepository(dependencies map[string]interface{}) (domain.ProductDeleteRepository, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductDeleteRepository(db, tableName), nil
}
