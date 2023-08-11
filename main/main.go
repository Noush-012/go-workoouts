package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("List 1")
	list1 := CreateList(5)
	fmt.Println("List 2")
	list2 := CreateList(5)
	fmt.Println("Result list")
	result := AddToListWithSingleDigitOutPut(list1, list2)
	listTraverse(result)

}
func MakeFirstLetterCaps(s string) string {

	var builder strings.Builder

	word := strings.Fields(s)
	fmt.Println(word)
	var res string
	for _, v := range word {

		builder.WriteString(strings.Title(v + " "))
	}
	res = builder.String()

	return res
}

func MakeFirstLetterCapsUsingSlice(s string) {
	word := strings.Fields(s)
	var res string

	for i, v := range word {
		if v[0] >= 'a' && v[0] <= 'z' {
			res = string(rune(v[0] - 32))
			word[i] = res + v[1:]

		} else {
			word[i] = v
		}

	}

	// fmt.Println("----", word)
	// fmt.Println("77", strings.Join(word, " "))
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

func CreateList(limit int) *list.List {
	list := list.New()

	for i := 0; i < limit; i++ {
		list.PushBack(rand.Intn(20))
	}
	listTraverse(list)

	return list

}

func listTraverse(list *list.List) {
	indx := 0

	for elem := list.Front(); elem != nil; elem = elem.Next() {
		indx++
		fmt.Printf(" %v ->", elem.Value)

	}
	fmt.Printf("\nList length: %v\n", indx)

}
