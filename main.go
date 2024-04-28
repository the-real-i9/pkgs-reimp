package main

import (
	"fmt"
	"i9lib/i9strings"
)

func main() {
	before, after, found := i9strings.Cut("TOKEN=adkajhfjh32rkh", "=")
	fmt.Println(before, after, found)
}
