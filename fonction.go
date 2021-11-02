package piscine

func SortIntegerTable_lettre(table []string) []string {
	length := len(table)

	for i := 0; i < length; i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
	return (table)
}

func SortIntegerTable_chiffre(table []int) []int {
	length := len(table)

	for i := 0; i < length; i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
	return (table)
}

func remove(s []rune, i int) []rune {
	return append(s[:i], s[i+1:]...)
}
