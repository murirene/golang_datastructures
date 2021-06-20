package datastructure

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (h *LinkedList) Add(value int) *LinkedList {
	if h.head == nil {
		node := &Node{
			value: value,
			next:  nil,
		}
		return &LinkedList{
			head: node,
			tail: node,
			size: 1,
		}
	}
	h.tail.next = &Node{
		value: value,
		next:  nil,
	}
	h.size = h.size + 1
	h.tail = h.tail.next

	return h
}

func (h *LinkedList) ToArray() []int {
	list := make([]int, h.size)
	i := 0
	for runner := h.head; runner != nil; runner = runner.next {
		list[i] = runner.value
		i++
	}

	return list
}
