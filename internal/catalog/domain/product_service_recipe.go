package domain

func CreateProductAdder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["ProductFindRepository"].(ProductFindRepository)
	saveRepository := dependencies["ProductSaveRepository"].(ProductSaveRepository)

	return NewProductAdder(findRepository, saveRepository), nil
}

func CreateProductDeleter(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["ProductFindRepository"].(ProductFindRepository)
	deleteRepository := dependencies["ProductDeleteRepository"].(ProductDeleteRepository)

	return NewProductDeleter(findRepository, deleteRepository), nil
}

func CreateProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["ProductFindRepository"].(ProductFindRepository)

	return NewProductFinder(findRepository), nil
}

func CreateAllProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findAllRepository := dependencies["ProductFindAllRepository"].(ProductFindAllRepository)

	return NewAllProductFinder(findAllRepository), nil
}

func CreateProductUpdater(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["ProductFindRepository"].(ProductFindRepository)
	saveRepository := dependencies["ProductSaveRepository"].(ProductSaveRepository)

	return NewProductUpdater(findRepository, saveRepository), nil
}
