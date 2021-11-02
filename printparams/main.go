package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	param := os.Args
	for i := 1; i < len(param); i++ {
		tab := []rune(param[i])
		for i := range tab {
			z01.PrintRune(tab[i])
		}
		z01.PrintRune('\n')
	}

}
