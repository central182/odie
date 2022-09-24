package dictionary_service_odapi

import (
	"errors"

	"github.com/central182/odie/internal/adapter/outbound/common/odapi"
	"github.com/central182/odie/internal/domain/dictionary/dictionary_service"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

type dictionaryService struct {
	client odapi.Client
}

func New(client odapi.Client) dictionary_service.DictionaryService {
	if client == nil {
		panic("a non-nil odapi.Client is expected")
	}

	return dictionaryService{client: client}
}

func (d dictionaryService) GetEntries(h headword.Headword) ([]entry.Entry, dictionary_service.GetEntriesError) {
	err := &getEntriesError{}

	if h == nil {
		err.setHasNilHeadword()
		return nil, err
	}

	resp, gerr := d.client.GetEntries(
		odapi.GetEntriesRequest{
			SourceLang: "en-gb",
			WordId:     h.Spelling(),
		},
	)
	if gerr != nil {
		if errors.Is(gerr, odapi.ErrEntryNotFound) {
			err.setNoEntryWasFound()
			return nil, err
		}

		err.setInfrastructuralFailure(gerr)
		return nil, err
	}

	result := unmarshalGetEntriesResponse(h.Spelling(), resp, err)
	if err.touched() {
		return result, err
	}

	return result, nil
}
