package sqlite

import (
	"database/sql"
	"encoding/json"
	"github.com/col3name/tts/pkg/model"
	_ "github.com/mattn/go-sqlite3"
)

type SettingRepoImpl struct {
	db *sql.DB
}

func NewSettingRepoImpl(db *sql.DB) (*SettingRepoImpl, error) {
	_, err := db.Exec(getCreateSettingTableSQL())
	if err != nil {
		return nil, err
	}
	return &SettingRepoImpl{db: db}, nil
}

func getCreateSettingTableSQL() string {
	return "CREATE TABLE IF NOT EXISTS settings\n(\n    id   SERIAL NOT NULL PRIMARY KEY,\n    data jsonb\n);"
}

func (r *SettingRepoImpl) GetSettings() (*model.SettingDB, error) {
	row, err := r.db.Query(r.getSettingSQL())
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Err() != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, err
	}

	var value string

	err = row.Scan(&value)
	if err != nil {
		return nil, err
	}

	return r.unmarshallSetting(value)
}

func (r *SettingRepoImpl) getSettingSQL() string {
	return "SELECT data FROM settings WHERE id = 1;"
}

func (r *SettingRepoImpl) SaveSettings(settings *model.SettingDB) error {
	data, err := r.marshalSetting(settings)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(r.getSaveSettingSQL(), data)
	return err
}

func (r *SettingRepoImpl) getSaveSettingSQL() string {
	return `INSERT OR REPLACE INTO settings (id, data) VALUES (1, $1);`
}

func (r *SettingRepoImpl) unmarshallSetting(value string) (*model.SettingDB, error) {
	var result model.SettingDB

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SettingRepoImpl) marshalSetting(settings *model.SettingDB) (string, error) {
	data, err := json.Marshal(*settings)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
