package application

import (
	"github.com/central182/odie/internal/domain/dictionary/dictionary_service"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

// Application is a facade over everything else in the domain,
// providing entry points for inbound adapters.
type Application interface {
	GetEntriesByHeadword(headword.Headword) ([]entry.Entry, GetEntriesByHeadwordError)
}

func New(ds dictionary_service.DictionaryService) Application {
	if ds == nil {
		panic("a non-nil dictionary_service.DictionaryService is expected")
	}

	return _application{dictionaryService: ds}
}

type _application struct {
	dictionaryService dictionary_service.DictionaryService
}

func (a _application) GetEntriesByHeadword(h headword.Headword) ([]entry.Entry, GetEntriesByHeadwordError) {
	err := &getEntriesByHeadwordError{}

	es, gerr := a.dictionaryService.GetEntries(h)
	if gerr != nil {
		err.setDictionaryServiceFailure(gerr)
	}

	if err.touched() {
		return es, err
	}

	return es, nil
}
