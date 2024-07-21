package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/application"
	"github.com/mateusmacedo/go-sls-marketplace/internal/catalog/domain"
	dbadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/db/gorm/adapter"
	httpadapter "github.com/mateusmacedo/go-sls-marketplace/internal/catalog/infrastructure/http/net/adapter"
	pkghttp "github.com/mateusmacedo/go-sls-marketplace/pkg/infrastructure/http"
)

func InitializeServer() (*mux.Router, error) {
	dbConn, err := gorm.Open(sqlite.Open("catalog.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = dbConn.AutoMigrate(&dbadapter.GormProductEntity{})
	if err != nil {
		return nil, err
	}

	productFindRepo := dbadapter.NewGormProductFindRepository(dbConn)
	productSaveRepo := dbadapter.NewGormProductSaveRepository(dbConn)
	adderProductService := domain.NewProductAdder(productFindRepo, productSaveRepo)
	addProductUseCase := application.NewAddProductUseCase(adderProductService)

	productDeleteRepo := dbadapter.NewGormProductDeleteRepository(dbConn)
	deleterProductService := domain.NewProductDeleter(productFindRepo, productDeleteRepo)
	deleteProductUseCase := application.NewDeleteProductUseCase(deleterProductService)

	productFindAllRepo := dbadapter.NewGormProductFindAllRepository(dbConn)
	fiderAllProductService := domain.NewAllProductFinder(productFindAllRepo)
	getAllProductsUseCase := application.NewGetAllProductsUseCase(fiderAllProductService)

	finderProductService := domain.NewProductFinder(productFindRepo)
	getProductUseCase := application.NewGetProductUseCase(finderProductService)

	updaterProductService := domain.NewProductUpdater(productFindRepo, productSaveRepo)
	updateProductUseCase := application.NewUpdateProductUseCase(updaterProductService)

	postHttpMethodGuard := pkghttp.NewHttpMethodGuard([]string{http.MethodPost})

	addProductHandler := httpadapter.NewNetHTTPAddProductAdapter(
		httpadapter.WithService(addProductUseCase),
		httpadapter.WithMethodGuard(postHttpMethodGuard),
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
