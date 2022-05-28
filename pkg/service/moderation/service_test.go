package moderation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate(Message{From: "shit", Text: "shit"})
	assert.Equal(t, " ", result)
}

func TestName2(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate(Message{From: "wtf", Text: "wtf"})
	assert.Equal(t, " ", result)
}

func TestName3(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	type testData struct {
		input    Message
		expected string
	}

	tests := []testData{
		{expected: "http:::/not.valid/a//a??a?b=&&c#hi  ", input: Message{From: "1", Text: "http:::/not.valid/a//a??a?b=&&c#hi"}},
		{expected: " ", input: Message{From: "1", Text: "http//google.com"}},
		{expected: " ", input: Message{From: "1", Text: "google.com"}},
		{expected: " hello  ", input: Message{From: "1", Text: "wtf google.com hello"}},
		{expected: "/foo/bar  ", input: Message{From: "1", Text: "/foo/bar"}},
		{expected: "http://  ", input: Message{From: "1", Text: "http://"}},
		{expected: " message send by me  ", input: Message{From: "1", Text: " message send by me"}},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, urlFilter.Moderate(test.input))
	}
}

func TestName4(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	result := urlFilter.Moderate(Message{From: "", Text: "Adjust position, velocity, accel?"})
	assert.Equal(t, "Adjust position, velocity, accel?  ", result)
	userFilter := NewUserFilterDecorator(urlFilter, []string{"spin"})
	result = userFilter.Moderate(Message{From: "spin", Text: "Adjust position, velocity, accel?"})
	assert.Equal(t, "", result)
}
