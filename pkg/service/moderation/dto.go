package moderation

import (
	"github.com/col3name/tts/pkg/util"
	"github.com/col3name/tts/pkg/util/container"
	"github.com/col3name/tts/pkg/util/separator"
	"strings"
)

type Setting struct {
	Id                      int
	ReplacementWordPair     container.Map
	IgnoreWords             []string
	Language                string
	LanguageDetectorEnabled bool
	UserBanList             []string
	ChannelsToListen        []string
	Volume                  int
}

func (s *Setting) SetIgnoreWords(value string) {
	s.IgnoreWords = strings.Split(value, separator.Item)
}

func (s *Setting) StoreIgnoreWord(word string) {
	s.IgnoreWords = util.ArrayStore(s.IgnoreWords, word)
}

func (s *Setting) DeleteIgnoreWord(word string) {
	s.IgnoreWords = util.ArrayDelete(s.IgnoreWords, word)
}

func (s *Setting) SetUserBanList(value string) {
	s.UserBanList = strings.Split(value, separator.Item)
}

func (s *Setting) StoreUserBanList(user string) {
	s.UserBanList = util.ArrayStore(s.UserBanList, user)
}

func (s *Setting) DeleteUserBanList(user string) {
	s.UserBanList = util.ArrayDelete(s.UserBanList, user)
}

func (s *Setting) SetChannelsToListen(value string) {
	s.ChannelsToListen = strings.Split(value, separator.Item)
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

func (s *Setting) SetReplacementWordPair(wordPair string) {
	builder := NewFilterMapBuilder()
	s.ReplacementWordPair = builder.Build(wordPair, separator.EmptyCharacter)
}
