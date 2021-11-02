package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil || rplc == nil {
		return nil
	}

	elem := node.Data
	if root == nil {
		root = rplc
	} else if elem < root.Data {
		root = BTreeTransplant(root.Left, node, rplc)
	} else if elem > root.Data {
		root = BTreeTransplant(root.Right, node, rplc)
	} else if elem == root.Data {
		root.Data = rplc.Data
	}
	rplc.Parent = node.Parent
	return root
}
