package helpers

import "strings"

func LimitString(s string, limit int) string {
	if len(s) > limit {
		return s[:limit-3] + "..."
	}

	return s
}

func RemoveSpecialChars(s string) string {
	strings.Replace("\n", s, "", -1)
	strings.Replace("\t", s, "", -1)
	strings.Replace("\r", s, "", -1)

	return s
}
