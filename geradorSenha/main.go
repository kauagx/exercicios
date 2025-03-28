package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	chars := "aeiou"
	numero := rand.IntN(5)
	fmt.Print(string(chars[numero]))
}
