package datastructure_test

import (
	"fmt"
	"go-ds-rene/datastructure"
	"testing"
)

func TestLinkedListAdd(t *testing.T) {
	list := datastructure.MakeLinkedList()
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

	for list.Size() > 0 {
		list.Pop()
	}

	if !list.IsEmpty() {
		t.Fatal("List was not empty")
	}
}

func TestLinkedListString(t *testing.T) {
	list := datastructure.MakeLinkedList()
	list.Push(4)
	list.Push(5)
	list.Push(1)
	list.Push(3)
	list.Push(2)

	listStr := fmt.Sprintf("%v", list)

	if listStr != "4->5->1->3->2" {
		t.Fatal("Failed to stringify the linked list")
	}
}
