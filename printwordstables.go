package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
func PrintWordsTables(a []string) {
	for i := range a {
		tab := []rune(a[i])
		for y := range tab {
			z01.PrintRune(tab[y])
		}
		z01.PrintRune('\n')
	}
}
