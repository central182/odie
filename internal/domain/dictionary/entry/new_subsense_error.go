package entry

type NewSubsenseError interface {
	HasEmptyDescription() bool
	EmptyExamples() []bool
}

type newSubsenseError struct {
	hasEmptyDescription bool
	emptyExamples       []bool
}

func (e *newSubsenseError) HasEmptyDescription() bool {
	return e.hasEmptyDescription
}

func (e *newSubsenseError) EmptyExamples() []bool {
	for _, ee := range e.emptyExamples {
		if ee {
			return e.emptyExamples
		}
	}

	return nil
}

func (e *newSubsenseError) touched() bool {
	if e.hasEmptyDescription {
		return true
	}

	for _, ee := range e.emptyExamples {
		if ee {
			return true
		}
	}

	return false
}

func (e *newSubsenseError) setHasEmptyDescription() {
	e.hasEmptyDescription = true
}

func (e *newSubsenseError) appendEmptyExample(isEmpty bool) {
	e.emptyExamples = append(e.emptyExamples, isEmpty)
}
