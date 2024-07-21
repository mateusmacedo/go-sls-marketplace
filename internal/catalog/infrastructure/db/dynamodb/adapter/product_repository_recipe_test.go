package adapter

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/go-sls-marketplace/test/mocks"
)

func TestCreateProductSaveRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDynamoDBAPI(ctrl)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "Products",
			},
			expectedErr: nil,
		},
		{
			name: "Missing dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     nil,
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI": mockDB,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Empty dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "",
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := CreateProductSaveRepository(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}

// Similar tests for other creation functions
func TestCreateProductFindRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDynamoDBAPI(ctrl)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "Products",
			},
			expectedErr: nil,
		},
		{
			name: "Missing dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     nil,
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI": mockDB,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Empty dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "",
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := CreateProductFindRepository(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}

func TestCreateProductFindAllRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDynamoDBAPI(ctrl)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "Products",
			},
			expectedErr: nil,
		},
		{
			name: "Missing dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     nil,
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI": mockDB,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Empty dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "",
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := CreateProductFindAllRepository(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}

func TestCreateProductDeleteRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDynamoDBAPI(ctrl)

	tests := []struct {
		name         string
		dependencies map[string]interface{}
		expectedErr  error
	}{
		{
			name: "Successful creation",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "Products",
			},
			expectedErr: nil,
		},
		{
			name: "Missing dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Nil dynamoDBAPI",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     nil,
				"dynamoTableName": "Products",
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Missing dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI": mockDB,
			},
			expectedErr: assert.AnError,
		},
		{
			name: "Empty dynamoTableName",
			dependencies: map[string]interface{}{
				"dynamoDBAPI":     mockDB,
				"dynamoTableName": "",
			},
			expectedErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := CreateProductDeleteRepository(tt.dependencies)

			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, repo)
			}
		})
	}
}
