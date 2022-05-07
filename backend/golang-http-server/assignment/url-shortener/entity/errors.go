package entity

import "fmt"

var ErrURLNotFound = fmt.Errorf("error URL not found")
var ErrBadRequest = fmt.Errorf("error bad request")
var ErrCustomURLIsExists = fmt.Errorf("error custom url is exists")
