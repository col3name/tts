package moderation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate("shit")
	assert.Empty(t, result)
}

func TestName2(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate("wtf")
	assert.Empty(t, result)
}
