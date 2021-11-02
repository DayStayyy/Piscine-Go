package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for i := range tab {
		if f(tab[i]) == true {
			nb++
		}
	}
	return nb
}
