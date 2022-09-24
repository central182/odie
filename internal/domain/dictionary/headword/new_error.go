package headword

import (
	"github.com/sanity-io/litter"
)

type NewError interface {
	error
	HasEmptySpelling() bool
}

type newError struct {
	hasEmptySpelling bool
}

func (e *newError) Error() string {
	if !e.touched() {
		return ""
	}
	return litter.Options{}.Sdump(*e)
}

func (e *newError) HasEmptySpelling() bool {
	return e.hasEmptySpelling
}

func (e *newError) touched() bool {
	return *e != newError{}
}

func (e *newError) setHasEmptySpelling() {
	e.hasEmptySpelling = true
}
