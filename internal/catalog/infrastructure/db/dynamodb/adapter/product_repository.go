package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type dynamoDbProductRepository struct {
	DB        DynamoDBAPI
	TableName string
}

func NewDynamoDbProductRepository(db DynamoDBAPI, tableName string) domain.ProductRepository {
	return &dynamoDbProductRepository{DB: db, TableName: tableName}
}

func (r *dynamoDbProductRepository) Save(product *domain.Product) error {
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

func (r *dynamoDbProductRepository) Find(id domain.ProductID) (*domain.Product, error) {
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

func (r *dynamoDbProductRepository) FindAll() ([]*domain.Product, error) {
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

func (r *dynamoDbProductRepository) Delete(id domain.ProductID) error {
	_, err := r.DB.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: string(id)},
		},
	})
	return err
}
