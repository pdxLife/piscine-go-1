package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {

	for i := 97; i <= 122; i++ {

	z01.PrintRune(rune(i))

	}
	fmt.Println()
}
