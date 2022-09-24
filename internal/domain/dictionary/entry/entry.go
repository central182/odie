package entry

import "github.com/central182/odie/internal/domain/dictionary/headword"

// An Entry is a detailed explanation of a Headword of a specific LexicalCategory.
// A pair of Headword and LexicalCategory doesn't uniquely identify an Entry, however,
// because there may exist homographs of the same LexicalCategory.
type Entry interface {
	Headword() headword.Headword
	LexicalCategory() LexicalCategory

	// An Entry contains at least one Sense.
	Senses() []Sense

	// There may be one or more Pronunciations related to an Entry.
	Pronunciations() ([]Pronunciation, bool)
}

type NewInput struct {
	Headword        string
	LexicalCategory string
	Senses          []NewSenseInput
	Pronunciations  []NewPronunciationInput
}

func New(input NewInput) (Entry, NewError) {
	err := &newError{}

	h, herr := headword.New(input.Headword)
	err.setHeadwordProblem(herr)

	l, lerr := newLexicalCategory(input.LexicalCategory)
	err.setLexicalCategoryProblem(lerr)

	if len(input.Senses) == 0 {
		err.setHasNoSenses()
	}

	var ss []Sense
	for _, si := range input.Senses {
		s, serr := newSense(si)
		err.appendSenseProblems(serr)
		ss = append(ss, s)
	}

	var ps []Pronunciation
	for _, pi := range input.Pronunciations {
		p, perr := newPronunciation(pi)
		err.appendPronunciationProblems(perr)
		ps = append(ps, p)
	}

	if err.touched() {
		return nil, err
	}

	return entry{
		headword:        h,
		lexicalCategory: l,
		senses:          ss,
		pronunciations:  ps,
	}, nil
}

type entry struct {
	headword        headword.Headword
	lexicalCategory LexicalCategory
	senses          []Sense
	pronunciations  []Pronunciation
}

func (e entry) Headword() headword.Headword {
	return e.headword
}

func (e entry) LexicalCategory() LexicalCategory {
	return e.lexicalCategory
}

func (e entry) Senses() []Sense {
	clones := make([]Sense, len(e.senses))
	copy(clones, e.senses)
	return clones
}

func (e entry) Pronunciations() ([]Pronunciation, bool) {
	if len(e.pronunciations) == 0 {
		return nil, false
	}

	clones := make([]Pronunciation, len(e.pronunciations))
	copy(clones, e.pronunciations)
	return clones, true
}
