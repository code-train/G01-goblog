package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString method converts
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToInt method
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// Uint64ToString method
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
