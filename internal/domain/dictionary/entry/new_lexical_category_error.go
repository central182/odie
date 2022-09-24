package entry

type NewLexicalCategoryError interface {
	HasUnknownLexicalCategory() bool
}

type newLexicalCategoryError struct {
	hasUnknownLexicalCategory bool
}

func (e *newLexicalCategoryError) HasUnknownLexicalCategory() bool {
	return e.hasUnknownLexicalCategory
}

func (e *newLexicalCategoryError) touched() bool {
	return *e != newLexicalCategoryError{}
}

func (e *newLexicalCategoryError) setHasUnknownLexicalCategory() {
	e.hasUnknownLexicalCategory = true
}
