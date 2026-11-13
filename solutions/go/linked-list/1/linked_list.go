package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}

type List[T comparable] struct {
	first *Node[T]
	last  *Node[T]
	size  int
}

// NewNode constructs a new Node with the given value & no next/prev links.
func NewNode[T comparable](v T) *Node[T] {
	return &Node[T]{Value: v, prev: nil, next: nil}
}

// NewList constructs a new List with the given values
func NewList[T comparable](elements ...T) *List[T] {
	if len(elements) == 0 {
		return &List[T]{first: nil, last: nil, size: 0}
	}

	dummy := &Node[T]{next: nil, prev: nil}
	cur := dummy
	for _, e := range elements {
		newNode := &Node[T]{Value: e, next: nil, prev: cur}
		cur.next = newNode
		cur = cur.next
	}
	newHead := dummy.next
	newHead.prev = nil
	newTail := cur
	return &List[T]{first: newHead, last: newTail, size: len(elements)}
}

// Next returns a pointer to the next node in the linked list.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev returns a pointer to the previous node in the linked list.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// Unshift pushes a new value before Head.
func (ll *List[T]) Unshift(v T) {
	newHead := &Node[T]{Value: v, prev: nil, next: ll.first}
	if ll.first != nil {
		ll.first.prev = newHead
	}
	ll.first = newHead
	if ll.last == nil {
		ll.last = newHead
	}
	ll.size += 1 // maintain list size
}

// Push pushes a new value after Tail.
func (ll *List[T]) Push(v T) {
	newNode := &Node[T]{Value: v, next: nil, prev: ll.last}
	if ll.last != nil {
		ll.last.next = newNode
	} else {
		ll.first = newNode
	}
	ll.last = newNode
	ll.size += 1 // maintian list size
}

// Shift pops the element at Head. It returns error if the linked list is empty.
func (ll *List[T]) Shift() (T, error) {
	if ll.size == 0 {
		var zero T
		return zero, errors.New("Empty list")
	}

	n := ll.first
	ll.first = n.next
	if ll.first != nil {
		ll.first.prev = nil
	} else {
		ll.last = nil
	}
	ll.size -= 1 // maintain list size
	return n.Value, nil
}

// Pop pops the element at Tail. It returns error if the linked list is empty.
func (ll *List[T]) Pop() (T, error) {
	if ll.size == 0 {
		var zero T
		return zero, errors.New("Empty list")
	}

	n := ll.last
	ll.last = n.Prev()
	if ll.last != nil {
		ll.last.next = nil
	} else {
		ll.first = nil
	}
	ll.size -= 1 // maintain list size
	return n.Value, nil
}

// Reverse reverses the given linked list in-place.
func (ll *List[T]) Reverse() {
	if ll.size == 0 {
		return
	}

	cur := ll.last
	newHead := &Node[T]{Value: cur.Value, next: nil, prev: nil}
	cur = cur.prev
	newCur := newHead
	for cur != nil {
		newNode := &Node[T]{Value: cur.Value, next: nil, prev: newCur}
		newCur.next = newNode
		newCur = newCur.next
		cur = cur.prev
	}

	ll.first = newHead
	ll.last = newCur
}

// First returns a pointer to the first node in the linked list (head).
func (ll *List[T]) First() *Node[T] {
	return ll.first
}

// Last returns a pointer to the last node in the linked list (tail).
func (ll *List[T]) Last() *Node[T] {
	return ll.last
}

// Count returns the number of nodes in the list.
func (ll *List[T]) Count() int {
	return ll.size
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List[T]) Delete(v T) bool {
	if ll.size == 0 {
		return false
	}
	if ll.first.Value == v {
		ll.Shift()
		return true
	}

	cur := ll.first
	prev := cur.prev
	for cur != nil {
		if cur.Value == v {
			if cur.next == nil {
				ll.Pop()
			} else {
				prev.next = cur.next
				cur.next.prev = prev
				ll.size -= 1 // maintain list size
			}
			return true
		}
		prev = cur
		cur = cur.next
	}
	return false
}
