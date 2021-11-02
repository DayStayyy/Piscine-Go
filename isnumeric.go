package piscine

func IsNumeric(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 47 && tab[i] < 58) {
			return false
		}
	}
	return true
}
