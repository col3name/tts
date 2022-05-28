package moderation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate("shit")
	assert.Equal(t, " ", result)
}

func TestName2(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	result := filterDefault.Moderate("wtf")
	assert.Equal(t, " ", result)
}

func TestName3(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	type testData struct {
		input    string
		expected string
	}

	tests := []testData{
		{expected: "http:::/not.valid/a//a??a?b=&&c#hi  ", input: "http:::/not.valid/a//a??a?b=&&c#hi"},
		{expected: " ", input: "http//google.com"},
		{expected: " ", input: "google.com"},
		{expected: " hello  ", input: "wtf google.com hello"},
		{expected: "/foo/bar  ", input: "/foo/bar"},
		{expected: "http://  ", input: "http://"},
		{expected: " message send by me  ", input: "wtf message send by me"},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, urlFilter.Moderate(test.input))
	}
}

func TestName4(t *testing.T) {
	filterDefault := NewFilterDefault("", "")
	urlFilter := NewUrlFilterDecorator(filterDefault)
	result := urlFilter.Moderate("Adjust position, velocity, accel?")
	assert.Equal(t, "Adjust position, velocity, accel?  ", result)
}
