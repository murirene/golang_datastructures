package datastructure

import "errors" 

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (h *LinkedList) Push(value int) {
	node := Node{
		value: value,
		next:  nil,
        prev:  nil,
	}

	if h == nil {
		h = &LinkedList{
			head: &node,
			tail: &node,
		}
	} else if h.tail == nil && h.head == nil {
		h.tail = &node
		h.head = &node
	} else {
		h.tail.next = &node
        node.prev = h.tail
		h.tail = h.tail.next
	}

	h.size = h.size + 1
}

func (h* LinkedList) Pop() (int, error) {
	if h.head == nil && h.tail == nil {
		return 0, errors.New("Can't pop an empty array\n")
	}

    h.size = h.size - 1
    value := h.head.value
    if h.head == h.tail {
        h.head = nil;
        h.tail = nil;
        h.size = 0
    } else {
        h.head = h.head.next
    }
    return value, nil
}

func (h *LinkedList) IsEmpty() bool {
    return h.size == 0
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
