package entry

type NewSenseError interface {
	HasEmptyDescription() bool
	EmptyExamples() []bool
	SubsenseProblems() []NewSubsenseError
}

type newSenseError struct {
	hasEmptyDescription bool
	emptyExamples       []bool
	subsenseProblems    []NewSubsenseError
}

func (e *newSenseError) HasEmptyDescription() bool {
	return e.hasEmptyDescription
}

func (e *newSenseError) EmptyExamples() []bool {
	for _, ee := range e.emptyExamples {
		if ee {
			return e.emptyExamples
		}
	}

	return nil
}

func (e *newSenseError) SubsenseProblems() []NewSubsenseError {
	for _, sp := range e.subsenseProblems {
		if sp != nil {
			return e.subsenseProblems
		}
	}

	return nil
}

func (e *newSenseError) touched() bool {
	if e.hasEmptyDescription {
		return true
	}

	for _, ee := range e.emptyExamples {
		if ee {
			return true
		}
	}

	for _, sp := range e.subsenseProblems {
		if sp != nil {
			return true
		}
	}

	return false
}

func (e *newSenseError) setHasEmptyDescription() {
	e.hasEmptyDescription = true
}

func (e *newSenseError) appendEmptyExample(isEmpty bool) {
	e.emptyExamples = append(e.emptyExamples, isEmpty)
}

func (e *newSenseError) appendSubsenseProblem(err NewSubsenseError) {
	e.subsenseProblems = append(e.subsenseProblems, err)
}
