package moderation

import (
	"github.com/col3name/tts/pkg/model"
	"net/url"
	"strings"
)

type Filter interface {
	SetFilterMap(filterMap FilterMap)
	Moderate(text *model.Message) (string, error)
}

type BaseFilterDecorator struct {
	filter Filter
}

func (f *BaseFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *BaseFilterDecorator) Moderate(text *model.Message) (string, error) {
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

func NewDefaultFilter(moderationPair, ignoreString string, users []string) Filter {
	baseFilter := NewBaseFilter(moderationPair, ignoreString)
	urlFilter := NewUrlFilterDecorator(baseFilter)
	userFilter := NewUserFilterDecorator(urlFilter, users)
	return NewTrimFilterDecorator(userFilter)
}

type UserFilterDecorator struct {
	BaseFilterDecorator
	users map[string]struct{}
}

const MinUsernameLength = 3

func NewUserFilterDecorator(filter Filter, users []string) *UserFilterDecorator {
	usersMap := map[string]struct{}{}
	for _, user := range users {
		if len(user) > MinUsernameLength {
			usersMap[user] = struct{}{}
		}
	}
	decorator := UserFilterDecorator{
		BaseFilterDecorator: BaseFilterDecorator{
			filter: filter,
		},
		users: usersMap,
	}
	return &decorator
}

func (f *UserFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *UserFilterDecorator) Moderate(message *model.Message) (string, error) {
	_, ok := f.users[message.From]
	if ok {
		return "", model.ErrorInvalidValue
	}
	return f.filter.Moderate(message)
}

type TrimFilterDecorator struct {
	BaseFilterDecorator
}

func NewTrimFilterDecorator(filter Filter) *TrimFilterDecorator {
	return &TrimFilterDecorator{
		BaseFilterDecorator{
			filter: filter,
		},
	}
}

func (f *TrimFilterDecorator) SetFilterMap(filterMap FilterMap) {
	f.filter.SetFilterMap(filterMap)
}

func (f *TrimFilterDecorator) Moderate(message *model.Message) (string, error) {
	text := message.Text
	text = strings.Trim(text, model.Space)
	fromLen := len(message.From)
	if fromLen > len(text) && len(text) == 0 {
		return "", model.ErrorInvalidValue
	}
	check := text[fromLen:]
	if strings.HasSuffix(check, "say    !") {
		return "", model.ErrorInvalidValue
	}
	message.Text = text
	return f.filter.Moderate(message)
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

func (f *UrlFilterDecorator) Moderate(message *model.Message) (string, error) {

	words := strings.Split(message.Text, model.Space)
	var text strings.Builder
	for _, word := range words {
		if f.isValidWord(word) {
			text.WriteString(word + model.Space)
		}
	}
	message.Text = text.String()
	return f.filter.Moderate(message)
}

func (f *UrlFilterDecorator) isValidWord(value string) bool {
	return !f.isValidUrl(value) && !f.isContainsTopLevelDomain(value)
}

func (f *UrlFilterDecorator) isValidUrl(value string) bool {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return false
	}

	u, err := url.Parse(value)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (f *UrlFilterDecorator) isContainsTopLevelDomain(value string) bool {
	for _, item := range mostPopularTLD {
		if strings.Contains(value, item) {
			return true
		}
	}
	return false
}

type BaseFilter struct {
	filterMap FilterMap
}

const defaultIgnore = "shit,slut,spunk,whore,fuck,nigger,sex,pussy,queer,sh1t,wank,wtf,anal,bitch,poop,tosser,vagina,balls,Goddamn,muff,clitoris,knobend,knob end,ballsack,bastard,bum,penis,arse,dick,f u c k,God damn,pube,anus,cunt,fellate,feck,felching,lmao,nigga,omg,bollok,dildo,fag,homo,turd,bugger,buttplug,dyke,bollock,flange,blowjob,boob,crap,labia,scrotum,s hit,smegma,ass,biatch,coon,lmfao,boner,fudge packer,jizz,hell,jerk,piss,tit,twat,bloody,butt,damn,blow job,cock,fellatio,fudgepacker,prick"

func NewBaseFilter(moderationPair, ignoreString string) *BaseFilter {
	f := new(BaseFilter)
	if len(moderationPair) > 0 || len(ignoreString) > 0 {
		builder := NewFilterMapBuilder()
		f.filterMap = *builder.Build(moderationPair, ignoreString+defaultIgnore)
	} else {
		f.filterMap = DefaultFilterMap
	}
	return f
}

func (f *BaseFilter) SetFilterMap(filterMap FilterMap) {
	f.filterMap = filterMap
}

func (f *BaseFilter) Moderate(message *model.Message) (string, error) {
	if len(message.Text) == 0 {
		return message.Text, model.ErrorInvalidValue
	}
	words := strings.Split(message.Text, model.Space)
	var text strings.Builder
	var value string
	var ok bool

	for _, word := range words {
		value, ok = f.filterMap.Get(strings.ToLower(word))
		if ok {
			text.WriteString(value)
		} else {
			text.WriteString(word)
		}
		text.WriteString(model.Space)
	}

	return text.String(), nil
}
