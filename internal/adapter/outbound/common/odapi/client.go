//go:generate mockgen -source $GOFILE -package odapi_mock -destination mock/$GOFILE

package odapi

import "errors"

// Client is a consumer of the Oxford Dictionaries API.
// The vocabularies used here are kept as close as possible to those that are defined
// in the Swagger documentation provided by Oxford Dictionaries API.
//
// Usage of OpenAPI Generator is deliberately avoided because the generated artifacts
// are not easily swappable in tests.
type Client interface {
	GetEntries(GetEntriesRequest) (GetEntriesResponse, error)
}

var (
	ErrEntryNotFound = errors.New("no entry was found")
)

type GetEntriesRequest struct {
	SourceLang string
	WordId     string
}

type GetEntriesResponse struct {
	Results []HeadwordEntry
}

type HeadwordEntry struct {
	LexicalEntries []LexicalEntry
}

type LexicalEntry struct {
	LexicalCategory LexicalCategory
	Entries         []Entry
}

type LexicalCategory struct {
	Id string
}

type Entry struct {
	Pronunciations []Pronunciation
	Senses         []Sense
}

type Pronunciation struct {
	AudioFile        string
	PhoneticSpelling string
}

type Sense struct {
	Definitions           []string
	CrossReferenceMarkers []string
	Examples              []Example
	Subsenses             []Sense
}

type Example struct {
	Text string
}

type ErrorDescription struct {
	Error string
}
