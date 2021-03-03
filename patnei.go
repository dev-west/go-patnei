package patnei

import (
	"errors"
)

type ExObj struct {
	S string
}

var eEmpty = &ExObj{ }

func NewObj() *ExObj {
	return &ExObj{ }
}

func (x *ExObj) Add(c string) (*ExObj, error) {
	eStr := "Add will lose data: " + x.S
	if len(x.S) != 0 { return eEmpty, errors.New(eStr) }

	x.S = c

	return x, nil
}

func (x *ExObj) Append(c string) (*ExObj, error) {
	if len(x.S) == 0 { return eEmpty, errors.New("Append to empty ExObj") }
	x.S = x.S + c

	return x, nil
}

