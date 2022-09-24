package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/central182/odie/internal/domain/application"
	dictionary_service_mock "github.com/central182/odie/internal/domain/dictionary/dictionary_service/mock"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

func TestApplication(t *testing.T) {
	t.Run("The Application", func(t *testing.T) {
		t.Run("can't be constructed from a nil DictionaryService", func(t *testing.T) {
			assert.Panics(t, func() {
				application.New(nil)
			})
		})
	})
}

func TestGetEntriesByHeadword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("The GetEntriesByHeadword method", func(t *testing.T) {
		t.Run("merely forwards the request to the underlying DictionaryService,", func(t *testing.T) {
			ds := dictionary_service_mock.NewMockDictionaryService(ctrl)
			h, es := fixture()
			ds.EXPECT().GetEntries(h).Return(es, nil)

			app := application.New(ds)
			actualEs, err := app.GetEntriesByHeadword(h)
			assert.Nil(t, err)
			assert.Equal(t, es, actualEs)
		})

		t.Run("and wraps errors if any.", func(t *testing.T) {
			ds := dictionary_service_mock.NewMockDictionaryService(ctrl)
			gerr := dictionary_service_mock.NewMockGetEntriesError(ctrl)
			h, es := fixture()
			ds.EXPECT().GetEntries(h).Return(es, gerr)

			app := application.New(ds)
			actualEs, err := app.GetEntriesByHeadword(h)
			assert.NotNil(t, err)
			assert.Equal(t, gerr, err.DictionaryServiceFailure())
			assert.Equal(t, es, actualEs)
		})
	})
}

func fixture() (headword.Headword, []entry.Entry) {
	h, herr := headword.New("simplified")
	if herr != nil {
		panic(herr)
	}

	e, eerr := entry.New(entry.NewInput{
		Headword:        "simplified",
		LexicalCategory: "adjective",
		Senses: []entry.NewSenseInput{
			{
				Description: "made simpler",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	es := []entry.Entry{e}

	return h, es
}
