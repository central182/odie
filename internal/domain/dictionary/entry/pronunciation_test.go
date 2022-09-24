package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPronunciation(t *testing.T) {
	phoneticSpelling := "wəːd"
	audio := "https://example.com/word.mp3"

	t.Run("A Pronunciation can't be constructed", func(t *testing.T) {
		t.Run("if an empty phonetic spelling is provided.", func(t *testing.T) {
			p, err := newPronunciation(NewPronunciationInput{})
			assert.Nil(t, p)
			assert.True(t, err.HasEmptyPhoneticSpelling())
			assert.Nil(t, err.AudioProblem())
		})

		t.Run("if a problematic audio URL is provided", func(t *testing.T) {
			p, err := newPronunciation(NewPronunciationInput{PhoneticSpelling: phoneticSpelling, Audio: ":"})
			assert.Nil(t, p)
			assert.False(t, err.HasEmptyPhoneticSpelling())
			assert.NotNil(t, err.AudioProblem())
		})
	})

	t.Run("A valid Pronunciation", func(t *testing.T) {
		t.Run("contains a non-empty phonetic spelling.", func(t *testing.T) {
			p, err := newPronunciation(NewPronunciationInput{PhoneticSpelling: phoneticSpelling})
			assert.Nil(t, err)
			assert.Equal(t, phoneticSpelling, p.PhoneticSpelling())
		})

		t.Run("may or may not refer to an audio.", func(t *testing.T) {
			var tests = []struct {
				audio    string
				hasAudio bool
			}{
				{"", false},
				{audio, true},
			}
			for _, tt := range tests {
				p, err := newPronunciation(NewPronunciationInput{PhoneticSpelling: phoneticSpelling, Audio: tt.audio})
				assert.Nil(t, err)

				a, ok := p.Audio()
				assert.Equal(t, tt.hasAudio, ok)
				assert.Equal(t, tt.audio, a.String())
			}
		})
	})
}
