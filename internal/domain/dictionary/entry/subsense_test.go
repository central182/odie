package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsense(t *testing.T) {
	description := "Explain it!"
	example := "A first example."
	example2 := "Another example."

	t.Run("A Subsense can't be constructed", func(t *testing.T) {
		t.Run("if an empty description is provided", func(t *testing.T) {
			s, err := newSubsense(NewSubsenseInput{})
			assert.Nil(t, s)
			assert.True(t, err.HasEmptyDescription())
			assert.Nil(t, err.EmptyExamples())
		})

		t.Run("if at least one empty example is provided", func(t *testing.T) {
			s, err := newSubsense(NewSubsenseInput{Description: description, Examples: []string{example, ""}})
			assert.Nil(t, s)
			assert.False(t, err.HasEmptyDescription())
			assert.Equal(t, []bool{false, true}, err.EmptyExamples())
		})
	})

	t.Run("A valid Subsense", func(t *testing.T) {
		t.Run("contains a non-empty description.", func(t *testing.T) {
			s, err := newSubsense(NewSubsenseInput{Description: description})
			assert.Nil(t, err)
			assert.Equal(t, description, s.Description())
		})

		t.Run("may or may not contain examples. If it does, all of the examples are non-empty sentences.", func(t *testing.T) {
			var tests = []struct {
				examples    []string
				hasExamples bool
			}{
				{nil, false},
				{[]string{example}, true},
				{[]string{example, example2}, true},
			}
			for _, tt := range tests {
				s, err := newSubsense(NewSubsenseInput{Description: description, Examples: tt.examples})
				assert.Nil(t, err)

				es, ok := s.Examples()
				assert.Equal(t, tt.hasExamples, ok)
				assert.Equal(t, tt.examples, es)
			}
		})
	})

	t.Run("The examples of a Subsense", func(t *testing.T) {
		t.Run("are distinct from the original argument passed to the constructor.", func(t *testing.T) {
			examples := []string{example}

			s, err := newSubsense(NewSubsenseInput{Description: description, Examples: examples})
			assert.Nil(t, err)

			examples[0] = example2

			es, ok := s.Examples()
			assert.True(t, ok)
			assert.Equal(t, []string{example}, es)
		})

		t.Run("can't be used to alter the Subsense's internal state.", func(t *testing.T) {
			s, err := newSubsense(NewSubsenseInput{Description: description, Examples: []string{example}})
			assert.Nil(t, err)

			es, ok := s.Examples()
			assert.True(t, ok)
			es[0] = example2

			es, ok = s.Examples()
			assert.True(t, ok)
			assert.Equal(t, []string{example}, es)
		})
	})
}
