package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for n2 != nil {
		n1 = listPushBack(n1, n2.Data)
		n2 = n2.Next
	}
	n1 = ListSort(n1)
	return n1
}
