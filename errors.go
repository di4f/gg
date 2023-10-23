package gg

import (
	"errors"
)

var (
	ObjectExistErr = errors.New("the object already exists")
	ObjectNotExistErr = errors.New("the object does not exist")
	ObjectNotImplementedErr = errors.New("none of object methods are implemented")
)

