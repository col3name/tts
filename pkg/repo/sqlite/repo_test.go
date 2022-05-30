package sqlite

import (
	"database/sql"
	"github.com/col3name/tts/pkg/repo/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	db, err := sql.Open("sqlite3", "./data.db")
	assert.NoError(t, err)
	repo, err := NewSettingRepoImpl(db)
	assert.NoError(t, err)
	common.MakeTest(t, repo)
}
