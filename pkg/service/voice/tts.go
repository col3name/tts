package voice

import (
	"github.com/col3name/gotts"
	"github.com/col3name/gotts/voices"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	lang_detection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/util"
	"strings"
)

type SpeechVoiceDTO struct {
	From string
	Text string
}
type SpeechVoiceService interface {
	Speak(text model.Message) error
}

type GoTtsService struct {
	speech              gotts.Speech
	language            string
	filter              moderation.Filter
	volume              float64
	from                string
	repo                repo.SettingRepo
	langDetector        lang_detection.LanguageDetectionService
	langDetectorEnabled bool
}

func NewSpeech(language string, volume float64) gotts.Speech {
	return gotts.Speech{Folder: "audio", Language: language, Volume: volume, Speed: 1}
}

func NewGoTtsService(language string, filter moderation.Filter, volume float64, repo repo.SettingRepo, langDetectorEnabled bool, langDetector lang_detection.LanguageDetectionService) *GoTtsService {
	s := new(GoTtsService)
	if len(language) == 0 {
		language = voices.English
	}
	if volume < 0 || volume > 15 {
		return nil
	}

	s.volume = volume
	s.speech = NewSpeech(language, volume)
	s.language = language
	s.filter = filter
	s.repo = repo
	s.langDetector = langDetector
	s.langDetectorEnabled = langDetectorEnabled
	return s
}

func (s *GoTtsService) Speak(text string) error {
	if s.repo != nil {
		settingDb, err := s.repo.GetSettings()
		if err != nil {
			return err
		}
		s.filter = moderation.NewDefaultFilter(settingDb.ReplacementWordPair,
			settingDb.IgnoreWords,
			util.StringOfEnumerationToArray(settingDb.UserBanList))
		if s.volume != settingDb.Volume {
			s.speech = NewSpeech(s.language, s.volume)
		}
		s.volume = settingDb.Volume
		s.langDetectorEnabled = settingDb.LanguageDetectorEnabled
		if err = s.repo.SaveSettings(settingDb); err != nil {
			return err
		}
		s.setLanguage(settingDb.Language)
	}

	if s.langDetectorEnabled {
		if err := s.detectLanguage(text); err != nil {
			return err
		}
	}

	result := s.filter.Moderate(model.Message{From: s.from, Text: text})
	result = strings.Trim(result, " ")
	fromLen := len(s.from)
	if fromLen > len(result) && len(result) == 0 {
		return nil
	}
	check := result[fromLen:]
	if strings.HasSuffix(check, "say    !") {
		return nil
	}
	return s.speech.Speak(result)
}

func (s *GoTtsService) setLanguage(language string) {
	s.speech = NewSpeech(language, s.volume)
	s.language = language
}

func (s *GoTtsService) detectLanguage(text string) error {
	langDetected, err := s.langDetector.Detect(text)
	if err != nil {
		return err
	}
	lang := string(*langDetected)
	s.setLanguage(lang)
	if s.repo != nil {
		settingDb, err := s.repo.GetSettings()
		if err != nil {
			return err
		}
		s.langDetectorEnabled = settingDb.LanguageDetectorEnabled
		settingDb.Language = lang
		return s.repo.SaveSettings(settingDb)
	}
	return nil
}
