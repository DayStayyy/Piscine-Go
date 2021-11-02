package main

import "github.com/01-edu/z01"

func main() {

	var lettre rune
	lettre = 48

	for i := 1; i < 11; i++ {
		z01.PrintRune(lettre)
		lettre++
	}
	z01.PrintRune('\n')

}
