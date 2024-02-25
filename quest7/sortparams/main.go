package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	for i := 0; i < len(name)-1; i++ {
		for j := 1 + i; j < len(name); j++ {
			if name[i] > name[j] {
				name[i], name[j] = name[j], name[i]
			}
		}
		// fmt.Println(name)
	}
	for i := 0; i < len(name); i++ {
		for j := 0; j < len(name[i]); j++ {
			z01.PrintRune(rune(name[i][j]))
		}
		z01.PrintRune('\n')

	}
}
