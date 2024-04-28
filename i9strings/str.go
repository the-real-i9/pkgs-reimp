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

func Contains(s, sep string) bool {
	sLen := len(s)
	sepLen := len(sep)

	lim := sLen - (sepLen - 1)

	for i := 0; i < lim; i++ {
		if s[i:i+sepLen] == sep {
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
//
// Time Complexity - O(n)
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

func ConstainsFunc(s string, f func(rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}

	return false
}

func ConstainsRune(s string, r rune) bool {
	for _, tr := range s {
		if r == tr {
			return true
		}
	}

	return false
}

func Count(s, sep string) int {
	if sep == "" {
		return len(s) + 1
	}

	sLen := len(s)
	sepLen := len(sep)

	lim := sLen - (sepLen - 1)

	cnt := 0

	for i := 0; i < lim; i++ {
		if s[i:i+sepLen] == sep {
			cnt++
		}
	}

	return cnt
}

func Cut(s, sep string) (before, after string, found bool) {
	if sep == "" {
		return "", s, true
	}

	sLen := len(s)
	sepLen := len(sep)

	lim := sLen - (sepLen - 1)

	for i := 0; i < lim; i++ {
		if s[i:i+sepLen] == sep {
			return s[0:i], s[i+sepLen:], true
		}
	}

	return s, "", false
}

func CutPrefix(s, prefix string) (after string, found bool) {
	if prefix == "" {
		return s, true
	}

	prefixLen := len(prefix)

	if s[0:prefixLen] == prefix {
		return s[prefixLen:], true
	}

	return s, false
}

func CutSuffix(s, suffix string) (before string, found bool) {
	if suffix == "" {
		return s, true
	}

	fromTo := len(s) - len(suffix)

	if s[fromTo:] == suffix {
		return s[0:fromTo], true
	}

	return s, false
}
