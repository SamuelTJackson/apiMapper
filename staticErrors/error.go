package staticErrors

import (
	"errors"
)

var EmptyRequest = errors.New("url can not be empty")

var IDDoesNotExists = errors.New("the requested id does not exists")
