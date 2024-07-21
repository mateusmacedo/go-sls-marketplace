package domain

import "fmt"

func CreateProductAdder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository, ok := dependencies["ProductFindRepository"].(ProductFindRepository)
	if !ok || findRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductFindRepository dependency")
	}

	saveRepository, ok := dependencies["ProductSaveRepository"].(ProductSaveRepository)
	if !ok || saveRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductSaveRepository dependency")
	}

	return NewProductAdder(findRepository, saveRepository), nil
}

func CreateProductDeleter(dependencies map[string]interface{}) (interface{}, error) {
	findRepository, ok := dependencies["ProductFindRepository"].(ProductFindRepository)
	if !ok || findRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductFindRepository dependency")
	}

	deleteRepository, ok := dependencies["ProductDeleteRepository"].(ProductDeleteRepository)
	if !ok || deleteRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductDeleteRepository dependency")
	}

	return NewProductDeleter(findRepository, deleteRepository), nil
}

func CreateProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository, ok := dependencies["ProductFindRepository"].(ProductFindRepository)
	if !ok || findRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductFindRepository dependency")
	}

	return NewProductFinder(findRepository), nil
}

func CreateAllProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findAllRepository, ok := dependencies["ProductFindAllRepository"].(ProductFindAllRepository)
	if !ok || findAllRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductFindAllRepository dependency")
	}

	return NewAllProductFinder(findAllRepository), nil
}

func CreateProductUpdater(dependencies map[string]interface{}) (interface{}, error) {
	findRepository, ok := dependencies["ProductFindRepository"].(ProductFindRepository)
	if !ok || findRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductFindRepository dependency")
	}

	saveRepository, ok := dependencies["ProductSaveRepository"].(ProductSaveRepository)
	if !ok || saveRepository == nil {
		return nil, fmt.Errorf("missing or nil ProductSaveRepository dependency")
	}

	return NewProductUpdater(findRepository, saveRepository), nil
}
