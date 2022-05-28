package moderation

import (
	"net/url"
	"strings"
)

type Message struct {
	From string
	Text string
}

type Filter interface {
	Moderate(text Message) string
	SetFilterMap(filterMap FilterMap)
}

type BaseFilterDecorator struct {
	filter Filter
}

func (f *BaseFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *BaseFilterDecorator) Moderate(text Message) string {
	return f.filter.Moderate(text)
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

type UserFilterDecorator struct {
	BaseFilterDecorator
	users map[string]struct{}
}

func NewUserFilterDecorator(filter Filter, users []string) *UserFilterDecorator {
	u := map[string]struct{}{}
	for _, user := range users {
		if len(user) > 3 {
			u[user] = struct{}{}
		}
	}
	decorator := UserFilterDecorator{
		BaseFilterDecorator: BaseFilterDecorator{
			filter: filter,
		},
		users: u,
	}
	return &decorator
}

func (f *UserFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *UserFilterDecorator) Moderate(msg Message) string {
	_, ok := f.users[msg.From]
	if ok {
		return ""
	}
	return f.filter.Moderate(msg)
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

func (f *UrlFilterDecorator) Moderate(msg Message) string {
	split := strings.Split(msg.Text, " ")
	var result string
	for _, word := range split {
		if !f.isValidUrl(word) && !f.isContainsTopLevelDomain(word) {
			result += word + " "
		}
	}
	msg.Text = result
	return f.filter.Moderate(msg)
}

func (f *UrlFilterDecorator) isValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (f *UrlFilterDecorator) isContainsTopLevelDomain(str string) bool {
	for _, item := range mostPopularTLD {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
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

func (f *FilterDefault) Moderate(msg Message) string {
	words := strings.Split(msg.Text, " ")
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
