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
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS settings\n(\n    id   SERIAL NOT NULL PRIMARY KEY,\n    data jsonb\n);")
	if err != nil {
		return nil, err
	}
	return &SettingRepoImpl{
		db: db,
	}, nil
}

func (r *SettingRepoImpl) GetSettings() (*model.SettingDB, error) {
	row, err := r.db.Query("SELECT data FROM settings WHERE id = 1")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Err() != nil {
		return nil, err
	}
	var result model.SettingDB
	var str string
	if row.Next() {
		err = row.Scan(&str)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(str), &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, err
}

func (r *SettingRepoImpl) SaveSettings(settings *model.SettingDB) error {
	const query = `INSERT OR REPLACE INTO settings (id, data) VALUES (1, $1);`
	data, err := json.Marshal(*settings)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, string(data))
	return err
}
