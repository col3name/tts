package lang_detection

import (
	"errors"
	"github.com/pemistahl/lingua-go"
	"strings"
)

var DefaultLanguages = []lingua.Language{
	lingua.English,
	lingua.Spanish,
	lingua.French,
	lingua.German,
	lingua.Russian,
	lingua.Chinese,
}

type Language string

var ErrUnsupportedLanguage = errors.New("unsupported language")

type LanguageDetectionService interface {
	Detect(text string) (*Language, error)
}

type LinguaDetectionService struct {
	detector lingua.LanguageDetector
}

func NewLinguaDetectionService(languages []lingua.Language) *LinguaDetectionService {
	return &LinguaDetectionService{
		detector: lingua.NewLanguageDetectorBuilder().
			FromLanguages(languages...).
			Build(),
	}
}

func (s *LinguaDetectionService) Detect(text string) (*Language, error) {
	lang, ok := s.detector.DetectLanguageOf(text)
	if !ok {
		return nil, ErrUnsupportedLanguage
	}

	language := Language(strings.ToLower(lang.IsoCode639_1().String()))
	return &language, nil
}
