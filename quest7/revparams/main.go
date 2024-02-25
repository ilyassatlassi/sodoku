package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for i := len(name) - 1; i >= 1; i-- {
		for j := 0; j < len(name[i]); j++ {
			z01.PrintRune(rune(name[i][j]))
		}
		z01.PrintRune('\n')

	}
	// for _, value := range name  {
	// 	for _, word := range value {
	// 		z01.PrintRune(word)
	// 	}

	// 	z01.PrintRune('\n')
	// }
}
