package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var tab []int
	first := l
	for l != nil {
		tab = append(tab, l.Data)
		l = l.Next
	}
	l = first

	SortIntegerTable(tab)
	for i := 0; l != nil && i < len(tab); i++ {
		l.Data = tab[i]
		l = l.Next
	}
	l = first

	return l
}
