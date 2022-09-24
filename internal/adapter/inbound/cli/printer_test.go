package cli_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/central182/odie/internal/adapter/inbound/cli"
	"github.com/central182/odie/internal/domain/application"
	application_mock "github.com/central182/odie/internal/domain/application/mock"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

func TestPrintEntriesOfHeadword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("The PrintEntriesOfHeadword method", func(t *testing.T) {
		t.Run("consumes from the Application and prints the return value appropriately,", func(t *testing.T) {
			h, es, txt := fixture()
			p := cli.NewPrinter(func() application.Application {
				app := application_mock.NewMockApplication(ctrl)
				app.EXPECT().GetEntriesByHeadword(h).Return(es, nil)
				return app
			})

			actualTxt, err := p.PrintEntriesOfHeadword(h.Spelling())
			assert.NoError(t, err)
			assert.Equal(t, txt, actualTxt)
		})

		t.Run("and reports errors if any.", func(t *testing.T) {
			h, _, _ := fixture()
			p := cli.NewPrinter(func() application.Application {
				app := application_mock.NewMockApplication(ctrl)
				err := application_mock.NewMockGetEntriesByHeadwordError(ctrl)
				err.EXPECT().Error().Return("Oops")
				app.EXPECT().GetEntriesByHeadword(h).Return(nil, err)
				return app
			})

			es, err := p.PrintEntriesOfHeadword(h.Spelling())
			assert.Empty(t, es)
			assert.ErrorContains(t, err, "Oops")
		})
	})
}

func fixture() (headword.Headword, []entry.Entry, string) {
	h, herr := headword.New("documentation")
	if herr != nil {
		panic(herr)
	}

	e1, eerr := entry.New(entry.NewInput{
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
	e2, eerr := entry.New(entry.NewInput{
		Headword:        "documentation",
		LexicalCategory: "verb",
		Senses: []entry.NewSenseInput{
			{
				Description: "documentation is not a verb :)",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	e3, eerr := entry.New(entry.NewInput{
		Headword:        "documentation",
		LexicalCategory: "interjection",
		Senses: []entry.NewSenseInput{
			{
				Description: "in a minute :)",
			},
		},
		Pronunciations: []entry.NewPronunciationInput{
			{
				PhoneticSpelling: "ˌdɒkjʊm(ɛ)nˈteɪʃn",
			},
			{
				PhoneticSpelling: "dɒk",
			},
		},
	})
	if eerr != nil {
		panic(eerr)
	}
	es := []entry.Entry{e1, e2, e3}

	txt := func() string {
		buf := new(bytes.Buffer)
		fmt.Fprintln(buf, color.New(color.Bold).Sprint("documentation"), "/ˌdɒkjʊm(ɛ)nˈteɪʃn/", color.New(color.Italic, color.FgYellow).Sprint("noun"))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "1", "some official material")
		fmt.Fprintln(buf, " ", color.New(color.Italic, color.Faint).Sprint("Complete the relevant documentation!"))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "1.1", "something accompanying a computer program")
		fmt.Fprintln(buf, "   ", color.New(color.Italic, color.Faint).Sprint("A user documentation."))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "2", "some process of classifying things")
		fmt.Fprintln(buf, " ", color.New(color.Italic, color.Faint).Sprint("Arrange the documentation of photographs!"))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, color.New(color.Bold).Sprint("documentation"), color.New(color.Italic, color.FgYellow).Sprint("verb"))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "1", "documentation is not a verb :)")
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, color.New(color.Bold).Sprint("documentation"), "/ˌdɒkjʊm(ɛ)nˈteɪʃn/, /dɒk/", color.New(color.Italic, color.FgYellow).Sprint("interjection"))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "1", "in a minute :)")
		fmt.Fprintln(buf)
		return buf.String()
	}()

	return h, es, txt
}
