package domain

import "errors"

type ProductAdder interface {
	AddProduct(id ProductID, name, description string, price float64) (*Product, error)
}

type AllProductFinder interface {
	GetAllProducts() ([]*Product, error)
}

type ProductFinder interface {
	GetProduct(id ProductID) (*Product, error)
}

type ProductService struct {
	saveRepository    ProductSaveRepository
	findRepository    ProductFindRepository
	findAllRepository ProductFindAllRepository
	deleteRepository  ProductDeleteRepository
}

func NewProductService(saveRepository ProductSaveRepository, findRepository ProductFindRepository, findAllRepository ProductFindAllRepository, deleteRepository ProductDeleteRepository) *ProductService {
	return &ProductService{
		saveRepository:    saveRepository,
		findRepository:    findRepository,
		findAllRepository: findAllRepository,
		deleteRepository:  deleteRepository,
	}
}

func (s *ProductService) AddProduct(id ProductID, name, description string, price float64) (*Product, error) {
	productExists, err := s.findRepository.Find(id)
	if err != nil {
		return nil, err
	}
	if productExists != nil {
		return nil, errors.New("product already exists")
	}

	product, err := NewProduct(id, name, description, price)
	if err != nil {
		return nil, err
	}

	err = s.saveRepository.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetAllProducts() ([]*Product, error) {
	records, err := s.findAllRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *ProductService) GetProduct(id ProductID) (*Product, error) {
	if id == "" {
		return nil, errors.New("invalid product ID")
	}
	return s.findRepository.Find(id)
}

func (s *ProductService) UpdateProduct(id ProductID, name, description string, price float64) (*Product, error) {
	if id == "" {
		return nil, errors.New("invalid product ID")
	}

	product, err := s.findRepository.Find(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if err := product.ChangeName(name); err != nil {
		return nil, err
	}

	if err := product.ChangeDescription(description); err != nil {
		return nil, err
	}

	if err := product.ChangePrice(price); err != nil {
		return nil, err
	}

	err = s.saveRepository.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id ProductID) error {
	if id == "" {
		return errors.New("invalid product ID")
	}

	product, err := s.findRepository.Find(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	err = s.deleteRepository.Delete(product.ID)
	if err != nil {
		return err
	}

	return nil
}
