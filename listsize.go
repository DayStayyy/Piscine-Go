package piscine

func ListSize(l *List) int {
	result := 0
	for l.Head != nil {
		result++
		l.Head = l.Head.Next
	}
	return result
}
