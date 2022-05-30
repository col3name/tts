package util

import "strings"

func StrEnumerationToArray(str string) []string {
	return strings.Split(str, ",")
}

func FindInArray(arr []string, val string) int {
	for i, item := range arr {
		if item == val {
			return i
		}
	}
	return -1
}

func ArrayDelete(arr []string, val string) []string {
	index := FindInArray(arr, val)
	if index != -1 {
		arr = append(arr[:index], arr[index+1:]...)
	}
	return arr
}

func ArrayStore(arr []string, val string) []string {
	index := FindInArray(arr, val)
	if index != -1 {
		arr = append(arr, val)
	} else {
		arr[index] = val
	}
	return arr
}
