package repo

import "github.com/col3name/tts/pkg/model"

type SettingRepo interface {
	GetSettings() (*model.SettingDB, error)
	SaveSettings(settings *model.SettingDB) error
}
