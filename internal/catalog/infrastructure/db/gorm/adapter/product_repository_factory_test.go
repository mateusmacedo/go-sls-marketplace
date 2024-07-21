package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test case structure
type testCase struct {
	name         string
	createRepoFn func(map[string]interface{}) (interface{}, error)
}

func TestCreateProductRepositories(t *testing.T) {
	db, _ := setupTestDB(t)
	dependencies := map[string]interface{}{
		"db": db,
	}

	// Wrapper functions to convert the return type to interface{}
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

	// Test cases
	testCases := []testCase{
		{
			name:         "CreateProductSaveRepository",
			createRepoFn: createProductSaveRepositoryWrapper,
		},
		{
			name:         "CreateProductFindRepository",
			createRepoFn: createProductFindRepositoryWrapper,
		},
		{
			name:         "CreateProductFindAllRepository",
			createRepoFn: createProductFindAllRepositoryWrapper,
		},
		{
			name:         "CreateProductDeleteRepository",
			createRepoFn: createProductDeleteRepositoryWrapper,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, err := tc.createRepoFn(dependencies)
			assert.NoError(t, err)
			assert.NotNil(t, repo)
		})
	}
}
