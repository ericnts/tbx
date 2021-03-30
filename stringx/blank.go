package stringx

import (
	"strings"
)

// 判断字符串为空
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 判断字符串不为空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
