package piscine

func TrimAtoi(s string) int {
	tab := []rune(s)
	for i := len(tab) - 1; i >= 0; i-- {
		if !(tab[i] > 47 && tab[i] < 58) && !(tab[i] == 45) {
			tab = remove(tab, i)
		}

	}
	if len(tab) == 0 {
		return 0
	}
	return Atoi2(string(tab))

}

func Atoi2(s string) int {
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
			if runes[i] == 45 {
				runes = remove(runes, i)
			}
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
