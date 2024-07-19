package adapter

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestGormProductRepository_Save(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	product := &domain.Product{
		ID:          "test-id",
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "product_entities" (.+) VALUES (.+) RETURNING "id"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(product.ID))
	mock.ExpectCommit()

	err := repo.Save(product)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_Save_Error_WhenProductIsNil(t *testing.T) {
	gormDB, _ := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	err := repo.Save(nil)
	assert.ErrorIs(t, err, domain.ErrInvalidProductID)
}

func TestGormProductRepository_Find(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	productID := domain.ProductID("test-id")

	// Create a fixed time for testing
	fixedTime := time.Now().Round(time.Second)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow("test-id", "Test Product", "Test Description", 9.99, fixedTime, fixedTime)

	// Adjust the SQL query expectation to match GORM's actual behavior
	mock.ExpectQuery(`SELECT \* FROM "product_entities" WHERE id = \$1 ORDER BY "product_entities"."id" LIMIT \$2`).
		WithArgs(string(productID), 1).
		WillReturnRows(rows)

	product, err := repo.Find(productID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	if product != nil {
		assert.Equal(t, string(productID), string(product.ID))
		assert.Equal(t, "Test Product", product.Name)
		assert.Equal(t, "Test Description", product.Description)
		assert.Equal(t, 9.99, product.Price)
		assert.Equal(t, fixedTime, product.CreatedAt)
		assert.Equal(t, fixedTime, product.UpdatedAt)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_Find_Error_WhenGormError(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	productID := domain.ProductID("test-id")

	// Adjust the SQL query expectation to match GORM's actual behavior
	mock.ExpectQuery(`SELECT \* FROM "product_entities" WHERE id = \$1 ORDER BY "product_entities"."id" LIMIT \$2`).
		WithArgs(string(productID), 1).
		WillReturnError(errors.New("unexpected error"))

	product, err := repo.Find(productID)
	assert.Nil(t, product)
	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_FindAll(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow("test-id-1", "Test Product 1", "Test Description 1", 9.99, time.Now(), time.Now()).
		AddRow("test-id-2", "Test Product 2", "Test Description 2", 19.99, time.Now(), time.Now())

	mock.ExpectQuery("SELECT \\* FROM \"product_entities\"").
		WillReturnRows(rows)

	products, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, products, 2)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_FindAll_Error_WhenGormError(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	mock.ExpectQuery("SELECT \\* FROM \"product_entities\"").
		WillReturnError(errors.New("unexpected error"))

	products, err := repo.FindAll()
	assert.Nil(t, products)
	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_Delete(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	productID := domain.ProductID("test-id")

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"product_entities\" WHERE id = \\$1").
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Delete(productID)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGormProductRepository_Find_NotFound(t *testing.T) {
	gormDB, mock := setupTestDB(t)
	repo := NewGormProductRepository(gormDB)

	productID := domain.ProductID("non-existent-id")

	// Adjust the SQL query expectation to match GORM's actual behavior
	mock.ExpectQuery(`SELECT \* FROM "product_entities" WHERE id = \$1 ORDER BY "product_entities"."id" LIMIT \$2`).
		WithArgs(string(productID), 1).
		WillReturnError(gorm.ErrRecordNotFound)

	product, err := repo.Find(productID)
	assert.Nil(t, product)
	assert.ErrorIs(t, err, domain.ErrNotFoundProduct)

	assert.NoError(t, mock.ExpectationsWereMet())
}
