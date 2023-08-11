package main

import (
	"fmt"
	linkedList "sammple/linked-list"
)

func main() {
	fmt.Println("List 1")
	list1 := linkedList.CreateList(5)
	fmt.Println("List 2")
	list2 := linkedList.CreateList(5)
	fmt.Println("Result list")
	result := linkedList.AddToListWithSingleDigitOutPut(list1, list2)
	linkedList.ListTraverse(result)
}
