package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	dbadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/db/gorm/adapter"
	httpadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http/net/adapter"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
)

func InitializeServer() (*mux.Router, error) {
	dbConn, err := gorm.Open(sqlite.Open("catalog.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		return nil, err
	}

	err = dbConn.AutoMigrate(&dbadapter.GormProductEntity{})
	if err != nil {
		return nil, err
	}

	productRepo := dbadapter.NewGormProductRepository(dbConn)

	productService := domain.NewProductService(productRepo, productRepo, productRepo, productRepo)

	addProductUseCase := application.NewProductAddUseCase(productService)
	deleteProductUseCase := application.NewDeleteProductUseCase(productService)
	getAllProductsUseCase := application.NewGetAllProductsUseCase(productService)
	getProductUseCase := application.NewGetProductUseCase(productService)
	updateProductUseCase := application.NewUpdateProductUseCase(productService)

	addProductHandler := httpadapter.NewNetHTTPAddProductAdapter(
		httpadapter.WithService(addProductUseCase),
		httpadapter.WithMethodGuard(
			pkghttp.NewHttpMethodGuard([]string{http.MethodPost}),
		),
	)

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

	return r, nil
}

func main() {
	r, err := InitializeServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
