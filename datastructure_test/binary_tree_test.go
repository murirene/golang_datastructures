package datastructure_test

import (
	"go-ds-rene/datastructure"
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree, err := datastructure.MakeBinaryTree(4)
	if err != nil {
		t.Fatal("Error creating a binary tree")
	}
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(9)
	tree.Insert(10)

	list := tree.ToArray()

	expectedList := []int{1, 4, 5, 9, 10}
	if !reflect.DeepEqual(list, expectedList) {
		t.Fatal("Arrays not the same!")
	}

}
