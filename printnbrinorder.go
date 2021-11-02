package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var tab []int
	b := 10

	for i := 0; n > 10; i++ {
		tab = append(tab, n%10)
		n = n / b
	}
	tab = append(tab, n)
	tab = SortIntegerTable_chiffre(tab)
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(rune(tab[i] + 48))
	}
}
