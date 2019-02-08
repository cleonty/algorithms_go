package linkedlist

func ExampleList_AddHead() {
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

func ExampleList_AddTail() {
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

func ExampleList_Delete() {
	list := New()
	list.AddHead(3)
	list.AddHead(2)
	list.AddHead(1)
	list.AddTail(4)
	list.AddTail(5)
	list.AddTail(6)
	list.Delete(3)
	list.Delete(2)
	list.Print()
	// Output: 1 4 5 6
}

func ExampleList_DeleteAll() {
	list := New()
	list.AddHead(3)
	list.AddHead(2)
	list.AddHead(3)
	list.AddHead(1)
	list.AddTail(4)
	list.AddHead(3)
	list.AddTail(5)
	list.AddTail(6)
	list.DeleteAll(3)
	list.Print()
	// Output: 1 2 4 5 6
}

func ExampleList_Reverse() {
	list := New()
	list.AddHead(3)
	list.AddHead(2)
	list.AddHead(1)
	list.AddTail(4)
	list.AddTail(5)
	list.AddTail(6)
	list.Reverse()
	list.Print()
	// Output: 6 5 4 3 2 1
}
