package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	dbadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/db/adapter"
	httpadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http/adapter"
)

func main() {
	dbConn, err := gorm.Open(sqlite.Open("catalog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.AutoMigrate(&dbadapter.ProductEntity{})
	if err != nil {
		log.Fatal(err)
	}

	productRepo := dbadapter.NewGormProductRepository(dbConn)

	productService := domain.NewProductService(productRepo, productRepo, productRepo, productRepo)

	addProductUseCase := application.NewProductAddUseCase(productService)
	deleteProductUseCase := application.NewDeleteProductUseCase(productService)
	getAllProductsUseCase := application.NewGetAllProductsUseCase(productService)
	getProductUseCase := application.NewGetProductsUseCase(productService)
	updateProductUseCase := application.NewUpdateProductUseCase(productService)

	addProductHandler := httpadapter.NewNetHTTPAddProductAdapter(addProductUseCase)
	deleteProductHandler := httpadapter.NewNetHTTPDeleteProductAdapter(deleteProductUseCase)
	getAllProductsHandler := httpadapter.NewNetHTTPGetAllProductsAdapter(getAllProductsUseCase)
	getProductHandler := httpadapter.NewNetHTTPGetProductAdapter(getProductUseCase)
	updateProductHandler := httpadapter.NewNetHTTPUpdateProductAdapter(updateProductUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/products", addProductHandler.Handle).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", deleteProductHandler.Handle).Methods(http.MethodDelete)
	r.HandleFunc("/products", getAllProductsHandler.Handle).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", getProductHandler.Handle).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", updateProductHandler.Handle).Methods(http.MethodPut)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
