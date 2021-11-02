package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []string(os.Args)
	valeurTransformer := 96
	trucdefou := 1
	if len(arg) > 1 {
		if arg[1] == "--upper" {
			valeurTransformer = 64
			trucdefou = 2
		}
	}
	for i := trucdefou; i < len(arg); i++ {
		chiffre := Atoi(arg[i])
		if chiffre > 0 && chiffre < 27 {
			z01.PrintRune(rune(chiffre + valeurTransformer))
		} else {
			z01.PrintRune(32)
		}

	}
	if len(arg) > 1 {
		z01.PrintRune('\n')
	}

}

func Atoi(s string) int {
	var b, y int
	b = 1
	runes := []rune(s)
	nombres := 0
	if !(len(runes)-1 < 0) {
		if runes[0] == 43 {
			runes[0] = 48
		}
		if runes[0] == 45 {
			runes[0] = 48
			y = 1

		}
	}

	for i := 0; i < len(runes); i++ {
		if (int(runes[i]-48) != 0) && (int(runes[i]-48) != 1) && (int(runes[i]-48) != 2) && (int(runes[i]-48) != 3) && (int(runes[i]-48) != 4) && (int(runes[i]-48) != 5) && (int(runes[i]-48) != 6) && (int(runes[i]-48) != 7) && (int(runes[i]-48) != 8) && (int(runes[i]-48) != 9) {
			b = 0
			break

		}
	}

	for i := 0; i < len(runes)-1 && b == 1; i++ {

		if len(runes)-1 < 0 {
			break
		}
		nombres += int(runes[i] - 48)
		nombres = nombres * 10
	}
	if !(len(runes)-1 < 0) {
		nombres += int(runes[len(runes)-1] - 48)
	}
	if y == 1 {
		nombres = nombres * -1
	}
	if b == 1 {
		return (nombres)

	} else {
		return (0)
	}

}
