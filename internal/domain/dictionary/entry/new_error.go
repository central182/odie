package entry

import (
	"github.com/central182/odie/internal/domain/dictionary/headword"
	"github.com/sanity-io/litter"
)

type NewError interface {
	error
	HeadwordProblem() headword.NewError
	LexicalCategoryProblem() NewLexicalCategoryError
	HasNoSenses() bool
	SenseProblems() []NewSenseError
	PronunciationProblems() []NewPronunciationError
}

type newError struct {
	headwordProblem        headword.NewError
	lexicalCategoryProblem NewLexicalCategoryError
	hasNoSenses            bool
	senseProblems          []NewSenseError
	pronunciationProblems  []NewPronunciationError
}

func (e *newError) Error() string {
	if !e.touched() {
		return ""
	}
	return litter.Options{}.Sdump(*e)
}

func (e *newError) HeadwordProblem() headword.NewError {
	return e.headwordProblem
}

func (e *newError) LexicalCategoryProblem() NewLexicalCategoryError {
	return e.lexicalCategoryProblem
}

func (e *newError) HasNoSenses() bool {
	return e.hasNoSenses
}

func (e *newError) SenseProblems() []NewSenseError {
	for _, sp := range e.senseProblems {
		if sp != nil {
			return e.senseProblems
		}
	}

	return nil
}

func (e *newError) PronunciationProblems() []NewPronunciationError {
	for _, pp := range e.pronunciationProblems {
		if pp != nil {
			return e.pronunciationProblems
		}
	}

	return nil
}

func (e *newError) touched() bool {
	if e.headwordProblem != nil {
		return true
	}

	if e.lexicalCategoryProblem != nil {
		return true
	}

	if e.hasNoSenses {
		return true
	}

	for _, sp := range e.senseProblems {
		if sp != nil {
			return true
		}
	}

	for _, pp := range e.pronunciationProblems {
		if pp != nil {
			return true
		}
	}

	return false
}

func (e *newError) setHeadwordProblem(err headword.NewError) {
	e.headwordProblem = err
}

func (e *newError) setLexicalCategoryProblem(err NewLexicalCategoryError) {
	e.lexicalCategoryProblem = err
}

func (e *newError) setHasNoSenses() {
	e.hasNoSenses = true
}

func (e *newError) appendSenseProblems(err NewSenseError) {
	e.senseProblems = append(e.senseProblems, err)
}

func (e *newError) appendPronunciationProblems(err NewPronunciationError) {
	e.pronunciationProblems = append(e.pronunciationProblems, err)
}
