package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type dynamoDbProductSaveRepository struct {
	DB        DynamoDBAPI
	TableName string
}

func NewDynamoDbProductSaveRepository(db DynamoDBAPI, tableName string) domain.ProductSaveRepository {
	return &dynamoDbProductSaveRepository{DB: db, TableName: tableName}
}

func (r *dynamoDbProductSaveRepository) Save(product *domain.Product) error {
	entity, err := NewProductEntityFromDomain(product)
	if err != nil {
		return err
	}

	item, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return err
	}

	_, err = r.DB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &r.TableName,
		Item:      item,
	})
	return err
}

type dynamoDbProductFindRepository struct {
	DB        DynamoDBAPI
	TableName string
}

func NewDynamoDbProductFindRepository(db DynamoDBAPI, tableName string) domain.ProductFindRepository {
	return &dynamoDbProductFindRepository{DB: db, TableName: tableName}
}

func (r *dynamoDbProductFindRepository) Find(id domain.ProductID) (*domain.Product, error) {
	result, err := r.DB.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: string(id)},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, domain.ErrNotFoundProduct
	}

	var entity DynamoDbProductEntity
	err = attributevalue.UnmarshalMap(result.Item, &entity)
	if err != nil {
		return nil, err
	}

	return entity.ToDomain()
}

type dynamoDbProductFindAllRepository struct {
	DB        DynamoDBAPI
	TableName string
}

func NewDynamoDbProductFindAllRepository(db DynamoDBAPI, tableName string) domain.ProductFindAllRepository {
	return &dynamoDbProductFindAllRepository{DB: db, TableName: tableName}
}

func (r *dynamoDbProductFindAllRepository) FindAll() ([]*domain.Product, error) {
	result, err := r.DB.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &r.TableName,
	})

	if err != nil {
		return nil, err
	}

	var entities []*DynamoDbProductEntity
	err = attributevalue.UnmarshalListOfMaps(result.Items, &entities)
	if err != nil {
		return nil, err
	}

	products := make([]*domain.Product, 0, len(entities))
	for _, entity := range entities {
		product, _ := entity.ToDomain()
		products = append(products, product)
	}

	return products, nil
}

type dynamoDbProductDeleteRepository struct {
	DB        DynamoDBAPI
	TableName string
}

func NewDynamoDbProductDeleteRepository(db DynamoDBAPI, tableName string) domain.ProductDeleteRepository {
	return &dynamoDbProductDeleteRepository{DB: db, TableName: tableName}
}

func (r *dynamoDbProductDeleteRepository) Delete(id domain.ProductID) error {
	_, err := r.DB.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: string(id)},
		},
	})
	return err
}
