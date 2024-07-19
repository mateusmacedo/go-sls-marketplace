package domain

type ProductSaveRepository interface {
	Save(product *Product) error
}

type ProductFindRepository interface {
	Find(id ProductID) (*Product, error)
}

type ProductFindAllRepository interface {
	FindAll() ([]*Product, error)
}

type ProductDeleteRepository interface {
	Delete(id ProductID) error
}

type ProductRepository interface {
	ProductSaveRepository
	ProductFindRepository
	ProductFindAllRepository
	ProductDeleteRepository
}
