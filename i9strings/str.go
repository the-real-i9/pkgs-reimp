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
	sLen := len(s)
	substrLen := len(substr)

	lim := sLen - (substrLen - 1)

	for i := 0; i < lim; i++ {
		if s[i:i+substrLen] == substr {
			return true
		}
	}

	return false
}

// tr - "tracker map", lcs - "less characters string", mcs = "more characters string"
//
// Find the string with less characters (lcs) and more charaters (mcs) between s and chars.
//
// Loop through the string with less characters and key each rune in the tracker map.
//
// Loop through the other string and check if the rune exists within the tracker map. If true, return true. Else, the loop runs to completion, in which case it returns false.
func ContainsAny(s, chars string) bool {
	var (
		tr       = make(map[rune]bool)
		lcs, mcs string
	)

	if len(s) < len(chars) {
		lcs, mcs = s, chars
	} else {
		lcs, mcs = chars, s
	}

	for _, r := range lcs {
		tr[r] = true
	}

	for _, r := range mcs {
		if tr[r] {
			return true
		}
	}

	return false
}
