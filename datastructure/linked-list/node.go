package datastructure

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) SetNext(value int) *Node {
	next := Node{
		value: value,
		next:  nil,
	}

	if n == nil {
		return &next
	}

	*n.next = next
	return &next
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}

	return fmt.Sprintf("%d\n", n.value)
}
