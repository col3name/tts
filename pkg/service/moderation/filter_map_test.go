package moderation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultMapFilterBuilder(t *testing.T) {
	builder := FilterMapBuilderImpl{}
	result := builder.Build(`bad:nice,sad:happy`, "")
	get, ok := result.get("bad")
	assert.True(t, ok)
	assert.Equal(t, get, "nice")
	get, ok = result.get("sad")
	assert.True(t, ok)
	assert.Equal(t, get, "happy")
}
