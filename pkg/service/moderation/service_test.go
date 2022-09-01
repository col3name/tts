package moderation

import (
	"github.com/col3name/tts/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	filterDefault := NewBaseFilter("", "")
	result, _ := filterDefault.Moderate(&model.Message{From: "shit", Text: "shit"})
	assert.Equal(t, " ", result)
}

func TestName2(t *testing.T) {
	filterDefault := NewBaseFilter("", "")
	result, _ := filterDefault.Moderate(&model.Message{From: "wtf", Text: "wtf"})
	assert.Equal(t, " ", result)
}

func TestName3(t *testing.T) {
	filterDefault := NewBaseFilter("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	type testData struct {
		name     string
		input    *model.Message
		expected string
	}

	tests := []testData{
		{name: "1", expected: "http:::/not.valid/a//a??a?b=&&c#hi  ", input: &model.Message{From: "1", Text: "http:::/not.valid/a//a??a?b=&&c#hi"}},
		{name: "2", expected: "", input: &model.Message{From: "1", Text: "http//google.com"}},
		{name: "3", expected: "", input: &model.Message{From: "1", Text: "google.com"}},
		{name: "4", expected: " hello  ", input: &model.Message{From: "1", Text: "wtf google.com hello"}},
		{name: "5", expected: "/foo/bar  ", input: &model.Message{From: "1", Text: "/foo/bar"}},
		{name: "6", expected: "http://  ", input: &model.Message{From: "1", Text: "http://"}},
		{name: "7", expected: " message send by me  ", input: &model.Message{From: "1", Text: " message send by me"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			text, _ := urlFilter.Moderate(test.input)
			assert.Equal(t, test.expected, text)
		})
	}
}

func TestName4(t *testing.T) {
	filterDefault := NewBaseFilter("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	result, _ := urlFilter.Moderate(&model.Message{From: "", Text: "Adjust position, velocity, accel?"})
	assert.Equal(t, "Adjust position, velocity, accel?  ", result)
	userFilter := NewUserFilterDecorator(urlFilter, []string{"spin"})
	result, _ = userFilter.Moderate(&model.Message{From: "spin", Text: "Adjust position, velocity, accel?"})
	assert.Equal(t, "", result)
}
