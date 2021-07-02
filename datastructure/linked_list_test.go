package datastructure

import (
	"fmt"
	"testing"
)

func TestLinkedListNode(t *testing.T) {
	values := [2]int{1, 2}
	node2 := Node{
		value: values[1],
	}
	node1 := Node{
		value: values[0],
		next:  &node2,
	}

	if node1.value != 1 && node2.value != 2 {
		t.Fatal("node value not set correctly!")
	}

	if node1.next != &node2 && node2.next == nil {
		t.Fatal("node next value is not set to node2")
	}

	var runner *Node = &node1
	for _, value := range values {
		if value != runner.value {
			t.Fatal(fmt.Sprintf("Node is not linked correctly %d %d", runner.value, value))
		}
		runner = node1.next
	}
}

func TestLinkedListAdd(t *testing.T) {
	var list LinkedList
	values := [4]int{100, 200, 300, 500}

	for _, value := range values {
		list.Push(value)
	}

	list_values := list.ToArray()

	if len(values) != len(list_values) {
		t.Fatal(fmt.Sprintf("Size was not the same on the list %d != %d", len(values), len(list_values)))
	}

	for i, v := range values {
		if v != list_values[i] {
			t.Fatal(fmt.Sprintf("%d != %d", v, list_values[i]))
		}
	}

	for list.size > 0 {
		list.Pop()
	}

	if list.head != nil && list.tail != nil && list.size == 0 {
		t.Fatal("List was not empty")
	}
}
