package binomialheap

type BinomialHeapNode struct {
	Value   int               // value of the node
	Degree  int               // number of children
	Parent  *BinomialHeapNode // points to parent node
	Child   *BinomialHeapNode // points to left most child
	Sibling *BinomialHeapNode // points to right side sibling
}

type LinkedList struct {
	Head *BinomialHeapNode
	Size int
}
