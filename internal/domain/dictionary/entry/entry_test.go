package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central182/odie/internal/domain/dictionary/headword"
)

func TestEntry(t *testing.T) {
	headwordInput := "example"
	headword := func() headword.Headword {
		h, err := headword.New(headwordInput)
		assert.Nil(t, err)
		return h
	}()
	lexicalCategoryInput := "noun"
	lexicalCategory := func() LexicalCategory {
		l, err := newLexicalCategory(lexicalCategoryInput)
		assert.Nil(t, err)
		return l
	}()
	senseInput := NewSenseInput{Description: "Explain it!"}
	sense := func() Sense {
		s, err := newSense(senseInput)
		assert.Nil(t, err)
		return s
	}()
	sense2Input := NewSenseInput{Description: "explaining again"}
	sense2 := func() Sense {
		s, err := newSense(sense2Input)
		assert.Nil(t, err)
		return s
	}()
	pronunciationInput := NewPronunciationInput{PhoneticSpelling: "ɪɡˈzɑːmpl"}
	pronunciation := func() Pronunciation {
		p, err := newPronunciation(pronunciationInput)
		assert.Nil(t, err)
		return p
	}()
	pronunciation2Input := NewPronunciationInput{PhoneticSpelling: "ɛɡˈzɑːmpl"}
	pronunciation2 := func() Pronunciation {
		p, err := newPronunciation(pronunciation2Input)
		assert.Nil(t, err)
		return p
	}()

	t.Run("An Entry can't be constructed", func(t *testing.T) {
		t.Run("if an empty Headword is provided", func(t *testing.T) {
			e, err := New(NewInput{LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}})
			assert.Nil(t, e)
			assert.NotNil(t, err.HeadwordProblem())
			assert.Nil(t, err.LexicalCategoryProblem())
			assert.False(t, err.HasNoSenses())
			assert.Nil(t, err.SenseProblems())
			assert.Nil(t, err.PronunciationProblems())
			assert.NotEmpty(t, err.Error())
		})

		t.Run("if an unknown LexicalCategory is provided", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, Senses: []NewSenseInput{senseInput}})
			assert.Nil(t, e)
			assert.Nil(t, err.HeadwordProblem())
			assert.NotNil(t, err.LexicalCategoryProblem())
			assert.False(t, err.HasNoSenses())
			assert.Nil(t, err.SenseProblems())
			assert.Nil(t, err.PronunciationProblems())
			assert.NotEmpty(t, err.Error())
		})

		t.Run("if no Sense is provided at all", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput})
			assert.Nil(t, e)
			assert.Nil(t, err.HeadwordProblem())
			assert.Nil(t, err.LexicalCategoryProblem())
			assert.True(t, err.HasNoSenses())
			assert.Nil(t, err.SenseProblems())
			assert.Nil(t, err.PronunciationProblems())
			assert.NotEmpty(t, err.Error())
		})

		t.Run("if at least one problematic Sense is provided", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput, {}}})
			assert.Nil(t, e)
			assert.Nil(t, err.HeadwordProblem())
			assert.Nil(t, err.LexicalCategoryProblem())
			assert.False(t, err.HasNoSenses())
			assert.Len(t, err.SenseProblems(), 2)
			assert.Nil(t, err.SenseProblems()[0])
			assert.NotNil(t, err.SenseProblems()[1])
			assert.Nil(t, err.PronunciationProblems())
			assert.NotEmpty(t, err.Error())
		})

		t.Run("if at least one problematic Pronunciation is provided", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}, Pronunciations: []NewPronunciationInput{pronunciationInput, {}}})
			assert.Nil(t, e)
			assert.Nil(t, err.HeadwordProblem())
			assert.Nil(t, err.LexicalCategoryProblem())
			assert.False(t, err.HasNoSenses())
			assert.Nil(t, err.SenseProblems())
			assert.Len(t, err.PronunciationProblems(), 2)
			assert.Nil(t, err.PronunciationProblems()[0])
			assert.NotNil(t, err.PronunciationProblems()[1])
			assert.NotEmpty(t, err.Error())
		})
	})

	t.Run("A valid Entry", func(t *testing.T) {
		t.Run("contains a valid Headword.", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}})
			assert.Nil(t, err)
			assert.Equal(t, headword, e.Headword())
		})

		t.Run("contains a valid LexicalCategory.", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}})
			assert.Nil(t, err)
			assert.Equal(t, lexicalCategory, e.LexicalCategory())
		})

		t.Run("contains at least one valid Senses.", func(t *testing.T) {
			var tests = []struct {
				senseInputs []NewSenseInput
				senses      []Sense
			}{
				{[]NewSenseInput{senseInput}, []Sense{sense}},
				{[]NewSenseInput{senseInput, sense2Input}, []Sense{sense, sense2}},
			}
			for _, tt := range tests {
				e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: tt.senseInputs})
				assert.Nil(t, err)
				assert.Equal(t, tt.senses, e.Senses())
			}
		})

		t.Run("may or may not contain Pronunciations.", func(t *testing.T) {
			var tests = []struct {
				pronunciationInputs []NewPronunciationInput
				hasPronuncations    bool
				pronunciations      []Pronunciation
			}{
				{nil, false, nil},
				{[]NewPronunciationInput{pronunciationInput}, true, []Pronunciation{pronunciation}},
				{[]NewPronunciationInput{pronunciationInput, pronunciation2Input}, true, []Pronunciation{pronunciation, pronunciation2}},
			}
			for _, tt := range tests {
				e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}, Pronunciations: tt.pronunciationInputs})
				assert.Nil(t, err)

				ps, ok := e.Pronunciations()
				assert.Equal(t, tt.hasPronuncations, ok)
				assert.Equal(t, tt.pronunciations, ps)
			}
		})
	})

	t.Run("The Senses of an Entry", func(t *testing.T) {
		t.Run("can't be used to alter the Sense's internal state.", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}})
			assert.Nil(t, err)

			ss := e.Senses()
			ss[0] = sense2

			ss = e.Senses()
			assert.Equal(t, []Sense{sense}, ss)
		})
	})

	t.Run("The Pronunciations of an Entry", func(t *testing.T) {
		t.Run("can't be used to alter the Sense's internal state.", func(t *testing.T) {
			e, err := New(NewInput{Headword: headwordInput, LexicalCategory: lexicalCategoryInput, Senses: []NewSenseInput{senseInput}, Pronunciations: []NewPronunciationInput{pronunciationInput}})
			assert.Nil(t, err)

			ps, ok := e.Pronunciations()
			assert.True(t, ok)
			ps[0] = pronunciation2

			ps, ok = e.Pronunciations()
			assert.True(t, ok)
			assert.Equal(t, []Pronunciation{pronunciation}, ps)
		})
	})
}
