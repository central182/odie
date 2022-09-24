package dictionary_service

import (
	"github.com/central182/odie/internal/domain/dictionary/entry"
)

type GetEntriesError interface {
	error
	HasNilHeadword() bool
	NoEntryWasFound() bool
	UnmarshallingFailures() []entry.NewError
	InfrastructuralFailure() error
}
