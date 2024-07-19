package domain

import "errors"

var ErrInvalidProductID = errors.New("invalid product ID")
var ErrInvalidProductName = errors.New("invalid product name")
var ErrInvalidProductDescription = errors.New("invalid product description")
var ErrInvalidProductPrice = errors.New("invalid product price")
var ErrAlreadyExistsProduct = errors.New("product already exists")
var ErrNotFoundProduct = errors.New("product not found")
var ErrRepositoryProduct = errors.New("error in repository")
