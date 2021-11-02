package piscine

func ToLower(s string) string {
	tab := []rune(s)
	for i := range tab {
		if tab[i] > 64 && tab[i] < 91 {
			tab[i] = tab[i] + 32
		}
	}
	return string(tab)
}
