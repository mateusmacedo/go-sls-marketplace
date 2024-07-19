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
go test ./... -timeout 10s -tags wireinject --cover --race
```

To generate mocks for the interfaces, you can use the following commands:

```bash
mockgen -source=path/to/source/interface.go -destination=path/to/mocks/interface_mock.go -package=mocks
```
