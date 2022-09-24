package headword

type Headword interface {
	// A Headword is represented by a non-empty spelling.
	Spelling() string
}

func New(spelling string) (Headword, NewError) {
	err := &newError{}

	if spelling == "" {
		err.setHasEmptySpelling()
		return nil, err
	}

	return headword{spelling: spelling}, nil
}

type headword struct {
	spelling string
}

func (h headword) Spelling() string {
	return h.spelling
}
