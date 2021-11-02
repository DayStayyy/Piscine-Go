package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	b := false
	order := false
	count := 0
	tab := a
	for index := range tab {
		count = index
	}
	if count == 0 {
		return true
	}
	if tab[0] < tab[1] {
		order = true
	}
	if order {
		for i := 0; i < count+1; i++ {
			for j := 0; j < count-i; j++ {
				if f(tab[j], tab[j+1]) <= 0 {
					b = true
				} else {
					return false

				}
			}
		}
	} else {
		for i := 0; i < count+1; i++ {
			for j := 0; j < count-i; j++ {
				if f(tab[j], tab[j+1]) >= 0 {
					b = true
				} else {
					return false

				}
			}
		}
	}
	return b
}
