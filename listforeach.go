package piscine

func ListForEach(l *List, f func(*NodeL)) {
	first := l.Head
	for l.Head != nil {
		f(l.Head)
		l.Head = l.Head.Next
	}
	l.Head = first

}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func ListSize2(l *List) int {
	result := 0
	for l.Head != nil {
		result++
		l.Head = l.Head.Next
	}
	return result
}
