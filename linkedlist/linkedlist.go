// linkedlist package contains implementation of a singly linked list.
package linkedlist

import (
	"fmt"
)

// List is singly linked list
type List struct {
	head  *Node
	count int
}

// Node represents a single element of the list
type Node struct {
	value interface{}
	next  *Node
}

// Size returns number of elements in the list
func (list *List) Size() int {
	return list.count
}

// IsEmpty checks whether the list is empty
func (list *List) IsEmpty() bool {
	return list.count == 0
}

// AddHead inserts value into the head of the list
func (list *List) AddHead(value interface{}) {
	list.head = &Node{
		value: value,
		next:  list.head,
	}
}

// AddTail inserts value into the tail of the list
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

// Delete deletes a first occurrence of value
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

// DeleteAll deletes all occurrences of value and returns number of deleted elements
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

// Reverse reverses a list
func (list *List) Reverse() {
	var cur *Node
	var prev *Node
	for cur = list.head; cur != nil; {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	list.head = prev
}

// Print prints a list into stdout
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

// New creates a new List
func New() *List {
	return &List{}
}
