package utils

import "strconv"

func FromStringToInt64(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
