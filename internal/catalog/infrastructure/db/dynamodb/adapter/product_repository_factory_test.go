package adapter

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	"github.com/mateusmacedo/go-sls-marketplace/test/mocks"
)

type testCase struct {
	name         string
	createRepoFn func(map[string]interface{}) (interface{}, error)
	expectedType interface{}
}

func TestCreateProductRepositories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mocks.NewMockDynamoDBAPI(ctrl)
	dependencies := map[string]interface{}{
		"dynamoDBAPI":     db,
		"dynamoTableName": "testTable",
	}

	createProductSaveRepositoryWrapper := func(dependencies map[string]interface{}) (interface{}, error) {
		return CreateProductSaveRepository(dependencies)
	}

	createProductFindRepositoryWrapper := func(dependencies map[string]interface{}) (interface{}, error) {
		return CreateProductFindRepository(dependencies)
	}

	createProductFindAllRepositoryWrapper := func(dependencies map[string]interface{}) (interface{}, error) {
		return CreateProductFindAllRepository(dependencies)
	}

	createProductDeleteRepositoryWrapper := func(dependencies map[string]interface{}) (interface{}, error) {
		return CreateProductDeleteRepository(dependencies)
	}

	testCases := []testCase{
		{
			name:         "CreateProductSaveRepository",
			createRepoFn: createProductSaveRepositoryWrapper,
			expectedType: (*domain.ProductSaveRepository)(nil),
		},
		{
			name:         "CreateProductFindRepository",
			createRepoFn: createProductFindRepositoryWrapper,
			expectedType: (*domain.ProductFindRepository)(nil),
		},
		{
			name:         "CreateProductFindAllRepository",
			createRepoFn: createProductFindAllRepositoryWrapper,
			expectedType: (*domain.ProductFindAllRepository)(nil),
		},
		{
			name:         "CreateProductDeleteRepository",
			createRepoFn: createProductDeleteRepositoryWrapper,
			expectedType: (*domain.ProductDeleteRepository)(nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, err := tc.createRepoFn(dependencies)
			assert.NoError(t, err)
			assert.NotNil(t, repo)
			assert.Implements(t, tc.expectedType, repo)
		})
	}
}
