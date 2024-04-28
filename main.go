package main

import (
	"fmt"
	"i9lib/i9strings"
)

func main() {
	eq := i9strings.EqualFold("aBcDe", "AbCdE")

	fmt.Println(eq)
}
