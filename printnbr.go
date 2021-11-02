package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var c, b, t, r, y, u, p int
	c = 1
	a := n

	if a < 0 {

		z01.PrintRune('-')
		if a == -a {
			y += 1
			u = (a % 10) * (-1)
			a = a / 10
			a *= -1

		} else {
			a *= -1
		}

	}

	for a > 9 {
		if (t == 1) && (b == 0) {
			r = 1
		}
		b = b * 10
		b += a % 10

		c = c * 10
		a = a / 10
		t += 1

	}
	if (b * 10) < 0 {
		z01.PrintRune(rune(a + 48))
	} else {
		b = b * 10
		b += a
	}
	for b > 9 {

		a = b % 10
		b = b / 10

		z01.PrintRune(rune(a + 48))

	}
	if n == 0 {
		p = 1
	}

	if p == 1 {
		z01.PrintRune(rune(b + 48))
		p = 3
	} else {
		z01.PrintRune(rune(b + 48))
	}

	if r == 1 {
		z01.PrintRune(48)
	}
	if y == 1 {
		z01.PrintRune(rune(u + 48))
	}
	z01.PrintRune('\n')
}
