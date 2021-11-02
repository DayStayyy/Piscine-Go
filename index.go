package piscine

func Index(s string, toFind string) int {
	tab := []rune(s)
	tabverif := []rune(toFind)
	nb := len(tabverif)

	for i := 0; i+nb < len(tab); i++ {
		if string(tab)[i:i+nb] == toFind {
			return i
		}
	}
	return -1
}
