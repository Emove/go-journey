package string

import (
	"strings"
	"unicode"
)

func UpperCamelCase(str string) string {
	if unicode.IsUpper(rune(str[0])) {
		return str
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func LowerCamelCae(str string) string {
	if unicode.IsLower(rune(str[0])) {
		return str
	}
	return strings.ToLower(str[:1]) + str[1:]
}

func CamelToSnake(value string) string {
	if len(value) == 0 {
		return ""
	}
	var result []rune
	for _, v := range value {
		if unicode.IsUpper(v) {
			result = append(result, '_')

		}
		result = append(result, unicode.ToLower(v))
	}
	if result[0] == '_' {
		result = result[1:]
	}
	return string(result)
}

func Trim(s string, ss string) string {
	return strings.Trim(s, ss)
}
