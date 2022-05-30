package model

import (
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/util"
	"strings"
)

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

func (s *SettingDB) SetReplacementWordPair(filter moderation.FilterMap) {
	filterMap := filter.Range()
	result := ""
	for key, value := range filterMap {
		result += key + ":" + value + ","
	}
	s.ReplacementWordPair = result[:len(result)-1]
}

func (s *SettingDB) SetIgnoreWords(words []string) {
	var result string
	for _, item := range words {
		result += item + ","
	}
	s.IgnoreWords = result[:len(result)-1]
}

func (s *SettingDB) SetUserBanList(users []string) {
	var result string
	for _, item := range users {
		result += item + ","
	}
	s.UserBanList = result[:len(result)-1]
}

func (s *SettingDB) SetChannelsToListen(list []string) {
	var result string
	for _, item := range list {
		result += item + ","
	}
	s.ChannelsToListen = result[:len(result)-1]
}

type Setting struct {
	Id                      int
	ReplacementWordPair     moderation.FilterMap
	IgnoreWords             []string
	Language                string
	LanguageDetectorEnabled bool
	UserBanList             []string
	ChannelsToListen        []string
	Volume                  int
}

func (s *Setting) SetIgnoreWords(str string) {
	s.IgnoreWords = strings.Split(str, ",")
}

func (s *Setting) StoreIgnoreWord(word string) {
	s.IgnoreWords = util.ArrayStore(s.IgnoreWords, word)
}

func (s *Setting) DeleteIgnoreWord(word string) {
	s.IgnoreWords = util.ArrayDelete(s.IgnoreWords, word)
}

func (s *Setting) SetUserBanList(str string) {
	s.UserBanList = strings.Split(str, ",")
}

func (s *Setting) StoreUserBanList(user string) {
	s.UserBanList = util.ArrayStore(s.UserBanList, user)
}

func (s *Setting) DeleteUserBanList(user string) {
	s.UserBanList = util.ArrayDelete(s.UserBanList, user)
}

func (s *Setting) SetChannelsToListen(str string) {
	s.ChannelsToListen = strings.Split(str, ",")
}

func (s *Setting) StoreChannelsToListen(user string) {
	s.ChannelsToListen = util.ArrayStore(s.ChannelsToListen, user)
}

func (s *Setting) DeleteChannelsToListen(user string) {
	s.ChannelsToListen = util.ArrayDelete(s.ChannelsToListen, user)
}

func (s *Setting) SetReplacementPair(key, value string) {
	s.ReplacementWordPair.Set(key, value)
}

func (s *Setting) RemoveReplacementPair(key string) {
	s.ReplacementWordPair.Remove(key)
}

func (s *Setting) GetReplacementPair(key string) (string, bool) {
	return s.ReplacementWordPair.Get(key)
}

func (s *Setting) SetReplacementWordPair(str string) {
	builder := moderation.FilterMapBuilderImpl{}
	s.ReplacementWordPair = *builder.Build(str, "")
}
