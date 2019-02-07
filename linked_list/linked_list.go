package linked_list

import (
	"fmt"
)

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

func (list *List) Delete(value interface{}) bool {
	var prev *Node
	deleted := false
	for cur := list.head; cur != nil; cur = cur.next {
		if cur.value == value {
			if prev == nil {
				list.head = cur.next
				list.count--
				deleted = true
				break
			}
			prev.next = cur.next
			list.count--
			deleted = true
		}
		prev = cur
	}
	return deleted
}

func (list *List) DeleteAll(value interface{}) int {
	var prev *Node
	deletedCount := 0
	for cur := list.head; cur != nil; cur = cur.next {
		if cur.value == value {
			if prev == nil {
				list.head = cur.next
				list.count--
				deletedCount++
			} else {
				prev.next = cur.next
				list.count--
				deletedCount++
			}
			continue
		}
		prev = cur
	}
	return deletedCount
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
