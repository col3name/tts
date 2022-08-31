package setting

import (
	"database/sql"
	"github.com/col3name/tts/pkg/config"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	"github.com/col3name/tts/pkg/repo/sqlite"
	"github.com/col3name/tts/pkg/util/logger"
)

func UpdateSetting(config *config.Config, settingRepo repo.SettingRepo) *model.SettingDB {
	setting := setupSettingFromConfig(config)

	settingsFromDatabase, err := settingRepo.GetSettings()
	logger.LogFatalIfNeed(err)
	mergeSetting(setting, settingsFromDatabase)
	logger.LogFatalIfNeed(err)
	return setting
}

func SetupSettingRepo() repo.SettingRepo {
	db, err := sql.Open("sqlite3", "./data.db")
	logger.LogFatalIfNeed(err)
	settingRepo, err := sqlite.NewSettingRepoImpl(db)
	logger.LogFatalIfNeed(err)
	return settingRepo
}

func mergeSetting(setting *model.SettingDB, settingsFromDatabase *model.SettingDB) {
	if settingsFromDatabase != nil {
		setting.ReplacementWordPair = settingsFromDatabase.ReplacementWordPair
		setting.IgnoreWords = settingsFromDatabase.IgnoreWords
		setting.UserBanList = settingsFromDatabase.UserBanList
		setting.ChannelsToListen = settingsFromDatabase.ChannelsToListen
	}
}

func setupSettingFromConfig(config *config.Config) *model.SettingDB {
	setting := &model.SettingDB{
		Id:                      1,
		ReplacementWordPair:     config.ModerationWordPairs,
		IgnoreWords:             config.IgnoreWords,
		Language:                config.Language,
		LanguageDetectorEnabled: config.LangDetectorEnabled,
		ChannelsToListen:        "",
		Volume:                  1,
	}
	setting.SetUserBanList(config.UserBanList)
	return setting
}
