package binarytree

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type Queue struct {
	count int
	Head  *BinaryTreeNode
	Tail  *BinaryTreeNode
}
