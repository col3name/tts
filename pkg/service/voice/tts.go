package voice

import (
	"github.com/col3name/gotts"
	"github.com/col3name/gotts/voices"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	lang_detection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/util/array"
	"github.com/col3name/tts/pkg/util/number"
	"github.com/col3name/tts/pkg/util/stringss"
)

type SpeechVoiceDTO struct {
	From string
	Text string
}
type SpeechVoiceService interface {
	Speak(text model.Message) error
}

type GoTtsService struct {
	speech                  gotts.Speech
	language                string
	filter                  moderation.Filter
	volume                  float64
	repo                    repo.SettingRepo
	langDetector            lang_detection.LanguageDetectionService
	languageDetectorEnabled bool
}

const AudioFolder = gotts.AudioFolder

func NewGoTtsService(language string, filter moderation.Filter, volume float64, repo repo.SettingRepo, langDetectorEnabled bool, langDetector lang_detection.LanguageDetectionService) *GoTtsService {
	if !number.InRange(volume, 0, 15) {
		return nil
	}
	if stringss.Empty(language) {
		language = voices.English
	}

	s := new(GoTtsService)
	s.volume = volume
	s.speech = *gotts.NewSpeech(language, volume)
	s.language = language
	s.filter = filter
	s.repo = repo
	s.langDetector = langDetector
	s.languageDetectorEnabled = langDetectorEnabled
	return s
}

func (s *GoTtsService) Speak(message model.Message) error {
	err := s.updateSettings(message.Text)
	if err != nil {
		return err
	}
	text, err := s.moderateMessageFromUser(&message)
	if err != nil {
		return err
	}
	return s.speak(text)
}

func (s *GoTtsService) updateSettings(text string) error {
	if s.repo == nil {
		return nil
	}
	settingsFromDB, err := s.getSettingsFromDB()
	if err != nil {
		return err
	}

	s.updateFilter(settingsFromDB)
	s.updateVolume(settingsFromDB)
	s.updateLanguageDectector(settingsFromDB)

	if err = s.updateSettingsInDB(settingsFromDB); err != nil {
		return err
	}
	s.updateLanguage(settingsFromDB.Language)

	return s.updateDetectedLanguage(text)
}

func (s *GoTtsService) updateLanguageDectector(settingsFromDB *model.SettingDB) {
	s.languageDetectorEnabled = settingsFromDB.LanguageDetectorEnabled
}
func (s *GoTtsService) moderateMessageFromUser(message *model.Message) (string, error) {
	return s.filter.Moderate(message)
}

func (s *GoTtsService) speak(text string) error {
	return s.speech.Speak(text)
}

func (s *GoTtsService) getSettingsFromDB() (*model.SettingDB, error) {
	return s.repo.GetSettings()
}

func (s *GoTtsService) updateSettingsInDB(settingsFromDB *model.SettingDB) error {
	return s.repo.SaveSettings(settingsFromDB)
}

func (s *GoTtsService) updateVolume(settingDb *model.SettingDB) {
	if s.volume != settingDb.Volume {
		s.speech = *gotts.NewSpeech(s.language, s.volume)
	}
	s.volume = settingDb.Volume
}

func (s *GoTtsService) updateFilter(settingDb *model.SettingDB) {
	users := array.FromString(settingDb.UserBanList)
	s.filter = moderation.NewMessageFilter(settingDb.ReplacementWordPair, settingDb.IgnoreWords, users)
}

func (s *GoTtsService) updateDetectedLanguage(text string) error {
	if !s.languageDetectorEnabled {
		return nil
	}
	langDetected, err := s.langDetector.Detect(text)
	if err != nil {
		return err
	}
	language := string(*langDetected)
	s.updateLanguage(language)

	if s.repo == nil {
		return nil
	}

	settingDb, err := s.getSettingsFromDB()
	if err != nil {
		return err
	}
	s.updateLanguageDectector(settingDb)
	settingDb.Language = language
	return s.updateSettingsInDB(settingDb)
}

func (s *GoTtsService) updateLanguage(language string) {
	s.speech = *gotts.NewSpeech(language, s.volume)
	s.language = language
}
