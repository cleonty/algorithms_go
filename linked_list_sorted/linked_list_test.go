package linkedlist

import "fmt"

func comp(a, b interface{}) bool {
	x := a.(int)
	y := b.(int)
	if x < y {
		return true
	}
	return false
}

func ExampleInsert() {
	list := New(comp)
	list.Insert(3)
	list.Insert(4)
	list.Insert(6)
	list.Insert(2)
	list.Insert(1)
	list.Insert(5)
	list.Print()
	fmt.Printf("\nSize: %d, empty? %v\n", list.Size(), list.IsEmpty())
	// Output:
	// 1 2 3 4 5 6
	// Size: 6, empty? false
}
