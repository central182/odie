package application

import (
	"github.com/sanity-io/litter"

	"github.com/central182/odie/internal/domain/dictionary/dictionary_service"
)

type GetEntriesByHeadwordError interface {
	error
	DictionaryServiceFailure() dictionary_service.GetEntriesError
}

type getEntriesByHeadwordError struct {
	dictionaryServiceFailure dictionary_service.GetEntriesError
}

func (e *getEntriesByHeadwordError) Error() string {
	return litter.Options{}.Sdump(*e)
}

func (e *getEntriesByHeadwordError) DictionaryServiceFailure() dictionary_service.GetEntriesError {
	return e.dictionaryServiceFailure
}

func (e *getEntriesByHeadwordError) touched() bool {
	return e.dictionaryServiceFailure != nil
}

func (e *getEntriesByHeadwordError) setDictionaryServiceFailure(err dictionary_service.GetEntriesError) {
	e.dictionaryServiceFailure = err
}
