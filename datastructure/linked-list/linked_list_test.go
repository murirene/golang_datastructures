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
