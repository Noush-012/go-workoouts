package linkedList

import (
	"container/list"
	"fmt"
	"math/rand"
)

func CreateList(limit int) *list.List {
	list := list.New()

	for i := 0; i < limit; i++ {
		list.PushBack(rand.Intn(20))
	}
	ListTraverse(list)

	return list

}
func ListTraverse(list *list.List) {
	indx := 0

	for elem := list.Front(); elem != nil; elem = elem.Next() {
		indx++
		fmt.Printf(" %v ->", elem.Value)

	}
	fmt.Printf("\nList length: %v\n", indx)
}

func AddToListWithSingleDigitOutPut(list1, list2 *list.List) *list.List {
	Reslist := list.New()
	var e1, e2 int
	for elem1, elem2 := list1.Front(), list2.Front(); elem1 != nil; elem1, elem2 = elem1.Next(), elem2.Next() {
		e1, _ = elem1.Value.(int)
		e2, _ = elem2.Value.(int)
		sum := e1 + e2
		// for sum > 0 {
		// 	dig := sum % 10
		// 	fmt.Println("-", dig)
		// 	digit = append(digit, dig)
		// 	fmt.Println("--", digit)
		// 	num /= 10
		// }
		Reslist.PushBack(sum)
	}
	return Reslist
}

// func IsCyclicList(list list.List) bool {
// 	currentNode := list.Front()

// 	for i != nil {

// 	}
// }
