package linked_list

import "fmt"

type List struct {
	head  *Node
	count int
}

type Node struct {
	value interface{}
	next  *Node
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return list.count == 0
}

func (list *List) AddHead(value interface{}) {
	list.head = &Node{
		value: value,
		next:  list.head,
	}
}

func (list *List) AddTail(value interface{}) {
	node := &Node{
		value: value,
		next:  nil,
	}
	if list.head == nil {
		list.head = node
		return
	}
	var n *Node
	for n = list.head; n.next != nil; n = n.next {
	}
	n.next = node
}

func (list *List) Print() {
	first := true
	for node := list.head; node != nil; node = node.next {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Printf("%v", node.value)
	}
}

func New() *List {
	return &List{}
}
