package constants

import "errors"

var ErrDuplicateData = errors.New("data already present in database")
var ErrDataNotFound = errors.New("data is not present in database")
