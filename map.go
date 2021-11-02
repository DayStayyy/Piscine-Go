package piscine

func Map(f func(int) bool, a []int) []bool {
	var tab []bool
	for i := range a {
		tab = append(tab, f(a[i]))
	}
	return tab
}
