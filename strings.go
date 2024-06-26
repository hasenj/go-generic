package generic

import "strings"
import "reflect"

func JoinToString[T any](list []T, sep string, stringer func(T) string) string {
	var out strings.Builder
	var listV = reflect.ValueOf(list)
	var length = listV.Len()
	var lastIndex = length - 1
	for index := 0; index < length; index++ {
		out.WriteString(stringer(list[index]))
		if index < lastIndex {
			out.WriteString(sep)
		}
	}
	return out.String()
}

func TrimSpace(s *string) {
	*s = strings.TrimSpace(*s)
}

// Generated by ChatGPT
// Extracts the shared prefix from two strings
func SharedPrefix(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	i := 0
	for i < len1 && i < len2 && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}

// Generated by ChatGPT
// Extracts the shared suffix from two strings
func SharedSuffix(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	i := 0
	for i < len1 && i < len2 && str1[len1-i-1] == str2[len2-i-1] {
		i++
	}
	return str1[len1-i:]
}
