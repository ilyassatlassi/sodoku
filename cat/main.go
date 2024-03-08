package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printtxt(str string) {
	for _, i := range str {
		z01.PrintRune(i)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input, _ := io.ReadAll(os.Stdin)
		printtxt(string(input))
	} else {
		for i := 0; i < len(args); i++ {
			content, err := os.ReadFile(args[i])
			if err != nil {
				printtxt("ERROR: ")
				printtxt(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printtxt(string(content))
		}
	}
}
