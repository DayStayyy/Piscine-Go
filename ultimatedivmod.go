package piscine

func UltimateDivMod(a *int, b *int) {
	var div, modulo int
	div = *a / *b
	modulo = *a % *b
	*a = div
	*b = modulo
}
