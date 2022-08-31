package model

import (
	"errors"
	"github.com/col3name/tts/pkg/util/container"
	"github.com/col3name/tts/pkg/util/separator"
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
	var textBuilder strings.Builder
	for key, value := range filterMap {
		textBuilder.WriteString(key)
		textBuilder.WriteString(separator.Pair)
		textBuilder.WriteString(value)
		textBuilder.WriteString(separator.Item)
	}
	text := textBuilder.String()
	s.ReplacementWordPair = text[:len(text)-1]
}

func (s *SettingDB) SetIgnoreWords(words []string) {
	var result string
	for _, item := range words {
		result += item + separator.Item
	}
	s.IgnoreWords = result[:len(result)-1]
}

func (s *SettingDB) SetUserBanList(users []string) {
	var result string
	for _, item := range users {
		result += item + separator.Item
	}
	s.UserBanList = result[:len(result)-1]
}

func (s *SettingDB) SetChannelsToListen(list []string) {
	var result string
	for _, item := range list {
		result += item + separator.Item
	}
	s.ChannelsToListen = result[:len(result)-1]
}
