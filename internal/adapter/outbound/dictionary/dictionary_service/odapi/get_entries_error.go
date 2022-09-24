package dictionary_service_odapi

import (
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/sanity-io/litter"
)

type getEntriesError struct {
	hasNilHeadword         bool
	noEntryWasFound        bool
	unmarshallingFailures  []entry.NewError
	infrastructuralFailure error
}

func (e *getEntriesError) Error() string {
	return litter.Options{}.Sdump(*e)
}

func (e *getEntriesError) HasNilHeadword() bool {
	return e.hasNilHeadword
}

func (e *getEntriesError) NoEntryWasFound() bool {
	return e.noEntryWasFound
}

func (e *getEntriesError) UnmarshallingFailures() []entry.NewError {
	return e.unmarshallingFailures
}

func (e *getEntriesError) InfrastructuralFailure() error {
	return e.infrastructuralFailure
}

func (e *getEntriesError) touched() bool {
	if e.hasNilHeadword {
		return true
	}

	if e.noEntryWasFound {
		return true
	}

	for _, uf := range e.unmarshallingFailures {
		if uf != nil {
			return true
		}
	}

	if e.infrastructuralFailure != nil {
		return true
	}

	return false
}

func (e *getEntriesError) setHasNilHeadword() {
	e.hasNilHeadword = true
}

func (e *getEntriesError) setNoEntryWasFound() {
	e.noEntryWasFound = true
}

func (e *getEntriesError) appendUnmarshallingFailure(err entry.NewError) {
	e.unmarshallingFailures = append(e.unmarshallingFailures, err)
}

func (e *getEntriesError) setInfrastructuralFailure(err error) {
	e.infrastructuralFailure = err
}
