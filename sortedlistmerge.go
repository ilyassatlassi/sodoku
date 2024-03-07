package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	head := n1
	for head.Next != nil {
		head = head.Next
	}
	head.Next = n2
	return ListSort(n1)
}
