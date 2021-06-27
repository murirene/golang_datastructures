package datastructure_test

import ( 
"testing"
"go-ds-rene/datastructure"
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
}
