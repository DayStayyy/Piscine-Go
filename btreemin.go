package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil {
		return root
	} else {
		root = BTreeMin(root.Left)
	}
	return root
}
