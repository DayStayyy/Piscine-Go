package piscine

func IsAlpha(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 96 && tab[i] < 123) && !(tab[i] > 64 && tab[i] < 91) {
			if !(tab[i] > 47 && tab[i] < 58) {
				return false
			}
		}
	}
	return true
}
