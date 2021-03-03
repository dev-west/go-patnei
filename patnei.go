package patnei

import (
	"errors"
)

type ExObj struct {
	s string
}

var eEmpty = &ExObj{ }

func NewObj() *ExObj {
	return &ExObj{ }
}

func (x *ExObj) Add(c string) (*ExObj, error) {
	if len(x.s) != 0 { return eEmpty, errors.New("Add will lose data: %s", x.s) }

	x.s = c

	return x
}

func (x *ExObj) Append(c string) (*ExObj, error) {
	if len(x.s) == 0 { return eEmpty, errors.New("Append to empty ExObj") }
	x.s = x.s + c

	return x
}

