package piscine

func Tierlist(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Data > root.Left.Data {
			return Tierlist(root.Left)
		} else {
			return false
		}

	} else if root.Right != nil {
		if root.Data < root.Right.Data {
			return Tierlist(root.Right)
		} else {
			return false
		}
	}
	return true

}
