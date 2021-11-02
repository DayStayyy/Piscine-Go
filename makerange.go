package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var tab []int
		return (tab)
	}
	var tab = make([]int, max-min)
	for i := min; i < max; i++ {
		tab[i-min] = i
	}

	return (tab)
}
