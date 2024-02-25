package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

// package piscine

func main() {
	// name := []string{"8", "5", "12", "12", "15"}
	if len(os.Args) == 1 {
		return
	}
	name := os.Args[1:]
	isupper := false
	slice := []int{}
	if name[0] == "--upper" {
		isupper = true
	}
	start := 0
	if isupper {
		start = 1
	}

	for i := start; i < len(name); i++ {
		num := 0

		for j := 0; j < len(name[i]); j++ {
			if name[i][j] >= '0' && name[i][j] <= '9' {

				num = num*10 + int(name[i][j]-'0')
				if len(name[i]) == j+1 {
					slice = append(slice, num)
				}

			} else {
				slice = append(slice, -1)
			}
		}

	}

	for _, value := range slice {
		if isupper {
			alpharange := value + 64
			if alpharange > 90 || value == -1 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(rune(alpharange))
			}

		} else {
			alpharange := value + 96
			if alpharange > 122 || value == -1 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(rune(alpharange))
			}

		}
	}
	if len(name) > 1 {
		z01.PrintRune('\n')
	}
	// if name[0] == "--upper" {

	// 	for i := 1; i < len(name); i++ {
	// 		if condition {
	// 			x := name[i] - '0'

	// 			name[i] = string('A' + x)
	// 		}

	// 	}
	// }
}
