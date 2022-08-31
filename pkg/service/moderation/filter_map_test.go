package moderation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultMapFilterBuilder(t *testing.T) {
	builder := NewFilterMapBuilder()
	result := builder.Build(`bad:nice,sad:happy`, "")
	get, ok := result.Get("bad")
	assert.True(t, ok)
	assert.Equal(t, get, "nice")
	get, ok = result.Get("sad")
	assert.True(t, ok)
	assert.Equal(t, get, "happy")
}

func TestName12(t *testing.T) {
	var result string
	for key := range bannedWord {
		result += key + ","
	}
	fmt.Println(result)
}
