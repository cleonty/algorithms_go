package linked_list

func ExampleOne() {
	list := New()
	list.AddHead(3)
	list.AddHead(2)
	list.AddHead(1)
	list.AddTail(4)
	list.AddTail(5)
	list.AddTail(6)
	list.Print()
	// Output: 1 2 3 4 5 6
}
