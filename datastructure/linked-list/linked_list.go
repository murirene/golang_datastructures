package datastructure

type LinkedList struct {
	head *Node
	tail *Node
}

func (h *LinkedList) Add(value int) *LinkedList {
	if h == nil {
		node := &Node{
			value: value,
			next:  nil,
		}
		return &LinkedList{
			head: node,
			tail: node,
		}
	}
	h.tail.next = &Node{
		value: value,
		next:  nil,
	}
	h.tail = h.tail.next

	return h
}
