package i9strings

import (
	"slices"
	"testing"
)

func TestCompare(t *testing.T) {
	// str1, str2, exp := "axc", "abz", 1 // in alphabhet listing, we expect to encounter "abz" before "axc"
	str1, str2, exp := "a", "aa", -1 // in alphabhet listing, we expect to encounter "a" before "aa"

	res := Compare(str1, str2)

	if res != exp {
		t.Errorf("Expected %d for EqualFold(%s, %s). Instead got %d.", exp, str1, str2, res)
	}
}

func TestEqualFold(t *testing.T) {
	// str1, str2, exp := "aBcDe", "AbCdE", true
	str1, str2, exp := "aBcDe", "AxCdE", false

	res := EqualFold(str1, str2)

	if res != exp {
		t.Errorf("Expected %t for EqualFold(%s, %s). Instead got %t.", exp, str1, str2, res)
	}
}

func TestFields(t *testing.T) {
	// str, exp := "I am a boy", []string{"I", "am", "a", "boy"}
	// str, exp := " 	", []string{}
	// str, exp := " 	h", []string{"h"}
	str, exp := "	foo bar  baz	", []string{"foo", "bar", "baz"}

	res := Fields(str)

	if slices.Compare(res, exp) != 0 {
		t.Errorf("Expected %q for Fields(%s). Instead got %q.", exp, str, res)
	}
}

func TestFieldsFunc(t *testing.T) {
	// str, exp := "I-am-a-boy", []string{"I", "am", "a", "boy"}
	// str, exp := "", []string{}
	// str, exp := "--------", []string{}
	str, exp := "2024-06-04", []string{"2024", "06", "04"}

	res := FieldsFunc(str, func(r rune) bool {
		return r == '-'
	})

	if slices.Compare(res, exp) != 0 {
		t.Errorf("Expected %q for FieldsFunc(%s, ). Instead got %q.", exp, str, res)
	}
}

func TestPrefix(t *testing.T) {
	str, prefix, exp := "table-water", "table", true
	// str, prefix, exp := "table-water", "fable", false

	res := HasPrefix(str, prefix)

	if res != exp {
		t.Errorf("Expected %t for Prefix(%s, %s). Instead got %t.", exp, str, prefix, res)
	}
}

func TestSuffix(t *testing.T) {
	// str, suffix, exp := "table-water", "water", true
	str, suffix, exp := "table-water", "wafer", false

	res := HasSuffix(str, suffix)

	if res != exp {
		t.Errorf("Expected %t for Suffix(%s, %s). Instead got %t.", exp, str, suffix, res)
	}
}

func TestIndex(t *testing.T) {
	// str, substr, exp := "marmalade", "lade", 5
	str, substr, exp := "marmalade", "mal", 3

	res := Index(str, substr)

	if res != exp {
		t.Errorf("Expected %d for Index(%s, %s). Instead got %d.", exp, str, substr, res)
	}
}
