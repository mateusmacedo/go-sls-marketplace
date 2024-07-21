package domain

type ProductAdder interface {
	AddProduct(id ProductID, name, description string, price float64) (*Product, error)
}

type productAdder struct {
	findRepository ProductFindRepository
	saveRepository ProductSaveRepository
}

func NewProductAdder(findRepository ProductFindRepository, repository ProductSaveRepository) ProductAdder {
	return &productAdder{
		findRepository: findRepository,
		saveRepository: repository,
	}
}

func (s *productAdder) AddProduct(id ProductID, name, description string, price float64) (*Product, error) {
	productExists, err := s.findRepository.Find(id)
	if err != nil {
		if err != ErrNotFoundProduct {
			return nil, err
		}
	}

	if productExists != nil {
		return nil, ErrAlreadyExistsProduct
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

type AllProductFinder interface {
	GetAllProducts() ([]*Product, error)
}

type allProductFinder struct {
	findAllRepository ProductFindAllRepository
}

func NewAllProductFinder(findAllRepository ProductFindAllRepository) AllProductFinder {
	return &allProductFinder{
		findAllRepository: findAllRepository,
	}
}

func (s *allProductFinder) GetAllProducts() ([]*Product, error) {
	records, err := s.findAllRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

type ProductFinder interface {
	GetProduct(id ProductID) (*Product, error)
}

type productFinder struct {
	findRepository ProductFindRepository
}

func NewProductFinder(findRepository ProductFindRepository) ProductFinder {
	return &productFinder{
		findRepository: findRepository,
	}
}

func (s *productFinder) GetProduct(id ProductID) (*Product, error) {
	if id == "" {
		return nil, ErrInvalidProductID
	}
	return s.findRepository.Find(id)
}

type ProductUpdater interface {
	UpdateProduct(id ProductID, name, description string, price float64) (*Product, error)
}

type productUpdater struct {
	findRepository ProductFindRepository
	saveRepository ProductSaveRepository
}

func NewProductUpdater(findRepository ProductFindRepository, saveRepository ProductSaveRepository) ProductUpdater {
	return &productUpdater{
		findRepository: findRepository,
		saveRepository: saveRepository,
	}
}

func (s *productUpdater) UpdateProduct(id ProductID, name, description string, price float64) (*Product, error) {
	if id == "" {
		return nil, ErrInvalidProductID
	}

	product, err := s.findRepository.Find(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrNotFoundProduct
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

type ProductDeleter interface {
	DeleteProduct(id ProductID) error
}

type productDeleter struct {
	findRepository   ProductFindRepository
	deleteRepository ProductDeleteRepository
}

func NewProductDeleter(findRepository ProductFindRepository, deleteRepository ProductDeleteRepository) ProductDeleter {
	return &productDeleter{
		findRepository:   findRepository,
		deleteRepository: deleteRepository,
	}
}

func (s *productDeleter) DeleteProduct(id ProductID) error {
	if id == "" {
		return ErrInvalidProductID
	}

	product, err := s.findRepository.Find(id)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrNotFoundProduct
	}

	err = s.deleteRepository.Delete(product.ID)
	if err != nil {
		return err
	}

	return nil
}
