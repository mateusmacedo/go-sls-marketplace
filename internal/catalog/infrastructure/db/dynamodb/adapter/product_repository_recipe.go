package adapter

import (
	"fmt"
)

func CreateProductSaveRepository(dependencies map[string]interface{}) (interface{}, error) {
	db, ok := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	if !ok || db == nil {
		return nil, fmt.Errorf("missing or invalid dynamoDBAPI dependency")
	}

	tableName, ok := dependencies["dynamoTableName"].(string)
	if !ok || tableName == "" {
		return nil, fmt.Errorf("missing or invalid dynamoTableName dependency")
	}

	return NewDynamoDbProductSaveRepository(db, tableName), nil
}

func CreateProductFindRepository(dependencies map[string]interface{}) (interface{}, error) {
	db, ok := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	if !ok || db == nil {
		return nil, fmt.Errorf("missing or invalid dynamoDBAPI dependency")
	}

	tableName, ok := dependencies["dynamoTableName"].(string)
	if !ok || tableName == "" {
		return nil, fmt.Errorf("missing or invalid dynamoTableName dependency")
	}

	return NewDynamoDbProductFindRepository(db, tableName), nil
}

func CreateProductFindAllRepository(dependencies map[string]interface{}) (interface{}, error) {
	db, ok := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	if !ok || db == nil {
		return nil, fmt.Errorf("missing or invalid dynamoDBAPI dependency")
	}

	tableName, ok := dependencies["dynamoTableName"].(string)
	if !ok || tableName == "" {
		return nil, fmt.Errorf("missing or invalid dynamoTableName dependency")
	}

	return NewDynamoDbProductFindAllRepository(db, tableName), nil
}

func CreateProductDeleteRepository(dependencies map[string]interface{}) (interface{}, error) {
	db, ok := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	if !ok || db == nil {
		return nil, fmt.Errorf("missing or invalid dynamoDBAPI dependency")
	}

	tableName, ok := dependencies["dynamoTableName"].(string)
	if !ok || tableName == "" {
		return nil, fmt.Errorf("missing or invalid dynamoTableName dependency")
	}

	return NewDynamoDbProductDeleteRepository(db, tableName), nil
}
