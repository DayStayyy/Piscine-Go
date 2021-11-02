package piscine

func ListClear(l *List) {
	for i := 0; l.Head != nil; i++ {
		l.Head = l.Head.Next
	}

}
