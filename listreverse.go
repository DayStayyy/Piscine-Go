package piscine

func ListReverse(l *List) {
	var temp *NodeL
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next := current.Next
		current.Next = temp
		temp = current
		current = next
	}
	l.Head = temp
}
