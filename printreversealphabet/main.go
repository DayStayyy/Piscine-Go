package main

import "github.com/01-edu/z01"

func main() {

	var lettre rune
	lettre = 122

	// Boucle 26 fois de la valeur 122 decimale a 97
	for i := 1; i < 27; i++ {
		z01.PrintRune(lettre)
		lettre--
	}
	z01.PrintRune('\n')
}
