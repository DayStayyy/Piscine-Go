package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		if root.Left != nil {
			if root.Data < root.Left.Data {
				return false
			}
		} else if root.Right != nil {
			if root.Data > root.Right.Data {
				return false
			}
		}

		if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
			return false
		}
	}
	return true

}
