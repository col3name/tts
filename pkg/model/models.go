package model

import (
	"errors"
	"github.com/col3name/tts/pkg/util/container"
	"github.com/col3name/tts/pkg/util/separator"
	"github.com/col3name/tts/pkg/util/stringss"
	"strings"
)

var (
	ErrorInvalidValue = errors.New("InvalidValue")
)

type Message struct {
	From string
	Text string
}

type SettingDB struct {
	Id                      int
	ReplacementWordPair     string
	IgnoreWords             string
	Language                string
	LanguageDetectorEnabled bool
	UserBanList             string
	ChannelsToListen        string
	Volume                  float64
}

func (s *SettingDB) SetReplacementWordPair(filter container.Map) {
	filterMap := filter.Range()
	var result strings.Builder
	for key, value := range filterMap {
		result.WriteString(key)
		result.WriteString(separator.Pair)
		result.WriteString(value)
		result.WriteString(separator.Item)
	}
	s.ReplacementWordPair = stringss.DeleteLast(result.String())
}

func (s *SettingDB) SetIgnoreWords(words []string) {
	s.IgnoreWords = stringss.FromArray(words)
}

func (s *SettingDB) SetUserBanList(users []string) {
	s.UserBanList = stringss.FromArray(users)
}

func (s *SettingDB) SetChannelsToListen(channels []string) {
	s.ChannelsToListen = stringss.FromArray(channels)
}
