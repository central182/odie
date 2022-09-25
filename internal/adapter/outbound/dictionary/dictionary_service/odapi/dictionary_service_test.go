package dictionary_service_odapi_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/central182/odie/internal/adapter/outbound/common/odapi"
	odapi_mock "github.com/central182/odie/internal/adapter/outbound/common/odapi/mock"
	dictionary_service_odapi "github.com/central182/odie/internal/adapter/outbound/dictionary/dictionary_service/odapi"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

func TestDictionaryService(t *testing.T) {
	t.Run("The DictionaryService", func(t *testing.T) {
		t.Run("can't be constructed from a nil Client.", func(t *testing.T) {
			assert.Panics(t, func() {
				dictionary_service_odapi.New(nil)
			})
		})
	})
}

func TestGetEntries(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("The GetEntries can't be called", func(t *testing.T) {
		t.Run("if a nil Headword is provided.", func(t *testing.T) {
			dsvc := dictionary_service_odapi.New(odapi_mock.NewMockClient(ctrl))
			es, err := dsvc.GetEntries(nil)
			assert.Zero(t, es)
			assert.True(t, err.HasNilHeadword())
			assert.NotEmpty(t, err.Error())
		})
	})

	t.Run("The GetEntries method", func(t *testing.T) {
		t.Run("returns Entries if the underlying Client responded with data in expected format.", func(t *testing.T) {
			fixture := documentation()

			c := odapi_mock.NewMockClient(ctrl)
			c.EXPECT().GetEntries(fixture.req).Return(fixture.resp, nil)

			dsvc := dictionary_service_odapi.New(c)
			es, err := dsvc.GetEntries(fixture.h)
			assert.Nil(t, err)
			assert.Equal(t, fixture.es, es)
		})

		t.Run("complains that no Entry could be found if the underlying Client said so.", func(t *testing.T) {
			c := odapi_mock.NewMockClient(ctrl)
			c.EXPECT().GetEntries(gomock.Any()).Return(odapi.GetEntriesResponse{}, odapi.ErrEntryNotFound)

			dsvc := dictionary_service_odapi.New(c)
			es, err := dsvc.GetEntries(documentation().h)
			assert.Zero(t, es)
			assert.True(t, err.NoEntryWasFound())
			assert.NotEmpty(t, err.Error())
		})

		t.Run("reports failure when the underlying Client failed unexpectedly.", func(t *testing.T) {
			c := odapi_mock.NewMockClient(ctrl)
			c.EXPECT().GetEntries(gomock.Any()).Return(odapi.GetEntriesResponse{}, errors.New("something is very wrong"))

			dsvc := dictionary_service_odapi.New(c)
			es, err := dsvc.GetEntries(documentation().h)
			assert.Zero(t, es)

			ierr := err.InfrastructuralFailure()
			assert.ErrorContains(t, ierr, "something is very wrong")
			assert.NotEmpty(t, err.Error())
		})

		t.Run("tries to unmarshal as many Entries as possible: only the data that is causing problems will be ignored.", func(t *testing.T) {
			fixture := bestEffort()

			c := odapi_mock.NewMockClient(ctrl)
			c.EXPECT().GetEntries(fixture.req).Return(fixture.resp, nil)

			dsvc := dictionary_service_odapi.New(c)
			es, err := dsvc.GetEntries(fixture.h)
			assert.Equal(t, fixture.es, es)
			assert.Len(t, err.UnmarshallingFailures(), 2)
			assert.NotEmpty(t, err.Error())
		})
	})
}

type getEntriesFixture struct {
	h    headword.Headword
	req  odapi.GetEntriesRequest
	resp odapi.GetEntriesResponse
	es   []entry.Entry
}

// Fixture for the word "documentation".
// Modelled after the real response given by Oxford Dictionaries API.
func documentation() getEntriesFixture {
	h, herr := headword.New("documentation")
	if herr != nil {
		panic(herr)
	}

	req := odapi.GetEntriesRequest{
		SourceLang: "en-gb",
		WordId:     "documentation",
	}

	resp := odapi.GetEntriesResponse{
		Results: []odapi.HeadwordEntry{
			{
				LexicalEntries: []odapi.LexicalEntry{
					{
						Entries: []odapi.Entry{
							{
								Pronunciations: []odapi.Pronunciation{
									{
										AudioFile:        "https://example.com/documentation.mp3",
										PhoneticSpelling: "ˌdɒkjʊm(ɛ)nˈteɪʃn",
									},
								},
								Senses: []odapi.Sense{
									{
										Definitions: []string{"some official material"},
										Examples:    []odapi.Example{{Text: "Complete the relevant documentation!"}},
										Subsenses: []odapi.Sense{
											{
												Definitions: []string{"something accompanying a computer program"},
												Examples:    []odapi.Example{{Text: "A user documentation."}},
											},
										},
									},
									{
										Definitions: []string{"some process of classifying things"},
										Examples:    []odapi.Example{{Text: "Arrange the documentation of photographs!"}},
									},
								},
							},
						},
						LexicalCategory: odapi.LexicalCategory{
							Id: "noun",
						},
					},
				},
			},
		},
	}

	e, eerr := entry.New(entry.NewInput{
		Headword:        "documentation",
		LexicalCategory: "noun",
		Senses: []entry.NewSenseInput{
			{
				Description: "some official material",
				Examples:    []string{"Complete the relevant documentation!"},
				Subsenses: []entry.NewSubsenseInput{
					{
						Description: "something accompanying a computer program",
						Examples:    []string{"A user documentation."},
					},
				},
			},
			{
				Description: "some process of classifying things",
				Examples:    []string{"Arrange the documentation of photographs!"},
			},
		},
		Pronunciations: []entry.NewPronunciationInput{
			{
				PhoneticSpelling: "ˌdɒkjʊm(ɛ)nˈteɪʃn",
				Audio:            "https://example.com/documentation.mp3",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	es := []entry.Entry{e}

	return getEntriesFixture{h, req, resp, es}
}

// Fixture for the "unmarshal-at-best-effort" case.
func bestEffort() getEntriesFixture {
	h, herr := headword.New("best-effort")
	if herr != nil {
		panic(herr)
	}

	req := odapi.GetEntriesRequest{
		SourceLang: "en-gb",
		WordId:     "best-effort",
	}

	resp := odapi.GetEntriesResponse{
		Results: []odapi.HeadwordEntry{
			{
				LexicalEntries: []odapi.LexicalEntry{
					{
						Entries: []odapi.Entry{
							{
								Senses: []odapi.Sense{
									{
										Definitions: []string{"provides no guarantee"},
									},
								},
							},
						},
						LexicalCategory: odapi.LexicalCategory{
							Id: "unknown",
						},
					},
				},
			},
			{
				LexicalEntries: []odapi.LexicalEntry{
					{
						Entries: []odapi.Entry{
							{},
						},
						LexicalCategory: odapi.LexicalCategory{
							Id: "noun",
						},
					},
				},
			},
			{
				LexicalEntries: []odapi.LexicalEntry{
					{
						Entries: []odapi.Entry{
							{
								Senses: []odapi.Sense{
									{
										Definitions: []string{"in a manner that", "provides no guarantee"},
									},
								},
							},
						},
						LexicalCategory: odapi.LexicalCategory{
							Id: "adjective",
						},
					},
				},
			},
			{
				LexicalEntries: []odapi.LexicalEntry{
					{
						Entries: []odapi.Entry{
							{
								Senses: []odapi.Sense{
									{
										Definitions: []string{"since when has best-effort become a combining form?"},
										Subsenses: []odapi.Sense{
											{
												CrossReferenceMarkers: []string{"refer to that subsense as well"},
											},
										},
									},
									{
										CrossReferenceMarkers: []string{"refer to that sense as well"},
									},
								},
							},
						},
						LexicalCategory: odapi.LexicalCategory{
							Id: "combiningForm",
						},
					},
				},
			},
		},
	}

	e1, eerr := entry.New(entry.NewInput{
		Headword:        "best-effort",
		LexicalCategory: "adjective",
		Senses: []entry.NewSenseInput{
			{
				Description: "in a manner that; provides no guarantee",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	e2, eerr := entry.New(entry.NewInput{
		Headword:        "best-effort",
		LexicalCategory: "combining form",
		Senses: []entry.NewSenseInput{
			{
				Description: "since when has best-effort become a combining form?",
				Subsenses: []entry.NewSubsenseInput{
					{
						Description: "refer to that subsense as well",
					},
				},
			},
			{
				Description: "refer to that sense as well",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	es := []entry.Entry{e1, e2}

	return getEntriesFixture{h, req, resp, es}
}
