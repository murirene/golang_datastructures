package datastructure
import "errors"

type TNode struct {
    left *TNode
    right *TNode
    value int
}

type BinaryTree struct {
    root *TNode
}


func MakeBinaryTree(value int) (BinaryTree, error) {
    tree := BinaryTree {
        root: &TNode {
            left: nil,
            right: nil,
            value: value,
        },
    }

    if tree.root.value != value {
        return tree, errors.New("Failed to create BinaryTree")
    }

   return tree, nil
}

func (treePtr *BinaryTree) Insert(value int) error {
    if ( treePtr == nil ){
       tree, err := MakeBinaryTree(value)
       if err != nil {
           return err
       }
       treePtr = &tree
       return nil
    }
    node := TNode {
            value: value,
            right: nil,
            left: nil,
    }

    runner := treePtr.root
    for {
        if runner.value >= value && runner.left != nil {
            runner = runner.left
        } else if  runner.value < value && runner.right != nil {
            runner = runner.right
        } else if runner.value >= value {
            runner.left = &node
            break;
        } else if runner.value < value {
            runner.right = &node
            break;
        } else {
            return errors.New("Error inserting")
        }
    }

    return nil
}
