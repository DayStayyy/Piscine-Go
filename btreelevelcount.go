package piscine

func BTreeLevelCount(root *TreeNode) int {
	nbgauche := 0
	nbdroite := 0
	if root != nil {
		nbgauche = BTreeLevelCount(root.Left)
		nbdroite = BTreeLevelCount(root.Right)
	} else {
		return 0
	}
	if nbgauche > nbdroite {
		return nbgauche + 1
	} else {
		return nbdroite + 1
	}
}
