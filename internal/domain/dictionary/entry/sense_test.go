package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSense(t *testing.T) {
	description := "Explain it!"
	example := "A first example."
	example2 := "Another example."
	subsenseInput := NewSubsenseInput{Description: "explaining the subsense"}
	subsense := func() Subsense {
		sub, serr := newSubsense(subsenseInput)
		assert.Nil(t, serr)
		return sub
	}()
	subsense2Input := NewSubsenseInput{Description: "explaining the subsense again"}
	subsense2 := func() Subsense {
		sub, serr := newSubsense(subsense2Input)
		assert.Nil(t, serr)
		return sub
	}()

	t.Run("A Sense can't be constructed", func(t *testing.T) {
		t.Run("if an empty description is provided", func(t *testing.T) {
			s, err := newSense(NewSenseInput{})
			assert.Nil(t, s)
			assert.True(t, err.HasEmptyDescription())
			assert.Nil(t, err.EmptyExamples())
			assert.Nil(t, err.SubsenseProblems())
		})

		t.Run("if at least one empty example is provided", func(t *testing.T) {
			s, err := newSense(NewSenseInput{Description: description, Examples: []string{example, ""}})
			assert.Nil(t, s)
			assert.False(t, err.HasEmptyDescription())
			assert.Equal(t, []bool{false, true}, err.EmptyExamples())
			assert.Nil(t, err.SubsenseProblems())
		})

		t.Run("if at least one problematic Subsense is provided", func(t *testing.T) {
			s, err := newSense(NewSenseInput{Description: description, Subsenses: []NewSubsenseInput{subsenseInput, {}}})
			assert.Nil(t, s)
			assert.False(t, err.HasEmptyDescription())
			assert.Nil(t, err.EmptyExamples())
			assert.Len(t, err.SubsenseProblems(), 2)
			assert.Nil(t, err.SubsenseProblems()[0])
			assert.NotNil(t, err.SubsenseProblems()[1])
		})
	})

	t.Run("A valid Sense", func(t *testing.T) {
		t.Run("contains a non-empty description.", func(t *testing.T) {
			s, err := newSense(NewSenseInput{Description: description})
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
				s, err := newSense(NewSenseInput{Description: description, Examples: tt.examples})
				assert.Nil(t, err)

				es, ok := s.Examples()
				assert.Equal(t, tt.hasExamples, ok)
				assert.Equal(t, tt.examples, es)
			}
		})

		t.Run("may or may not contain Subsenses.", func(t *testing.T) {
			var tests = []struct {
				subsenseInputs []NewSubsenseInput
				hasSubsenses   bool
				subsenses      []Subsense
			}{
				{nil, false, nil},
				{[]NewSubsenseInput{subsenseInput}, true, []Subsense{subsense}},
				{[]NewSubsenseInput{subsenseInput, subsense2Input}, true, []Subsense{subsense, subsense2}},
			}
			for _, tt := range tests {
				s, err := newSense(NewSenseInput{Description: description, Subsenses: tt.subsenseInputs})
				assert.Nil(t, err)

				subs, ok := s.Subsenses()
				assert.Equal(t, tt.hasSubsenses, ok)
				assert.Equal(t, tt.subsenses, subs)
			}
		})
	})

	t.Run("The examples of a Sense", func(t *testing.T) {
		t.Run("are distinct from the original argument passed to the constructor.", func(t *testing.T) {
			examples := []string{example}

			s, err := newSense(NewSenseInput{Description: description, Examples: examples})
			assert.Nil(t, err)

			examples[0] = example2

			es, ok := s.Examples()
			assert.True(t, ok)
			assert.Equal(t, []string{example}, es)
		})

		t.Run("can't be used to alter the Sense's internal state.", func(t *testing.T) {
			s, err := newSense(NewSenseInput{Description: description, Examples: []string{example}})
			assert.Nil(t, err)

			es, ok := s.Examples()
			assert.True(t, ok)
			es[0] = example2

			es, ok = s.Examples()
			assert.True(t, ok)
			assert.Equal(t, []string{example}, es)
		})
	})

	t.Run("The Subsenses of a Sense", func(t *testing.T) {
		t.Run("can't be used to alter the Sense's internal state.", func(t *testing.T) {
			s, err := newSense(NewSenseInput{Description: description, Subsenses: []NewSubsenseInput{subsenseInput}})
			assert.Nil(t, err)

			subs, ok := s.Subsenses()
			assert.True(t, ok)
			subs[0] = subsense2

			subs, ok = s.Subsenses()
			assert.True(t, ok)
			assert.Equal(t, []Subsense{subsense}, subs)
		})
	})
}
