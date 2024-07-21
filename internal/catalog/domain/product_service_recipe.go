package domain

func CreateProductAdder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["findRepository"].(ProductFindRepository)
	saveRepository := dependencies["saveRepository"].(ProductSaveRepository)

	return NewProductAdder(findRepository, saveRepository), nil
}

func CreateProductDeleter(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["findRepository"].(ProductFindRepository)
	deleteRepository := dependencies["deleteRepository"].(ProductDeleteRepository)

	return NewProductDeleter(findRepository, deleteRepository), nil
}

func CreateProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["findRepository"].(ProductFindRepository)

	return NewProductFinder(findRepository), nil
}

func CreateAllProductFinder(dependencies map[string]interface{}) (interface{}, error) {
	findAllRepository := dependencies["findAllRepository"].(ProductFindAllRepository)

	return NewAllProductFinder(findAllRepository), nil
}

func CreateProductUpdater(dependencies map[string]interface{}) (interface{}, error) {
	findRepository := dependencies["findRepository"].(ProductFindRepository)
	saveRepository := dependencies["saveRepository"].(ProductSaveRepository)

	return NewProductUpdater(findRepository, saveRepository), nil
}
