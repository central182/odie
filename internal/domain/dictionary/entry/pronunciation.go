package entry

import (
	"net/url"
)

type Pronunciation interface {
	// A Pronunciation is represented orthographically by a phonetic spelling in IPA.
	PhoneticSpelling() string

	// A Pronunciation may be represented acoustically by a downloadable audio file.
	Audio() (url.URL, bool)
}

type NewPronunciationInput struct {
	PhoneticSpelling string
	Audio            string
}

func newPronunciation(input NewPronunciationInput) (Pronunciation, NewPronunciationError) {
	err := &newPronunciationError{}

	if input.PhoneticSpelling == "" {
		err.setHasEmptyPhoneticSpelling()
	}

	u := func() url.URL {
		if input.Audio == "" {
			return url.URL{}
		}

		u, uerr := url.Parse(input.Audio)
		if uerr != nil {
			err.setAudioProblem(uerr)
			return url.URL{}
		}

		return *u
	}()

	if err.touched() {
		return nil, err
	}

	return pronunciation{
		phoneticSpelling: input.PhoneticSpelling,
		audio:            u,
	}, nil
}

type pronunciation struct {
	phoneticSpelling string
	audio            url.URL
}

func (p pronunciation) PhoneticSpelling() string {
	return p.phoneticSpelling
}

func (p pronunciation) Audio() (url.URL, bool) {
	var zeroUrl url.URL
	if p.audio == zeroUrl {
		return zeroUrl, false
	}

	return p.audio, true
}
