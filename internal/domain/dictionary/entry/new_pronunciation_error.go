package entry

type NewPronunciationError interface {
	HasEmptyPhoneticSpelling() bool
	AudioProblem() error
}

type newPronunciationError struct {
	hasEmptyPhoneticSpelling bool
	audioProblem             error
}

func (e *newPronunciationError) HasEmptyPhoneticSpelling() bool {
	return e.hasEmptyPhoneticSpelling
}

func (e *newPronunciationError) AudioProblem() error {
	return e.audioProblem
}

func (e *newPronunciationError) touched() bool {
	return *e != newPronunciationError{}
}

func (e *newPronunciationError) setHasEmptyPhoneticSpelling() {
	e.hasEmptyPhoneticSpelling = true
}

func (e *newPronunciationError) setAudioProblem(err error) {
	e.audioProblem = err
}
