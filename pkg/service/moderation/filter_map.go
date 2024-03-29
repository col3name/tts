package moderation

import (
	"errors"
	"github.com/col3name/tts/pkg/util/array"
	"github.com/col3name/tts/pkg/util/separator"
	"github.com/col3name/tts/pkg/util/stringss"
	"strings"
)

type FilterMap struct {
	data map[string]string
}

func NewFilterMap() *FilterMap {
	f := new(FilterMap)
	f.data = map[string]string{}
	return f
}

func (m *FilterMap) Get(from string) (string, bool) {
	s, ok := m.data[from]
	return s, ok
}

func (m *FilterMap) Range() map[string]string {
	return m.data
}

func (m *FilterMap) String() string {
	var result strings.Builder
	for key, value := range m.data {
		result.WriteString(key)
		result.WriteString(separator.Pair)
		result.WriteString(value)
		result.WriteString(separator.Item)
	}
	return stringss.DeleteLast(result.String())
}

func (m *FilterMap) Set(from, to string) {
	m.data[from] = to
}

func (m *FilterMap) Remove(from string) {
	delete(m.data, from)
}

var bannedWord = map[string]string{"anal": "",
	"anus":         "",
	"arse":         "",
	"ass":          "",
	"ballsack":     "",
	"balls":        "",
	"bastard":      "",
	"bitch":        "",
	"biatch":       "",
	"bloody":       "",
	"blowjob":      "",
	"blow job":     "",
	"bollock":      "",
	"bollok":       "",
	"boner":        "",
	"boob":         "",
	"bugger":       "",
	"bum":          "",
	"butt":         "",
	"buttplug":     "",
	"clitoris":     "",
	"cock":         "",
	"coon":         "",
	"crap":         "",
	"cunt":         "",
	"damn":         "",
	"dick":         "",
	"dildo":        "",
	"dyke":         "",
	"fag":          "",
	"feck":         "",
	"fellate":      "",
	"fellatio":     "",
	"felching":     "",
	"fuck":         "",
	"f u c k":      "",
	"fudgepacker":  "",
	"fudge packer": "",
	"flange":       "",
	"Goddamn":      "",
	"God damn":     "",
	"hell":         "",
	"homo":         "",
	"jerk":         "",
	"jizz":         "",
	"knobend":      "",
	"knob end":     "",
	"labia":        "",
	"lmao":         "",
	"lmfao":        "",
	"muff":         "",
	"nigger":       "",
	"nigga":        "",
	"omg":          "",
	"penis":        "",
	"piss":         "",
	"poop":         "",
	"prick":        "",
	"pube":         "",
	"pussy":        "",
	"queer":        "",
	"scrotum":      "",
	"sex":          "",
	"shit":         "",
	"s hit":        "",
	"sh1t":         "",
	"slut":         "",
	"smegma":       "",
	"spunk":        "",
	"tit":          "",
	"tosser":       "",
	"turd":         "",
	"twat":         "",
	"vagina":       "",
	"wank":         "",
	"whore":        "",
	"wtf":          ""}

var DefaultFilterMap = FilterMap{
	data: bannedWord,
}

type FilterMapBuilder interface {
	build(value string) FilterMap
}

type FilterMapBuilderImpl struct{}

func NewFilterMapBuilder() *FilterMapBuilderImpl {
	return &FilterMapBuilderImpl{}
}

var (
	ErrorInvalidValue = errors.New("InvalidValue")
)

func (b *FilterMapBuilderImpl) Build(wordPairs string, ignoreString string) *FilterMap {
	filterMap := NewFilterMap()
	err := b.fillFilterMap(filterMap, wordPairs, b.handleWordPair)
	if err != nil {
		return nil
	}
	err = b.fillFilterMap(filterMap, ignoreString, b.handleIgnoreWord)
	if err != nil {
		return nil
	}
	return filterMap
}

func (b *FilterMapBuilderImpl) fillFilterMap(filterMap *FilterMap, value string, fn func(filterMap *FilterMap, pair string) error) error {
	itemArray := array.FromString(value)

	for _, item := range itemArray {
		if err := fn(filterMap, item); err != nil {
			return ErrorInvalidValue
		}
	}

	return nil
}

func (b *FilterMapBuilderImpl) handleWordPair(filterMap *FilterMap, pair string) error {
	splitPair := strings.Split(pair, separator.Pair)
	if len(splitPair) != 2 {
		return ErrorInvalidValue
	}
	filterMap.Set(strings.ToLower(splitPair[0]), strings.ToLower(splitPair[1]))
	return nil
}

func (b *FilterMapBuilderImpl) handleIgnoreWord(filterMap *FilterMap, word string) error {
	filterMap.Set(strings.ToLower(word), "")
	return nil
}
