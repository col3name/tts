package moderation

import (
	"net/url"
	"strings"
)

type Filter interface {
	Moderate(text string) string
	SetFilterMap(filterMap FilterMap)
}

type FilterMap struct {
	data map[string]string
}

func NewFilterMap() *FilterMap {
	f := new(FilterMap)
	f.data = map[string]string{}
	return f
}

func (m *FilterMap) get(from string) (string, bool) {
	s, ok := m.data[from]
	return s, ok
}

func (m *FilterMap) push(from, to string) {
	m.data[from] = to
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
	pairs := strings.Split(value, ",")
	result := NewFilterMap()
	for _, pair := range pairs {
		splitPair := strings.Split(pair, ":")
		if len(splitPair) != 2 {
			return nil
		}
		result.push(strings.ToLower(splitPair[0]), strings.ToLower(splitPair[1]))
	}

	split := strings.Split(ignoreString, ",")
	for _, ignore := range split {
		result.push(strings.ToLower(ignore), "")
	}

	return result
}

var DefaultFilterMap = FilterMap{
	data: bannedWord,
}

type BaseFilterDecorator struct {
	filter Filter
}

func (f *BaseFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *BaseFilterDecorator) Moderate(text string) string {
	return f.filter.Moderate(text)
}

type UrlFilterDecorator struct {
	BaseFilterDecorator
}

func NewUrlFilterDecorator(filter Filter) *UrlFilterDecorator {
	return &UrlFilterDecorator{
		BaseFilterDecorator{
			filter: filter,
		},
	}
}

func (f *UrlFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *UrlFilterDecorator) isValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

var mostPopularTLD = []string{
	".com",
	".net",
	".org",
	".de",
	".icu",
	".uk",
	".ru",
	".info",
	".top",
	".xyz",
	".tk",
	".cn",
	".ga",
	".cf",
	".nl",
}

func (f *UrlFilterDecorator) isContainsTopLevelDomain(str string) bool {
	for _, item := range mostPopularTLD {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
}

func (f *UrlFilterDecorator) Moderate(text string) string {
	split := strings.Split(text, " ")
	var result string
	for _, word := range split {
		if !f.isValidUrl(word) && !f.isContainsTopLevelDomain(word) {
			result += word + " "
		}
	}
	return f.filter.Moderate(result)
}

type FilterDefault struct {
	filterMap FilterMap
}

func NewFilterDefault(moderationPair, ignoreString string) *FilterDefault {
	f := new(FilterDefault)
	if len(moderationPair) > 0 || len(ignoreString) > 0 {
		builder := FilterMapBuilderImpl{}
		f.filterMap = *builder.Build(moderationPair, ignoreString)
	} else {
		f.filterMap = DefaultFilterMap
	}
	return f
}

func (f *FilterDefault) SetFilterMap(filterMap FilterMap) {
	f.filterMap = filterMap
}

func (f *FilterDefault) Moderate(text string) string {
	words := strings.Split(text, " ")
	var result string
	var val string
	var ok bool

	for _, word := range words {
		val, ok = f.filterMap.get(strings.ToLower(word))
		if ok {
			result += val
		} else {
			result += word
		}
		result += " "
	}

	return result
}
