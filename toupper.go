package piscine

func ToUpper(s string) string {
	tab := []rune(s)
	for i := range tab {
		if tab[i] > 96 && tab[i] < 123 {
			tab[i] = tab[i] - 32
		}
	}
	return string(tab)
}
