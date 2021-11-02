package piscine

func Trucdemerde(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	root.Left = Trucdemerde(root.Left, data)
	root.Left.Parent = root
	return root

}
