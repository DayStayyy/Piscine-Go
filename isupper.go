package piscine

func IsUpper(s string) bool {
	tab := []rune(s)
	for i := range tab {
		if !(tab[i] > 64 && tab[i] < 91) {
			return false
		}
	}
	return true
}
