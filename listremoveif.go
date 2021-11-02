package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	link3 := &List{}
	for l.Head != nil {
		if l.Head.Data != data_ref {
			ListPushBack(link3, l.Head.Data)
		}
		l.Head = l.Head.Next
	}
	l.Head = link3.Head

}
