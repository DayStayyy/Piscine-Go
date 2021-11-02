package piscine

func ListMerge(l1 *List, l2 *List) {
	for l2.Head != nil {
		ListPushBack(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}
	/*if l2.Head != nil && l1.Head != nil {
		l1.Tail.Next = l2.Head
	}*/

}
