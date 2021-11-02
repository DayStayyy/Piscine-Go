package piscine

import "github.com/01-edu/z01"

var f, as, bs, cs string
var a, b, c, d, e, g, z rune

func PrintComb() {
	for e != 3 {
		for a >= b {
			b += 1
		}
		for b >= c {
			c += 1
		}
		if c > 9 {
			c = 1
			b += 1
			if b >= 10 {
				b = 0
				a += 1
			}
		}
		if a < b {
			d += 1
		}
		if b < c {
			d += 1
		}
		if c <= 10 {
			d += 1
		}
		if d == 3 {
			z01.PrintRune(a + 48)
			z01.PrintRune(b + 48)
			z01.PrintRune(c + 48)
			g++
			if g == 120 {
				z01.PrintRune('\n')
				e = 3
				break
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
				c += 1
			}

		} else {
			d = 0
		}
	}

}
