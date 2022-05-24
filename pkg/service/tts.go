package service

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type SpeakService interface {
	Speak(text string) error
}

type HtgoTtsService struct {
	speech htgotts.Speech
}

func NewHtgoTtsService() *HtgoTtsService {
	h := new(HtgoTtsService)
	h.speech = htgotts.Speech{Folder: "audio", Language: voices.Russian}
	return h
}

func (s *HtgoTtsService) Speak(text string) error {
	return s.speech.Speak(text)
}
