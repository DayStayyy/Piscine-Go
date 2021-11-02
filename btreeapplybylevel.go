package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	depth := BTreeLevelCount(root)
	for i := 1; i < depth+1; i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		Level(root.Left, level-1, f)
		Level(root.Right, level-1, f)
	}
}
