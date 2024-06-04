package i9strings

import (
	"cmp"
	"unicode"
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

	prefixEnd := len(prefix)

	if s[0:prefixEnd] == prefix {
		return s[prefixEnd:], true
	}

	return s, false
}

func CutSuffix(s, suffix string) (before string, found bool) {
	if suffix == "" {
		return s, true
	}

	suffixStart := len(s) - len(suffix)

	if s[suffixStart:] == suffix {
		return s[0:suffixStart], true
	}

	return s, false
}

/*
Both strings must be of equal length, if not, we immediately return false.

The code point difference between same characters of opposite cases is exactly 32.

Since EqualFold tests for same characters case-insensitively, if codepoint difference of two characters is zero (same case), 32 (lowercase - uppercase), or -32 (uppercase - lowercase), the two characters are the same case-insensitively, otherwise they're not same.

If unequality is detected, we return false.
*/
func EqualFold(s, t string) bool {
	sLen := len(s)
	if sLen != len(t) {
		return false
	}

	isSameChar := func(sChar byte, tChar byte) bool {
		codePointDiff := int(sChar) - int(tChar)

		return codePointDiff == 0 || codePointDiff == 32 || codePointDiff == -32
	}

	for i := 0; i < sLen; i++ {
		if isSameChar(s[i], t[i]) {
			continue
		}
		return false
	}

	return true
}

func Fields(s string) []string {

	strSlice := []string{}

	currStr := ""

	for i, char := range s {
		isSpace := unicode.IsSpace(char)
		isLastChar := i == len(s)-1

		if !isSpace {
			currStr += string(char)
		}

		if isSpace || isLastChar {
			if len(currStr) > 0 {
				strSlice = append(strSlice, currStr)
				currStr = ""
			}
		}
	}

	return strSlice
}
