package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if (k*10 + l) > (i*10 + j) {

						if i == 0 {
							z01.PrintRune(48)
						} else if i == 1 {
							z01.PrintRune(49)
						} else if i == 2 {
							z01.PrintRune(50)
						} else if i == 3 {
							z01.PrintRune(51)
						} else if i == 4 {
							z01.PrintRune(52)
						} else if i == 5 {
							z01.PrintRune(53)
						} else if i == 6 {
							z01.PrintRune(54)
						} else if i == 7 {
							z01.PrintRune(55)
						} else if i == 8 {
							z01.PrintRune(56)
						} else if i == 9 {
							z01.PrintRune(57)
						}
						if j == 0 {
							z01.PrintRune(48)
						} else if j == 1 {
							z01.PrintRune(49)
						} else if j == 2 {
							z01.PrintRune(50)
						} else if j == 3 {
							z01.PrintRune(51)
						} else if j == 4 {
							z01.PrintRune(52)
						} else if j == 5 {
							z01.PrintRune(53)
						} else if j == 6 {
							z01.PrintRune(54)
						} else if j == 7 {
							z01.PrintRune(55)
						} else if j == 8 {
							z01.PrintRune(56)
						} else if j == 9 {
							z01.PrintRune(57)
						}
						z01.PrintRune(32)

						if k == 0 {
							z01.PrintRune(48)
						} else if k == 1 {
							z01.PrintRune(49)
						} else if k == 2 {
							z01.PrintRune(50)
						} else if k == 3 {
							z01.PrintRune(51)
						} else if k == 4 {
							z01.PrintRune(52)
						} else if k == 5 {
							z01.PrintRune(53)
						} else if k == 6 {
							z01.PrintRune(54)
						} else if k == 7 {
							z01.PrintRune(55)
						} else if k == 8 {
							z01.PrintRune(56)
						} else if k == 9 {
							z01.PrintRune(57)
						}
						if l == 0 {
							z01.PrintRune(48)
						} else if l == 1 {
							z01.PrintRune(49)
						} else if l == 2 {
							z01.PrintRune(50)
						} else if l == 3 {
							z01.PrintRune(51)
						} else if l == 4 {
							z01.PrintRune(52)
						} else if l == 5 {
							z01.PrintRune(53)
						} else if l == 6 {
							z01.PrintRune(54)
						} else if l == 7 {
							z01.PrintRune(55)
						} else if l == 8 {
							z01.PrintRune(56)
						} else if l == 9 {
							z01.PrintRune(57)
						}
						if i == 9 && j == 8 && k == 9 && l == 9 {
							z01.PrintRune('\n')
							break
						}
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}
