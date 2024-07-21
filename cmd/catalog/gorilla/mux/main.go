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
	pkgapplication "github.com/mateusmacedo/go-sls-marketplace/pkg/application"
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

	serviceLocator := pkgapplication.NewSimpleServiceLocator()
	factory := pkgapplication.NewFactory(serviceLocator)

	serviceLocator.Register("dbConn", dbConn)

	// Repositories
	factory.RegisterRecipe("ProductSaveRepository", pkgapplication.Recipe{
		Dependencies: []string{"dbConn"},
		Factory:      dbadapter.CreateProductSaveRepository,
	})
	factory.RegisterRecipe("ProductFindRepository", pkgapplication.Recipe{
		Dependencies: []string{"dbConn"},
		Factory:      dbadapter.CreateProductFindRepository,
	})
	factory.RegisterRecipe("ProductFindAllRepository", pkgapplication.Recipe{
		Dependencies: []string{"dbConn"},
		Factory:      dbadapter.CreateProductFindAllRepository,
	})
	factory.RegisterRecipe("ProductDeleteRepository", pkgapplication.Recipe{
		Dependencies: []string{"dbConn"},
		Factory:      dbadapter.CreateProductDeleteRepository,
	})
	// Domain Services
	factory.RegisterRecipe("ProductAdder", pkgapplication.Recipe{
		Dependencies: []string{"ProductFindRepository", "ProductSaveRepository"},
		Factory:      domain.CreateProductAdder,
	})
	factory.RegisterRecipe("ProductDeleter", pkgapplication.Recipe{
		Dependencies: []string{"ProductFindRepository", "ProductDeleteRepository"},
		Factory:      domain.CreateProductDeleter,
	})
	factory.RegisterRecipe("ProductFinder", pkgapplication.Recipe{
		Dependencies: []string{"ProductFindRepository"},
		Factory:      domain.CreateProductFinder,
	})
	factory.RegisterRecipe("AllProductFinder", pkgapplication.Recipe{
		Dependencies: []string{"ProductFindAllRepository"},
		Factory:      domain.CreateAllProductFinder,
	})
	// Application Use Cases
	factory.RegisterRecipe("AddProductUseCase", pkgapplication.Recipe{
		Dependencies: []string{"ProductAdder"},
		Factory:      application.CreateAddProductUseCase,
	})
	factory.RegisterRecipe("DeleteProductUseCase", pkgapplication.Recipe{
		Dependencies: []string{"ProductDeleter"},
		Factory:      application.CreateDeleteProductUseCase,
	})
	factory.RegisterRecipe("GetAllProductsUseCase", pkgapplication.Recipe{
		Dependencies: []string{"AllProductFinder"},
		Factory:      application.CreateGetAllProductsUseCase,
	})
	factory.RegisterRecipe("GetProductUseCase", pkgapplication.Recipe{
		Dependencies: []string{"ProductFinder"},
		Factory:      application.CreateGetProductUseCase,
	})
	factory.RegisterRecipe("UpdateProductUseCase", pkgapplication.Recipe{
		Dependencies: []string{"ProductUpdater"},
		Factory:      application.CreateUpdateProductUseCase,
	})

	addProductUseCase, err := factory.Create("AddProductUseCase")
	if err != nil {
		panic(err)
	}

	deleteProductUseCase, err := factory.Create("DeleteProductUseCase")
	if err != nil {
		panic(err)
	}

	getAllProductsUseCase, err := factory.Create("GetAllProductsUseCase")
	if err != nil {
		panic(err)
	}

	getProductUseCase, err := factory.Create("GetProductUseCase")
	if err != nil {
		panic(err)
	}

	updateProductUseCase, err := factory.Create("UpdateProductUseCase")
	if err != nil {
		panic(err)
	}

	postHttpMethodGuard := pkghttp.NewHttpMethodGuard([]string{http.MethodPost})

	addProductHandler := httpadapter.NewNetHTTPAddProductAdapter(
		httpadapter.WithService(addProductUseCase.(application.AddProductUseCase)),
		httpadapter.WithMethodGuard(postHttpMethodGuard),
	)

	deleteProductHandler := httpadapter.NewNetHTTPDeleteProductAdapter(deleteProductUseCase.(application.DeleteProductUseCase))
	getAllProductsHandler := httpadapter.NewNetHTTPGetAllProductsAdapter(getAllProductsUseCase.(application.GetAllProductsUseCase))
	getProductHandler := httpadapter.NewNetHTTPGetProductAdapter(getProductUseCase.(application.GetProductUseCase))
	updateProductHandler := httpadapter.NewNetHTTPUpdateProductAdapter(updateProductUseCase.(application.UpdateProductUseCase))

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
