package adapter

import (
	"gorm.io/gorm"
)

func CreateProductSaveRepository(dependencies map[string]interface{}) (interface{}, error) {
	dbConn := dependencies["dbConn"].(*gorm.DB)

	return NewGormProductSaveRepository(dbConn), nil
}

func CreateProductFindRepository(dependencies map[string]interface{}) (interface{}, error) {
	dbConn := dependencies["dbConn"].(*gorm.DB)

	return NewGormProductFindRepository(dbConn), nil
}

func CreateProductFindAllRepository(dependencies map[string]interface{}) (interface{}, error) {
	dbConn := dependencies["dbConn"].(*gorm.DB)

	return NewGormProductFindAllRepository(dbConn), nil
}

func CreateProductDeleteRepository(dependencies map[string]interface{}) (interface{}, error) {
	dbConn := dependencies["dbConn"].(*gorm.DB)

	return NewGormProductDeleteRepository(dbConn), nil
}
