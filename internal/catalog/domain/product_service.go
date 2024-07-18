package domain

import "errors"

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
	return s.findAllRepository.FindAll()
}

func (s *ProductService) GetProduct(id ProductID) (*Product, error) {
	return s.findRepository.Find(id)
}

func (s *ProductService) UpdateProduct(id ProductID, name, description string, price float64) (*Product, error) {
	product, err := s.findRepository.Find(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	product.ChangeName(name)
	product.ChangeDescription(description)
	product.ChangePrice(price)

	err = s.saveRepository.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id ProductID) error {
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
