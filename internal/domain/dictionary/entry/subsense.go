package entry

type Subsense interface {
	// A Subsense is explained in a non-empty description.
	Description() string

	// There may exist some examples related to a Subsense, all of which should be non-empty sentences.
	Examples() ([]string, bool)
}

type NewSubsenseInput struct {
	Description string
	Examples    []string
}

func newSubsense(input NewSubsenseInput) (Subsense, NewSubsenseError) {
	err := &newSubsenseError{}

	if input.Description == "" {
		err.setHasEmptyDescription()
	}

	var es []string
	for _, e := range input.Examples {
		err.appendEmptyExample(e == "")
		es = append(es, e)
	}

	if err.touched() {
		return nil, err
	}

	return subsense{
		description: input.Description,
		examples:    es,
	}, nil
}

type subsense struct {
	description string
	examples    []string
}

func (s subsense) Description() string {
	return s.description
}

func (s subsense) Examples() ([]string, bool) {
	if len(s.examples) == 0 {
		return nil, false
	}

	clones := make([]string, len(s.examples))
	copy(clones, s.examples)
	return clones, true
}
