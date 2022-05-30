package repo

import (
	"fmt"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	connector, err := GetConnector(Config{
		DbAddress:      "localhost:5432",
		DbName:         "twitch_tts",
		DbUser:         "postgres",
		DbPassword:     "postgres",
		MaxConnections: 10,
		AcquireTimeout: 1,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	pool, err := NewConnectionPool(connector)
	assert.NoError(t, err)
	repo := SettingRepoImpl{connPool: pool}
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
	err = repo.SaveSettings(&in)
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
