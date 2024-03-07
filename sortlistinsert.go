package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		return nil
	}
	if l.Data > newNode.Data {
		newNode.Next = l
		l = newNode
		return l
	}
	current := l

	for current.Next != nil {
		prev := current
		current = current.Next
		if current.Data > newNode.Data {
			newNode.Next = current
			prev.Next = newNode
			return l
		} else if current.Next == nil {
			current.Next = newNode
			newNode.Next = nil
		}
	}
	return l
}
