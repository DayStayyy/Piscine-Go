package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return nil
	}
	elem := node.Data
	if elem < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if elem > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		root.Data = BTreeMin(root.Right).Data
		root.Right = BTreeDeleteNode(root.Right, BTreeMin(root.Right))
	}

	return root

}
