package lang_detection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	input    string
	expected Language
}

func TestDetermineLanguage(t *testing.T) {
	tests := []testCase{
		{" Each language is assigned a two-letter ", "en"},
		{"你好，世界", "zh"},
		{"Hallo Welt", "de"},
		{"A cada idioma se le asigna una letra de dos letras", "es"},
		{"Chaque langue se voit attribuer une lettre à deux lettres", "fr"},
		{"Каждому языку присваивается двухбуквенный", "ru"},
	}

	service := NewLinguaDetectionService(DefaultLanguages)
	for _, test := range tests {
		lang, err := service.Detect(test.input)
		fmt.Println(lang, err, test.expected)
		assert.Equal(t, test.expected, *lang)
	}
}
