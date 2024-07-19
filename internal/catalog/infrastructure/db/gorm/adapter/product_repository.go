package adapter

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) domain.ProductRepository {
	return &gormProductRepository{
		db: db,
	}
}

func (repo *gormProductRepository) Save(product *domain.Product) error {
	entity, err := NewProductEntityFromDomain(product)
	if err != nil {
		return err
	}
	result := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "description", "price", "updated_at"}),
	}).Create(entity)

	return result.Error
}

func (repo *gormProductRepository) Find(id domain.ProductID) (*domain.Product, error) {
	var entity GormProductEntity
	err := repo.db.First(&entity, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFoundProduct
		}
		return nil, err
	}
	return entity.ToDomain()
}

func (repo *gormProductRepository) FindAll() ([]*domain.Product, error) {
	var entities []GormProductEntity
	err := repo.db.Find(&entities).Error
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

func (repo *gormProductRepository) Delete(id domain.ProductID) error {
	return repo.db.Delete(&GormProductEntity{}, "id = ?", id).Error
}
