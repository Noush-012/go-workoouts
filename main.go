package main

import (
	"fmt"
	"runtime"
)

//           ----------           To Do              ------------                         //

// Linked list find cyclic

func main() {
	// fmt.Println("List 1")
	// list1 := linkedList.CreateList(5)
	// fmt.Println("List 2")
	// list2 := linkedList.CreateList(5)
	// fmt.Println("Result list")
	// result := linkedList.AddToListWithSingleDigitOutPut(list1, list2)
	// linkedList.ListTraverse(result)

	// str := "A man, a plan, a canal: Panama"

	// strn := strings.TrimFunc(str, func(r rune) bool {
	// 	return unicode.IsSpace(' ') && unicode.IsSymbol(':') && unicode.IsSymbol(',')
	// })

	// t1 := stringOps.BinaryAddition("1010", "1011")

	// t2 := stringOps.BinaryAdditionMath("1010", "1011")

	// fmt.Println(t1, t2)

	// api.GetAllCarMakes()
	// hashmap.Mapping()

	// c := Closure()
	// fmt.Println(c(5))
	// fmt.Println(c(5))
	// fmt.Println(c(5))

	// i := new(person)
	// i.name = "John"

	// ManualGarbageCollection()

	// arr := []int{2, 6, 89, 8, -8, 7, 6, 9, 25, 1, 3}
	// sort.Ints(arr)
	// fmt.Println(arr)
	// fib := FibClosure()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(fib())
	// }

	// id := "123"

	// employeeData, _ := api.CallEmployeeDetails(id)
	// emplyrData, _ := api.CallEmployerDetails(id)

	// final := api.Data{
	// 	Employee: employeeData,
	// 	Employer: emplyrData,
	// }

	// fmt.Println(final)

	// api.GetAllCarMakes()
	//str := "Nou sha                             d"

	//res := strings.ReplaceAll(str, " ", "")
	//res := strings.TrimSpace(str)

	//for i := len(str) - 1; i >= 0; i-- {

	//res += string(str[i])

	//}
	//fmt.Println(res)

	// var a interface{} = 10
	// fmt.Println(a)
	// b := reflect.Int
	// fmt.Println(b)

	// a = b
	// fmt.Println(a)

	// fmt.Println(reflect.DeepEqual(b, a))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(integer.FibonacciR(i))
	// }

}

func FibClosure() func() int {
	n1 := 0
	n2 := 1

	return func() int {

		t3 := n1
		n1, n2 = n2, n1+n2

		return t3
	}
}

type Person struct {
	Name string
}

func CreatePerson() *Person {
	p := &Person{
		Name: "Anas",
	}

	if p.Name == "" {
		return &Person{}
	}
	return nil
}

func ManualGarbageCollection() {
	// Allocate memory using the runtime package
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("Initial memory allocated: %d bytes\n", mem.Alloc)

	arr := []int{1, 5, 55, 216, 16, 5}

	arr = append(arr, 1)

	// Deallocate memory manually
	runtime.GC()

	// Check memory after garbage collection
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory allocated after GC: %d bytes\n", mem.Alloc)
}
