package utils

import (
	"strings"
)

// ToSnakeCase converts in string to snake_case
func ToSnakeCase(in string) string {
	var isPrevLower bool
	var buf strings.Builder

	for _, r := range in {
		if 'A' <= r && r <= 'Z' && isPrevLower {
			buf.WriteByte('_')
			buf.WriteString(strings.ToLower(string(r)))
			isPrevLower = false
			continue
		}

		isPrevLower = 'a' <= r && r <= 'z'
		buf.WriteString(strings.ToLower(string(r)))
	}

	return buf.String()
}

// FirstLower makes first symbol lowercase in the string
func FirstLower(s string) string {
	if len(s) < 1 {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}

	return strings.ToLower(s[:1]) + s[1:]
}
