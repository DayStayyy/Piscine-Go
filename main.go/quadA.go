package main

import "os"

func main() {
	QuadA()
}

func QuadA() {
	x := 0
	y := 0
	tab := os.Args
	if len(tab) == 3 {
		boolean1 := IsNumeric(tab[1])
		boolean2 := IsNumeric(tab[2])
		if boolean1 == true && boolean2 == true {
			x = Atoi(tab[1])
			y = Atoi(tab[2])
		}

	}
	if x >= 1 && y >= 1 {
		for i := y; i > 0; i-- {
			for o := x; o > 0; o-- {
				if (o == 1 || o == x) && (i == 1 || i == y) {
					print("o")
				} else if i == 1 || i == y {
					print("-")
				} else if o == 1 || o == x {
					print("|")
				} else {
					print(" ")
				}
			}
			print("\n")
		}
	}
}

func IsNumeric(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 47 && tab[i] < 58) {
			return false
		}
	}
	return true
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
