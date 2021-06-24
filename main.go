package main

import (
	"fmt"
	"go-ds-rene/datastructure"
)

func main() {
	var list datastructure.LinkedList
	list.Push(1)
	list.Push(2)
	list.Push(3)
    var arr []int = list.ToArray()
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

    for !list.IsEmpty() {
        value, err := list.Pop()
        if err != nil {
            fmt.Println(err)
            break
        }
        fmt.Printf("*Pop %d\n", value)
    }
}
