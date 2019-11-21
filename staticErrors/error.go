package staticErrors

import (
	"errors"
)

var EmptyRequest = errors.New("url can not be empty")
