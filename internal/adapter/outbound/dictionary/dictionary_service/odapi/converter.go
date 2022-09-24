package dictionary_service_odapi

import (
	"strings"

	"github.com/central182/odie/internal/adapter/outbound/common/odapi"
	"github.com/central182/odie/internal/domain/dictionary/entry"
)

func unmarshalPronunciation(
	p odapi.Pronunciation,
) entry.NewPronunciationInput {
	return entry.NewPronunciationInput{
		PhoneticSpelling: p.PhoneticSpelling,
		Audio:            p.AudioFile,
	}
}

func unmarshalPronunciations(
	subs []odapi.Pronunciation,
) []entry.NewPronunciationInput {
	var result []entry.NewPronunciationInput
	for _, sub := range subs {
		result = append(result, unmarshalPronunciation(sub))
	}
	return result
}

func unmarshalSubsense(
	sub odapi.Sense,
) entry.NewSubsenseInput {
	return entry.NewSubsenseInput{
		Description: strings.Join(sub.Definitions, "; "),
		Examples: func() (egs []string) {
			for _, eg := range sub.Examples {
				egs = append(egs, eg.Text)
			}
			return
		}(),
	}
}

func unmarshalSubsenses(
	subs []odapi.Sense,
) []entry.NewSubsenseInput {
	var result []entry.NewSubsenseInput
	for _, sub := range subs {
		result = append(result, unmarshalSubsense(sub))
	}
	return result
}

func unmarshalSense(
	s odapi.Sense,
) entry.NewSenseInput {
	return entry.NewSenseInput{
		Description: strings.Join(s.Definitions, "; "),
		Examples: func() (egs []string) {
			for _, eg := range s.Examples {
				egs = append(egs, eg.Text)
			}
			return
		}(),
		Subsenses: unmarshalSubsenses(s.Subsenses),
	}
}

func unmarshalSenses(
	ss []odapi.Sense,
) []entry.NewSenseInput {
	var result []entry.NewSenseInput
	for _, s := range ss {
		result = append(result, unmarshalSense(s))
	}
	return result
}

func unmarshalEntries(
	headword string,
	lexicalCategory string,
	es []odapi.Entry,
	err *getEntriesError,
) []entry.Entry {
	var result []entry.Entry
	for _, e := range es {
		newInput := entry.NewInput{
			Headword:        headword,
			LexicalCategory: lexicalCategory,
			Senses:          unmarshalSenses(e.Senses),
			Pronunciations:  unmarshalPronunciations(e.Pronunciations),
		}

		e, eerr := entry.New(newInput)
		if eerr == nil {
			result = append(result, e)
		} else {
			err.appendUnmarshallingFailure(eerr)
		}
	}
	return result
}

func unmarshalGetEntriesResponse(
	headword string,
	resp odapi.GetEntriesResponse,
	err *getEntriesError,
) []entry.Entry {
	var result []entry.Entry
	for _, he := range resp.Results {
		for _, le := range he.LexicalEntries {
			result = append(result, unmarshalEntries(headword, le.LexicalCategory.Id, le.Entries, err)...)
		}
	}
	return result
}
