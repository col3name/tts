package common

import (
	"fmt"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func MakeTest(t *testing.T, repo repo.SettingRepo) {
	filterMap := moderation.DefaultFilterMap.Range()
	ignoreWords := ""
	for key := range filterMap {
		ignoreWords += key + ","
	}
	ignoreWords = ignoreWords[:len(ignoreWords)-1]

	settingId := 1
	in := model.SettingDB{
		Id:                      settingId,
		ReplacementWordPair:     "bad:nice",
		IgnoreWords:             ignoreWords,
		Language:                "en",
		LanguageDetectorEnabled: false,
		UserBanList:             "",
		ChannelsToListen:        "spiinlock",
		Volume:                  1,
	}
	err := repo.SaveSettings(&in)
	assert.NoError(t, err)
	out, err := repo.GetSettings()
	assert.NoError(t, err)
	reflect.DeepEqual(in, out)
	fmt.Println(out)
	setting := model.Setting{}
	setting.Id = out.Id
	setting.Volume = out.Volume
	setting.Language = out.Language
	setting.LanguageDetectorEnabled = out.LanguageDetectorEnabled
	setting.SetReplacementWordPair(out.ReplacementWordPair)
	setting.SetIgnoreWords(out.IgnoreWords)
	setting.SetUserBanList(out.UserBanList)
	setting.SetChannelsToListen(out.ChannelsToListen)
	fmt.Println(setting)
}
