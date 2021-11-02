package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	texte := "x = 42, y = 21\n"
	tab := []rune(texte)
	for i := range tab {
		z01.PrintRune(tab[i])
	}
}
