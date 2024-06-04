package i9strings

import (
	"slices"
	"testing"
)

func TestEqualFold(t *testing.T) {
	str1 := "aBcDe"
	str2 := "AbCdE"

	res := EqualFold(str1, str2)

	exp := true

	if res != exp {
		t.Errorf("Expected %t for EqualFold(%s, %s). Instead got %t.", exp, str1, str2, res)
	}
}

func TestFields(t *testing.T) {
	// str := "I am a boy"
	// str := " 	"
	// str := " 	h"
	str := "	foo bar  baz	"

	res := Fields(str)

	// exp := []string{"I", "am", "a", "boy"}
	// exp := []string{}
	// exp := []string{"h"}
	exp := []string{"foo", "bar", "baz"}

	if slices.Compare(res, exp) != 0 {
		t.Errorf("Expected %q for Fields(%s). Instead got %q.", exp, str, res)
	}
}
