//go:generate mockgen -source $GOFILE -package dictionary_service_mock -destination mock/$GOFILE

package dictionary_service

import (
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

// DictionaryService is a gateway to an external dictionary service.
type DictionaryService interface {
	// There may exist one or more Entries having the given Headword.
	GetEntries(headword.Headword) ([]entry.Entry, GetEntriesError)
}
