package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
}
