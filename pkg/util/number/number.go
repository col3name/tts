package number

import "strconv"

func InRange(value, min, max float64) bool {
	return value > min && value < max
}

func ParseFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}
