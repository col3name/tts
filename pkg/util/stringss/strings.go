package stringss

import (
	"github.com/col3name/tts/pkg/util/separator"
	"strings"
)

func DeleteLast(value string) string {
	return value[:len(value)-1]
}

func FromArray(array []string) string {
	var result strings.Builder
	for _, item := range array {
		result.WriteString(item)
		result.WriteString(separator.Item)
	}
	return DeleteLast(result.String())
}

func Empty(value string) bool {
	return len(value) == 0
}
