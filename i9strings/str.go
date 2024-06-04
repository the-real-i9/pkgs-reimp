package i9strings

import (
	"unicode"
)

func Clone(s string) string {
	return s[:]
}

/*
The rule of string comparison is that, when comparing two characters at equal positions, the first comparison to be either 1 or -1 becomes the result of the whole comparison.

Think of a character string as having "place values" like the mathematical number string; 1 comes before 12, just as "a" comes before "ab". We can say that "aa" is the next alphabet after "z", just as 11 is the next number after 10.

That said, I choose to stick with language constructs.
*/
func Compare(a, b string) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func Contains(s, substr string) bool {
	sLen := len(s)
	substrLen := len(substr)

	lim := sLen - (substrLen - 1)

	for i := 0; i < lim; i++ {
		nextSubstrLenChars := s[i : i+substrLen] // for instance: next2Chars, provided substrLen == 2

		if nextSubstrLenChars == substr {
			return true
		}
	}

	return false
}

// lcs - "less characters string", mcs = "more characters string"
//
// Find the string with less characters (lcs) and more charaters (mcs) between s and chars.
//
// Loop through the lcs and key each character in the contains map.
//
// Loop through the mcs and check if the character exists in the contains map. If true, return true. Else, the loop runs to completion, in which case it returns false.
//
// Time Complexity - O(n)
func ContainsAny(s, chars string) bool {
	var (
		contains = make(map[rune]bool)
		lcs, mcs string
	)

	if len(s) < len(chars) {
		lcs, mcs = s, chars
	} else {
		lcs, mcs = chars, s
	}

	for _, c := range lcs {
		contains[c] = true
	}

	for _, c := range mcs {
		if contains[c] {
			return true
		}
	}

	return false
}

func ContainsFunc(s string, f func(rune) bool) bool {
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

func FieldsFunc(s string, f func(rune) bool) []string {

	strSlice := []string{}

	currStr := ""

	for i, char := range s {
		funcSat := f(char)
		isLastChar := i == len(s)-1

		if !funcSat {
			currStr += string(char)
		}

		if funcSat || isLastChar {
			if len(currStr) > 0 {
				strSlice = append(strSlice, currStr)
				currStr = ""
			}
		}
	}

	return strSlice
}

func HasPrefix(s, prefix string) bool {
	prefixEnd := len(prefix)

	return s[0:prefixEnd] == prefix
}

func HasSuffix(s, suffix string) bool {
	suffixStart := len(s) - len(suffix)

	return s[suffixStart:] == suffix
}

func Index(s, substr string) int {
	sLen := len(s)
	substrLen := len(substr)

	lim := sLen - (substrLen - 1)

	for i := 0; i < lim; i++ {
		nextSubstrLenChars := s[i : i+substrLen] // e.g. next2Chars, provided substrLen == 2

		if nextSubstrLenChars == substr {
			return i
		}
	}

	return -1
}

func IndexAny(s, chars string) int {
	var (
		exists = make(map[rune]bool)
	)

	for _, c := range chars {
		exists[c] = true
	}

	for i, c := range s {
		if exists[c] {
			return i
		}
	}

	return -1
}

func IndexByte(s string, c byte) int {
	for i, ch := range s {
		if c == byte(ch) {
			return i
		}
	}

	return -1
}

func IndexFunc(s string, f func(rune) bool) int {
	for i, c := range s {
		if f(c) {
			return i
		}
	}

	return -1
}

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}

	return -1
}

func Join(elems []string, sep string) string {
	res := ""

	for i, str := range elems {
		res += str

		isLastStr := i == len(elems)-1

		if !isLastStr {
			res += sep
		}
	}

	return res
}
