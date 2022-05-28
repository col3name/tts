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
	volume   int
}

func NewHtgoTtsService(language string, filter moderation.Filter, volume int) *HtgoTtsService {
	h := new(HtgoTtsService)
	if len(language) == 0 {
		language = voices.English
	}
	if volume < 0 || volume > 15 {
		return nil
	}

	h.language = language
	h.volume = volume
	h.speech = htgotts.Speech{Folder: "audio", Language: language, Volume: volume}
	h.filter = filter
	return h
}

func (s *HtgoTtsService) Speak(text string) error {
	result := s.filter.Moderate(text)
	if len(result) < 3 {
		return nil
	}
	return s.speech.Speak(result)
}
