package main

import (
	"fmt"
	"i9lib/i9strings"
)

func main() {
	before, found := i9strings.CutSuffix("TOKEN=adkajhfjh32rkh", "adkajhfjh32rkh")
	fmt.Println(before, found)
}
