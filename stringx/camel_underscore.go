package stringx

import (
	"bytes"
	"strings"
	"unicode"
)

var commonAbbr = []string{"ID", "API", "GUID", "IP", "SN"}
var titleReplace *strings.Replacer
var upReplace *strings.Replacer

func init() {
	var lowerForReplacer []string
	var upForReplacer []string
	for _, abbr := range commonAbbr {
		title := strings.Title(strings.ToLower(abbr))
		lowerForReplacer = append(lowerForReplacer, abbr, title)
		upForReplacer = append(upForReplacer, title, abbr)
	}
	titleReplace = strings.NewReplacer(lowerForReplacer...)
	upReplace = strings.NewReplacer(upForReplacer...)
}

// 驼峰式写法转为下划线写法
func UnderscoreName(name string) string {
	name = titleReplace.Replace(name)
	buffer := bytes.Buffer{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteByte('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = upReplace.Replace(strings.Title(name))
	return strings.Replace(name, " ", "", -1)
}

// 下划线首字母小写的驼峰
func LowerCamelName(name string) string {
	var sb strings.Builder
	for i, item := range strings.Split(name, "_") {
		if i == 0 {
			sb.WriteString(item)
		} else {
			sb.WriteString(upReplace.Replace(strings.Title(item)))
		}
	}
	return sb.String()
}
