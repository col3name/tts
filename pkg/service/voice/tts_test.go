package voice

import (
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	service := NewHtgoTtsService("", moderation.NewFilterDefault("", ""), 1)
	assert.NoError(t, service.Speak("text"))
}
