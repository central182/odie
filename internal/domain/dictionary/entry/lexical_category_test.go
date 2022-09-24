package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicalCategory(t *testing.T) {
	t.Run("A LexicalCategory can't be constructed from an unknown name.", func(t *testing.T) {
		lc, err := newLexicalCategory("unknown lexical category")
		assert.Nil(t, lc)
		assert.True(t, err.HasUnknownLexicalCategory())
	})

	t.Run("A LexicalCategory can be reflexively constructed from its own name.", func(t *testing.T) {
		// Which is to say, that a LexicalCategory can be constructed from a known, defined name.

		var tests = []struct {
			lc LexicalCategory
		}{
			{Adjective{}},
			{Adverb{}},
			{CombiningForm{}},
			{Conjunction{}},
			{Contraction{}},
			{Determiner{}},
			{Idiomatic{}},
			{Interjection{}},
			{Noun{}},
			{Numeral{}},
			{Other{}},
			{Particle{}},
			{Predeterminer{}},
			{Prefix{}},
			{Preposition{}},
			{Pronoun{}},
			{Residual{}},
			{Suffix{}},
			{Verb{}},
		}
		for _, tt := range tests {
			lc, err := newLexicalCategory(tt.lc.Name())
			assert.Nil(t, err)
			assert.Equal(t, tt.lc, lc)
		}
	})
}
