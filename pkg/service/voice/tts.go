package voice

import (
	"github.com/col3name/tts/pkg/repo"
	lang_detection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/util"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"strings"
)

type SetterFromService interface {
	SetFrom(from string)
}

type SpeechVoiceService interface {
	SetterFromService
	Speak(text string) error
}

type HtgoTtsService struct {
	speech              htgotts.Speech
	language            string
	filter              moderation.Filter
	volume              int
	from                string
	repo                repo.SettingRepo
	langDetector        lang_detection.LanguageDetectionService
	langDetectorEnabled bool
}

func NewHtgoTtsService(language string, filter moderation.Filter, volume int, repo repo.SettingRepo, langDetectorEnabled bool, langDetector lang_detection.LanguageDetectionService) *HtgoTtsService {
	s := new(HtgoTtsService)
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

func (s *HtgoTtsService) SetFrom(from string) {
	s.from = from
}

func (s *HtgoTtsService) setLanguage(language string) {
	s.speech = NewSpeech(language, s.volume)
	s.language = language
}

func (s *HtgoTtsService) detectLanguage(text string) error {
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

func (s *HtgoTtsService) Speak(text string) error {
	if s.repo != nil {
		settingDb, err := s.repo.GetSettings()
		if err != nil {
			return err
		}
		s.filter = moderation.NewDefaultFilter(settingDb.ReplacementWordPair, settingDb.IgnoreWords, util.StrEnumerationToArray(settingDb.UserBanList))
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

	result := s.filter.Moderate(moderation.Message{From: s.from, Text: text})
	result = strings.Trim(result, " ")
	check := result[len(s.from):]
	suffix := strings.HasSuffix(check, "say    !")
	if suffix {
		return nil
	}
	return s.speech.Speak(result)
}

func NewSpeech(language string, volume int) htgotts.Speech {
	return htgotts.Speech{Folder: "audio", Language: language, Volume: volume}
}
