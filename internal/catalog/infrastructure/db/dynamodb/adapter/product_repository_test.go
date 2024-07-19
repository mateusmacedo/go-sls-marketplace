package adapter

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/mocks"
)

func TestSaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	product := &domain.Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 100,
	}

	mockDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)

	err := repo.Save(product)
	assert.NoError(t, err)
}

func TestSaveProductErrorWhenProductIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	err := repo.Save(nil)
	assert.Error(t, err)
}

func TestSaveProductErrorWhenMarshalMapFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	product := &domain.Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 100,
	}

	mockDB.EXPECT().PutItem(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	err := repo.Save(product)
	assert.Error(t, err)
}

func TestFindProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	productID := "1"
	mockOutput := &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: productID},
			"name":  &types.AttributeValueMemberS{Value: "Product 1"},
			"price": &types.AttributeValueMemberN{Value: "100"},
		},
	}

	mockDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(mockOutput, nil)

	product, err := repo.Find(domain.ProductID(productID))
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, domain.ProductID(productID), product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100.00, product.Price)
}

func TestFindProductErrorWhenDynamoDBGetItemFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	productID := "1"

	mockDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	product, err := repo.Find(domain.ProductID(productID))
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestFindProductErrorWhenProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	productID := "1"

	mockOutput := &dynamodb.GetItemOutput{}

	mockDB.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(mockOutput, nil)

	product, err := repo.Find(domain.ProductID(productID))
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestFindAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	mockOutput := &dynamodb.ScanOutput{
		Items: []map[string]types.AttributeValue{
			{
				"id":    &types.AttributeValueMemberS{Value: "1"},
				"name":  &types.AttributeValueMemberS{Value: "Product 1"},
				"price": &types.AttributeValueMemberN{Value: "100"},
			},
			{
				"id":    &types.AttributeValueMemberS{Value: "2"},
				"name":  &types.AttributeValueMemberS{Value: "Product 2"},
				"price": &types.AttributeValueMemberN{Value: "200"},
			},
		},
	}

	mockDB.EXPECT().Scan(gomock.Any(), gomock.Any()).Return(mockOutput, nil)

	products, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
}

func TestFindAllProductsErrorWhenDynamoDBScanFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	mockDB.EXPECT().Scan(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	products, err := repo.FindAll()
	assert.Error(t, err)
	assert.Nil(t, products)
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDynamoDBAPI(ctrl)
	repo := NewDynamoDbProductRepository(mockDB, "ProductsTable")

	productID := "1"

	mockDB.EXPECT().DeleteItem(gomock.Any(), gomock.Any()).Return(&dynamodb.DeleteItemOutput{}, nil)

	err := repo.Delete(domain.ProductID(productID))
	assert.NoError(t, err)
}
