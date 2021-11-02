package piscine

func NRune(s string, n int) rune {
	tab := []rune(s)
	taille := 0
	n = n - 1
	if n < 0 {
		return 0
	}
	for i := range tab {
		taille = i
	}
	if n > taille {
		return 0
	}
	return tab[n]
}
