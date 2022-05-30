package repo

import (
	"encoding/json"
	"github.com/col3name/tts/pkg/model"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"net"
	"time"
)

type Config struct {
	DbUser         string
	DbPassword     string
	DbAddress      string
	DbName         string
	MaxConnections int
	AcquireTimeout int
}

func GetConnector(config Config) (pgx.ConnPoolConfig, error) {
	databaseUri := "postgres://" + config.DbUser + ":" + config.DbPassword + "@" + config.DbAddress + "/" + config.DbName
	pgxConnConfig, err := pgx.ParseURI(databaseUri)
	if err != nil {
		return pgx.ConnPoolConfig{}, errors.Wrap(err, "failed to parse database URI from environment variable")
	}
	pgxConnConfig.Dial = (&net.Dialer{Timeout: 10 * time.Second, KeepAlive: 5 * time.Minute}).Dial
	pgxConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
	}
	pgxConnConfig.PreferSimpleProtocol = true

	return pgx.ConnPoolConfig{
		ConnConfig:     pgxConnConfig,
		MaxConnections: config.MaxConnections,
		AcquireTimeout: time.Duration(config.AcquireTimeout) * time.Second,
	}, nil
}

func NewConnectionPool(config pgx.ConnPoolConfig) (*pgx.ConnPool, error) {
	return pgx.NewConnPool(config)
}

type SettingRepo interface {
	GetSettings() (*model.SettingDB, error)
	SaveSettings(settings *model.SettingDB) error
}

type SettingRepoImpl struct {
	connPool *pgx.ConnPool
}

func NewSettingRepo(connPool *pgx.ConnPool) *SettingRepoImpl {
	return &SettingRepoImpl{
		connPool: connPool,
	}
}

func (r *SettingRepoImpl) GetSettings() (*model.SettingDB, error) {
	row, err := r.connPool.Query("SELECT data FROM settings WHERE id = 1")
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
	sql := `INSERT INTO settings (id, data) VALUES (1, $1)
	ON CONFLICT (id) DO UPDATE
		SET data = excluded.data`
	data, err := json.Marshal(*settings)
	if err != nil {
		return err
	}
	_, err = r.connPool.Exec(sql, string(data))
	return err

}
