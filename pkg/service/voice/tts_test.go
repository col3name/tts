package voice

import (
	"github.com/col3name/tts/pkg/service/moderation"
	"testing"
)

func TestName(t *testing.T) {
	service := NewHtgoTtsService("", moderation.NewFilterDefault("", ""))
	service.Speak("text")
}
