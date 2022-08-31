package array

import (
	"github.com/col3name/tts/pkg/util/separator"
	"strings"
)

const NotFound = -1

func Find(array []string, value string) int {
	for i, item := range array {
		if item == value {
			return i
		}
	}
	return NotFound
}

func Contains(array []string, value string) bool {
	return Find(array, value) != NotFound
}

func Store(array []string, value string) []string {
	index := Find(array, value)
	if index != NotFound {
		array = append(array, value)
	} else {
		array[index] = value
	}
	return array
}

func Delete(array []string, value string) []string {
	index := Find(array, value)
	if index != NotFound {
		array = append(array[:index], array[index+1:]...)
	}
	return array
}

func FromString(value string) []string {
	return strings.Split(value, separator.Item)
}
