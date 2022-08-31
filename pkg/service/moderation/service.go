package moderation

import (
	"github.com/col3name/tts/pkg/model"
	"net/url"
	"strings"
)

type Filter interface {
	Moderate(text model.Message) string
	SetFilterMap(filterMap FilterMap)
}

type BaseFilterDecorator struct {
	filter Filter
}

func (f *BaseFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *BaseFilterDecorator) Moderate(text model.Message) string {
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

func NewDefaultFilter(moderationPair, ignoreString string, users []string) *UserFilterDecorator {
	filterDefault := NewFilterDefault(moderationPair, ignoreString)
	urlFilter := NewUrlFilterDecorator(filterDefault)
	return NewUserFilterDecorator(urlFilter, users)
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

func (f *UserFilterDecorator) Moderate(msg model.Message) string {
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

func (f *UrlFilterDecorator) Moderate(msg model.Message) string {
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

const defaultIgnore = "shit,slut,spunk,whore,fuck,nigger,sex,pussy,queer,sh1t,wank,wtf,anal,bitch,poop,tosser,vagina,balls,Goddamn,muff,clitoris,knobend,knob end,ballsack,bastard,bum,penis,arse,dick,f u c k,God damn,pube,anus,cunt,fellate,feck,felching,lmao,nigga,omg,bollok,dildo,fag,homo,turd,bugger,buttplug,dyke,bollock,flange,blowjob,boob,crap,labia,scrotum,s hit,smegma,ass,biatch,coon,lmfao,boner,fudge packer,jizz,hell,jerk,piss,tit,twat,bloody,butt,damn,blow job,cock,fellatio,fudgepacker,prick"

func NewFilterDefault(moderationPair, ignoreString string) *FilterDefault {
	f := new(FilterDefault)
	if len(moderationPair) > 0 || len(ignoreString) > 0 {
		builder := FilterMapBuilderImpl{}
		f.filterMap = *builder.Build(moderationPair, ignoreString+defaultIgnore)
	} else {
		f.filterMap = DefaultFilterMap
	}
	return f
}

func (f *FilterDefault) SetFilterMap(filterMap FilterMap) {
	f.filterMap = filterMap
}

func (f *FilterDefault) Moderate(msg model.Message) string {
	words := strings.Split(msg.Text, " ")
	var result string
	var val string
	var ok bool

	for _, word := range words {
		val, ok = f.filterMap.Get(strings.ToLower(word))
		if ok {
			result += val
		} else {
			result += word
		}
		result += " "
	}

	return result
}
