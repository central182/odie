//go:generate mockgen -source $GOFILE -package dictionary_service_mock -destination mock/$GOFILE

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
