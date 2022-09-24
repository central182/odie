package headword_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central182/odie/internal/domain/dictionary/headword"
)

func TestHeadword(t *testing.T) {
	t.Run("A Headword can't be constructed from an empty spelling.", func(t *testing.T) {
		h, err := headword.New("")
		assert.Nil(t, h)
		assert.True(t, err.HasEmptySpelling())
		assert.NotEmpty(t, err.Error())
	})

	t.Run("A valid Headword is represented by a non-empty spelling.", func(t *testing.T) {
		var tests = []struct {
			spelling string
		}{
			{"word"},

			// There's no easy way to test whether a given string will likely appear in a dictionary or not.
			// And therefore,
			{"Any string that is not empty is considered valid."},
		}
		for _, tt := range tests {
			h, err := headword.New(tt.spelling)
			assert.Nil(t, err)
			assert.Equal(t, tt.spelling, h.Spelling())
		}
	})
}
