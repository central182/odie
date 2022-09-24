package entry

type Sense interface {
	// A Sense is explained in a non-empty description.
	Description() string

	// There may exist some examples related to a Sense, all of which should be non-empty sentences.
	Examples() ([]string, bool)

	// A Sense may further contain some Subsenses.
	Subsenses() ([]Subsense, bool)
}

type NewSenseInput struct {
	Description string
	Examples    []string
	Subsenses   []NewSubsenseInput
}

func newSense(input NewSenseInput) (Sense, NewSenseError) {
	err := &newSenseError{}

	if input.Description == "" {
		err.setHasEmptyDescription()
	}

	var es []string
	for _, e := range input.Examples {
		err.appendEmptyExample(e == "")
		es = append(es, e)
	}

	var subs []Subsense
	for _, subi := range input.Subsenses {
		sub, serr := newSubsense(subi)
		err.appendSubsenseProblem(serr)
		subs = append(subs, sub)
	}

	if err.touched() {
		return nil, err
	}

	return sense{
		description: input.Description,
		examples:    es,
		subsenses:   subs,
	}, nil
}

type sense struct {
	description string
	examples    []string
	subsenses   []Subsense
}

func (s sense) Description() string {
	return s.description
}

func (s sense) Examples() ([]string, bool) {
	if len(s.examples) == 0 {
		return nil, false
	}

	clones := make([]string, len(s.examples))
	copy(clones, s.examples)
	return clones, true
}

func (s sense) Subsenses() ([]Subsense, bool) {
	if len(s.subsenses) == 0 {
		return nil, false
	}
	clones := make([]Subsense, len(s.subsenses))
	copy(clones, s.subsenses)
	return clones, true
}
