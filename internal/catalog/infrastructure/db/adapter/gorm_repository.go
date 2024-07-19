package adapter

import (
	"errors"

	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{
		db: db,
	}
}

func (repo *GormProductRepository) Save(product *domain.Product) error {
	entity, err := NewProductEntityFromDomain(product)
	if err != nil {
		return err
	}
	return repo.db.Create(entity).Error
}

func (repo *GormProductRepository) Find(id domain.ProductID) (*domain.Product, error) {
	var entity ProductEntity
	err := repo.db.First(&entity, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFoundProduct
		}
		return nil, err
	}
	return entity.ToDomain()
}

func (repo *GormProductRepository) FindAll() ([]*domain.Product, error) {
	var entities []ProductEntity
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

func (repo *GormProductRepository) Delete(id domain.ProductID) error {
	return repo.db.Delete(&ProductEntity{}, "id = ?", id).Error
}