package piscine

func IsLower(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 96 && tab[i] < 123) {
			return false
		}
	}
	return true
}
