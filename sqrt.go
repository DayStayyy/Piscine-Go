package piscine

func Sqrt(nb int) int {
	j := 1
	for i := 1; 1 == 1; i, j = i+2, j+1 {
		nb -= i
		if nb == 0 {
			return (j)
		} else if nb < 0 {
			return (0)
		}
	}
	return (j)

}
