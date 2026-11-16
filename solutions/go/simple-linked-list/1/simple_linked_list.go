package simplelinkedlist

import (
	"errors"
	"slices"
)

// Define the List and Element types here.
type Element struct {
	data int
	next *Element
}

type List struct {
	first *Element
	size  int
}

func New(elements []int) *List {
	dummy := &Element{}
	cursor := dummy
	for _, elem := range elements {
		cursor.next = &Element{data: elem}
		cursor = cursor.next
	}
	return &List{first: dummy.next, size: len(elements)}
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	if l.size == 0 {
		l.first = &Element{data: element}
		l.size++
		return
	}

	cursor := l.first
	for cursor.next != nil {
		cursor = cursor.next
	}
	cursor.next = &Element{data: element}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("list is empty")
	}

	var elem *Element
	if l.size == 1 {
		elem = l.first
		l.first = nil
		l.size--
		return elem.data, nil
	}

	cursor := l.first
	for cursor.next.next != nil {
		cursor = cursor.next
	}
	elem = cursor.next
	cursor.next = nil
	l.size--
	return elem.data, nil
}

func (l *List) Peek() (int, error) {
	if l.size == 0 {
		return 0, errors.New("list is empty")
	}
	cursor := l.first
	for cursor.next != nil {
		cursor = cursor.next
	}
	return cursor.data, nil
}

func (l *List) Array() []int {
	output := make([]int, 0, l.size)
	cursor := l.first
	for cursor != nil {
		output = append(output, cursor.data)
		cursor = cursor.next
	}
	return output
}

func (l *List) Reverse() *List {
	arr := l.Array()
	slices.Reverse(arr)
	return New(arr)
}
