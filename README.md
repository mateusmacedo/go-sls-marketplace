# Golang Serverless Marketplace System

This is a serverless marketplace system that is built using Golang and AWS Lambda. The system is designed to be a simple marketplace system.

## Features

- Create a new product in the marketplace
- Read an product from the marketplace
- Update an product in the marketplace
- Delete an product from the marketplace
- List all products in the marketplace
- Search for products in the marketplace - to be implemented
- Filter products in the marketplace - to be implemented
- Sort products in the marketplace - to be implemented
- Paginate products in the marketplace - to be implemented
- Rate an product in the marketplace - to be implemented
- Comment on an product in the marketplace - to be implemented
- Report an product in the marketplace - to be implemented
- Flag an product in the marketplace - to be implemented
- Add an product to the user's wishlist - to be implemented
- Remove an product from the user's wishlist - to be implemented
- List all products in the user's wishlist - to be implemented
- Search for products in the user's wishlist - to be implemented
- Filter products in the user's wishlist - to be implemented
- Sort products in the user's wishlist - to be implemented
- Paginate products in the user's wishlist - to be implemented

## Testing

To test the system, you can use the following commands:

```bash
# Run the tests
go test ./internal/catalog/application -timeout 30s -tags wireinject --cover --race -count=1
go test ./internal/catalog/domain -timeout 30s -tags wireinject --cover --race -count=1
go test ./internal/catalog/infrastructure/db/dynamodb/adapter -timeout 30s -tags wireinject --cover --race -count=1
go test ./internal/catalog/infrastructure/db/gorm/adapter -timeout 30s -tags wireinject --cover --race -count=1
go test ./internal/catalog/infrastructure/http/aws/adapter -timeout 30s -tags wireinject --cover --race -count=1
go test ./internal/catalog/infrastructure/http/net/adapter -timeout 30s -tags wireinject --cover --race -count=1
go test ./pkg/application -timeout 30s -tags wireinject --cover --race -count=1
go test ./pkg/infrastructure/http -timeout 30s -tags wireinject --cover --race -count=1
go test ./pkg/infrastructure/log -timeout 30s -tags wireinject --cover --race -count=1
```

To generate mocks for the interfaces, you can use the following commands:

```bash
mockgen -destination=test/domain/mocks/product_save_repository.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductSaveRepository
mockgen -destination=test/domain/mocks/product_find_repository.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductFindRepository
mockgen -destination=test/domain/mocks/product_find_all_repository.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductFindAllRepository
mockgen -destination=test/domain/mocks/product_delete_repository.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductDeleteRepository
mockgen -destination=test/domain/mocks/product_adder.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductAdder
mockgen -destination=test/domain/mocks/product_allproduct_finder.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain AllProductFinder
mockgen -destination=test/domain/mocks/product_product_finder.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductFinder
mockgen -destination=test/domain/mocks/product_updater.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductUpdater
mockgen -destination=test/domain/mocks/product_deleter.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain ProductDeleter

mockgen -destination=test/application/mocks/add_product_use_case.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application AddProductUseCase
mockgen -destination=test/application/mocks/delete_product_use_case.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application DeleteProductUseCase
mockgen -destination=test/application/mocks/get_all_products_use_case.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application GetAllProductsUseCase
mockgen -destination=test/application/mocks/get_product_use_case.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application GetProductUseCase
mockgen -destination=test/application/mocks/update_product_use_case.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application UpdateProductUseCase

mockgen -destination=test/infrastructure/mocks/dynamo_dbapi.go -package=mocks github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/db/dynamodb/adapter DynamoDBAPI
```
