package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	issort := false
	result := l
	if l == nil {
		return nil
	}
	if l.Next == nil {
		return l
	}
	for !issort {
		curr := l
		issort = true
		for curr.Next != nil {
			next := curr.Next
			if curr.Data > next.Data {
				curr.Data, next.Data = next.Data, curr.Data
				issort = false
			}
			curr = next
		}
	}
	return result
}
