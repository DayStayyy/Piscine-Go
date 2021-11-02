package piscine

func Swap(a *int, b *int) {
	var z int
	z = *a + 1 - 1
	*a = *b
	*b = z

}
