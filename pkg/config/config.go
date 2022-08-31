package config

import (
	"github.com/col3name/tts/pkg/service/voice"
	"github.com/col3name/tts/pkg/util/separator"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const (
	MessageInvalidVolume        = "volume range [0; 2]"
	MessageFailedLoadConfigFile = "Error loading .env file"
	MessageNotSupportedLanguage = "Not supported Language."
)
const (
	EnvVolume            = "VOLUME"
	EnvUserIgnore        = "USER_IGNORE"
	EnvLangDetectEnabled = "LANG_DETECT_ENABLED"
	EnvTwitchChannel     = "TWITCH_CHANNEL"
	EnvRestAddress       = "SERVE_REST_ADDRESS"
	EnvLanguage          = "LANGUAGE"
	EnvModeration        = "MODERATION"
	EnvIgnoredWords      = "IGNORE"
)

type Config struct {
	Volume              float64
	LangDetectorEnabled bool
	WebViewAddress      string
	StaticApiAddress    string
	RestAddress         string
	Language            string
	ModerationWordPairs string
	IgnoreWords         string
	ChannelsList        []string
	UserBanList         []string
}

func NewConfig() *Config {
	return &Config{
		Volume:              1.0,
		LangDetectorEnabled: false,
		WebViewAddress:      "http://localhost:3000",
		StaticApiAddress:    ":3000",
		RestAddress:         ":8000",
	}
}

func (c *Config) Parse() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(MessageFailedLoadConfigFile)
	}
	c.setServeRestAddress()
	c.setChannelsList()
	c.setLanguage()
	c.setModerationWordPair()
	c.setIgnoreWords()
	c.setLanguageDetectorEnabled()
	c.setUserBanList()
	c.setVolume()
}

func (c *Config) setVolume() {
	volumeString := os.Getenv(EnvVolume)
	if len(volumeString) == 0 {
		return
	}
	num, err := strconv.ParseFloat(volumeString, 10)
	if err != nil {
		log.Fatal(MessageInvalidVolume)
	}
	if num < 0 || num > 2.0 {
		log.Fatal(MessageInvalidVolume)
	}
	c.Volume = num
}

func (c *Config) setUserBanList() {
	userIgnore := os.Getenv(EnvUserIgnore)
	c.UserBanList = strings.Split(userIgnore, separator.Item)
	c.UserBanList = append(c.UserBanList, c.ChannelsList...)
}

func (c *Config) setLanguageDetectorEnabled() {
	langDetectorEnabledString := os.Getenv(EnvLangDetectEnabled)
	if langDetectorEnabledString == "" || langDetectorEnabledString == "false" {
		c.LangDetectorEnabled = false
	} else if langDetectorEnabledString == "true" {
		c.LangDetectorEnabled = true
	}
}

func (c *Config) setChannelsList() {
	channels := os.Getenv(EnvTwitchChannel)
	if len(channels) == 0 {
		log.Fatal(MessageFailedLoadConfigFile)
	}
	c.ChannelsList = strings.Split(channels, separator.Pair)
}

func (c *Config) setServeRestAddress() {
	serveRestAddress := os.Getenv(EnvRestAddress)
	if len(c.RestAddress) != 0 {
		c.RestAddress = serveRestAddress
	}
}

func (c *Config) setLanguage() {
	c.Language = os.Getenv(EnvLanguage)
	if !voice.IsSupported(c.Language) {
		log.Fatal(MessageNotSupportedLanguage)
	}
}

func (c *Config) setModerationWordPair() {
	c.ModerationWordPairs = os.Getenv(EnvModeration)
}

func (c *Config) setIgnoreWords() {
	c.IgnoreWords = os.Getenv(EnvIgnoredWords)
}
