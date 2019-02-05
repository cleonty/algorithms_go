package linkedlist

import (
	"fmt"
)

type Comparator func(a, b interface{}) bool

type List struct {
	head  *Node
	count int
	comp  Comparator
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

func (list *List) Insert(value interface{}) {
	node := &Node{
		value: value,
		next:  nil,
	}
	list.count++
	if list.head == nil {
		list.head = node
		return
	}
	var cur, prev *Node
	for cur = list.head; cur != nil; cur = cur.next {
		if list.comp(value, cur.value) {
			if prev != nil {
				prev.next = node
				node.next = cur
				return
			}
			node.next = list.head
			list.head = node
			return
		}
		prev = cur
	}
	prev.next = node
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

func New(comp Comparator) *List {
	return &List{
		head:  nil,
		count: 0,
		comp:  comp,
	}
}
