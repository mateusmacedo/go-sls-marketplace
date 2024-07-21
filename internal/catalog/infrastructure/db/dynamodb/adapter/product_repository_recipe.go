package adapter

func CreateProductSaveRepository(dependencies map[string]interface{}) (interface{}, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductSaveRepository(db, tableName), nil
}

func CreateProductFindRepository(dependencies map[string]interface{}) (interface{}, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductFindRepository(db, tableName), nil
}

func CreateProductFindAllRepository(dependencies map[string]interface{}) (interface{}, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductFindAllRepository(db, tableName), nil
}

func CreateProductDeleteRepository(dependencies map[string]interface{}) (interface{}, error) {
	db := dependencies["dynamoDBAPI"].(DynamoDBAPI)
	tableName := dependencies["dynamoTableName"].(string)

	return NewDynamoDbProductDeleteRepository(db, tableName), nil
}
