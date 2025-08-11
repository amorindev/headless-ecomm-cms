package errors

import "errors"

var ErrProductNotFound = errors.New("product-not-found")
var ErrProductItemNotFound = errors.New("product-item-not-found")