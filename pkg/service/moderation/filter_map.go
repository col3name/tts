package moderation

import (
	"github.com/col3name/tts/pkg/util"
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
	var result string
	for key, value := range m.data {
		result += key + ":" + value + ","
	}
	return result[:len(result)-1]
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

type FilterMapBuilder interface {
	build(value string) FilterMap
}

type FilterMapBuilderImpl struct{}

func (b *FilterMapBuilderImpl) Build(value string, ignoreString string) *FilterMap {
	pairs := util.StrEnumerationToArray(value)
	result := NewFilterMap()
	for _, pair := range pairs {
		splitPair := strings.Split(pair, ":")
		if len(splitPair) != 2 {
			return nil
		}
		result.Set(strings.ToLower(splitPair[0]), strings.ToLower(splitPair[1]))
	}

	split := util.StrEnumerationToArray(ignoreString)
	for _, ignore := range split {
		result.Set(strings.ToLower(ignore), "")
	}

	return result
}

var DefaultFilterMap = FilterMap{
	data: bannedWord,
}
