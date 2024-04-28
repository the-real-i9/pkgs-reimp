package i9strings

import (
	"cmp"
)

func Clone(s string) string {
	return s[:]
}

func Compare(a, b string) int {
	return cmp.Compare(a, b)
}

func Contains(s, substr string) bool {
	strLen := len(s)
	substrLen := len(substr)

	lim := strLen - (substrLen - 1)

	for i := 0; i < lim; i++ {
		if s[i:i+substrLen] == substr {
			return true
		}
	}

	return false
}
