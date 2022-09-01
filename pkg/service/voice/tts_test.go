package voice

import (
	"github.com/col3name/tts/pkg/model"
	langdetection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestVoice(t *testing.T) {
	skipCI(t)
	detectionService := langdetection.NewLinguaDetectionService(langdetection.DefaultLanguages)

	service := NewGoTtsService("", moderation.NewBaseFilter("", ""), 1, nil, true, detectionService)
	tests := []struct {
		name     string
		input    string
		language string
	}{
		{name: "1", input: "Each language is assigned a two-letter!", language: "english"},
		{name: "2", input: "每种语言分配一个两个字母!", language: "chinese"},
		{name: "3", input: "Jeder Sprache wird ein Zweibuchstabe zugewiesen!", language: "german"},
		{name: "4", input: "A cada idioma se le asigna una letra de dos letras!", language: "spanish"},
		{name: "5", input: "Chaque langue se voit attribuer une lettre à deux lettres!", language: "french"},
		{name: "6", input: "Каждому языку присваивается двухбуквенный символ!", language: "русский"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.NoError(t, service.Speak(model.Message{
				From: "",
				Text: test.input,
			}))
		})
	}
}
