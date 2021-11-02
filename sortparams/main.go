package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []string(os.Args)
	texte := []string{}
	for i := 1; i < len(arg); i++ {
		texte = append(texte, arg[i])
	}
	texte = SortIntegerTable_lettre(texte)

	for i := range texte {
		tab := []rune(texte[i])
		for i := range tab {
			z01.PrintRune(tab[i])
		}
		z01.PrintRune('\n')
	}
	/*
		texte := ""
		for i := len(arg) - 1; i > 0; i-- {
			texte += arg[i]
		}
		param := []rune(texte)
		chiffre := []int{}
		lettre := []string{}
		for i := len(param) - 1; i >= 0; i-- {
			if param[i] == 32 {
				continue
			} else if param[i] > 47 && param[i] < 58 {
				chiffre = append(chiffre, int(param[i]-48))
			} else {
				lettre = append(lettre, string(param[i]))
			}

		}
		chiffre = SortIntegerTable_chiffre(chiffre)
		lettre = SortIntegerTable_lettre(lettre)

		for i := range chiffre {
			z01.PrintRune(rune(chiffre[i] + 48))
			z01.PrintRune('\n')
		}
		texte = ""
		for i := range lettre {
			texte += lettre[i]
		}
		truc := []rune(texte)
		for i := range truc {
			z01.PrintRune(rune(truc[i]))
			z01.PrintRune('\n')
		}*/
}

func SortIntegerTable_lettre(table []string) []string {
	length := len(table)

	for i := 0; i < length; i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
	return (table)
}

func SortIntegerTable_chiffre(table []int) []int {
	length := len(table)

	for i := 0; i < length; i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
	return (table)
}
func remove(s []rune, i int) []rune {
	return append(s[:i], s[i+1:]...)
}
