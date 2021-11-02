package piscine

func AlphaCount(s string) int {
	tab := []rune(s)
	nb := 0
	for i := range tab {
		if (tab[i] > 96 && tab[i] < 123) || (tab[i] > 64 && tab[i] < 91) {
			nb += 1

		}
	}
	return (nb)

}
