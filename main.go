package main

import (
	"fmt"
	datastructure "go-ds-rene/datastructure/linked-list"
)

func main() {
	var list datastructure.LinkedList
	list = *list.Add(1)
	list = *list.Add(2)
	list = *list.Add(3)
	var arr []int = list.ToArray()
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
