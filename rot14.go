package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			if (runes[i] + 14) > 122 {
				a := runes[i] + 14 - 122
				runes[i] = 96 + a
			} else {
				runes[i] += 14
			}
		}
		if runes[i] >= 65 && runes[i] <= 90 {
			if (runes[i] + 14) > 90 {
				a := runes[i] + 14 - 90
				runes[i] = 64 + a
			} else {
				runes[i] += 14
			}
		}
	}
	return string(runes)
}
