package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductSaveRepository(t *testing.T) {
	db, _ := setupTestDB(t)

	dependencies := map[string]interface{}{
		"db": db,
	}

	repo, err := CreateProductSaveRepository(dependencies)

	assert.NoError(t, err)
	assert.NotNil(t, repo)
}

func TestCreateProductFindRepository(t *testing.T) {
	db, _ := setupTestDB(t)

	dependencies := map[string]interface{}{
		"db": db,
	}

	repo, err := CreateProductFindRepository(dependencies)

	assert.NoError(t, err)
	assert.NotNil(t, repo)
}

func TestCreateProductFindAllRepository(t *testing.T) {
	db, _ := setupTestDB(t)

	dependencies := map[string]interface{}{
		"db": db,
	}

	repo, err := CreateProductFindAllRepository(dependencies)

	assert.NoError(t, err)
	assert.NotNil(t, repo)
}

func TestCreateProductDeleteRepository(t *testing.T) {
	db, _ := setupTestDB(t)

	dependencies := map[string]interface{}{
		"db": db,
	}

	repo, err := CreateProductDeleteRepository(dependencies)

	assert.NoError(t, err)
	assert.NotNil(t, repo)
}
