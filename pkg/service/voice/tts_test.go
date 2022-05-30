package voice

import (
	langdetection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	detectionService := langdetection.NewLinguaDetectionService(langdetection.DefaultLanguages)

	service := NewHtgoTtsService("", moderation.NewFilterDefault("", ""), 1, nil, true, detectionService)
	tests := []struct {
		input    string
		language string
	}{
		{"Each language is assigned a two-letter!", "english"},
		{"每种语言分配一个两个字母!", "chinese"},
		{"Jeder Sprache wird ein Zweibuchstabe zugewiesen!", "german"},
		{"A cada idioma se le asigna una letra de dos letras!", "spanish"},
		{"Chaque langue se voit attribuer une lettre à deux lettres!", "french"},
		{"Каждому языку присваивается двухбуквенный символ!", "русский"},
	}
	for _, test := range tests {
		assert.NoError(t, service.Speak(test.input))
	}
}
