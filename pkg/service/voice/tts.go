package voice

import (
	"github.com/col3name/tts/pkg/service/moderation"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type SpeechVoiceService interface {
	Speak(text string) error
}

type HtgoTtsService struct {
	speech   htgotts.Speech
	language string
	filter   moderation.Filter
}

func NewHtgoTtsService(language string, filter moderation.Filter) *HtgoTtsService {
	h := new(HtgoTtsService)
	if len(language) == 0 {
		language = voices.English
	}

	h.language = language
	h.speech = htgotts.Speech{Folder: "audio", Language: language}
	h.filter = filter
	return h
}

func (s *HtgoTtsService) Speak(text string) error {
	result := s.filter.Moderate(text)
	return s.speech.Speak(result)
}
