package list

import (
	"fmt"
	"strings"
)

// Thing is used as a placeholder for some other data type.
type Thing string

// List is a single linked list.
type List struct {
	Head *Node
	Tail *Node // Used for quick appending
	Size int
}

// Node is a node in a List.
type Node struct {
	Data Thing
	Next *Node
}

// Empty returns true if the list is empty.
func (l *List) Empty() bool {
	return l.Size == 0
}

// Prepend prepends to the front of the list.
func (l *List) Prepend(t Thing) {
	n := &Node{Data: t, Next: l.Head}

	l.Head = n
	if l.Size == 0 {
		l.Tail = n
	}

	l.Size++
}

// Append appends to the end of the list.
func (l *List) Append(t Thing) {
	n := &Node{Data: t}

	l.Tail = n
	if l.Size == 0 {
		l.Head = n
	} else {
		l.Tail.Next = n
	}

	l.Size++
}

// Remove removes the specified item. True is returned if the item was found
// and removed else false is returned.
func (l *List) Remove(t Thing) bool {

	var prev, n *Node
	for n := l.Head; n != nil; n = n.Next {
		if n.Data == t {
			goto FOUND
		}
		prev = n
	}

	return false

FOUND:
	if n == l.Head {
		l.Head = n.Next
	}

	if n == l.Tail {
		l.Tail = prev
	}

	if prev != nil {
		prev.Next = n.Next
	}

	l.Size--
	return true
}

// Foreach applies the function to each item in the list.
func (l *List) Foreach(f func(Thing)) {
	for n := l.Head; n != nil; n = n.Next {
		f(n.Data)
	}
}

// SliceAll returns all list items as a slice, preserving order.
func (l *List) SliceAll() []Thing {
	r := make([]Thing, l.Size)
	for i, n := 0, l.Head; n != nil; i, n = i+1, n.Next {
		r[i] = n.Data
	}
	return r
}

// String returns the list items as a comma delimited string.
func (l *List) String() string {
	sb := strings.Builder{}

	for n := l.Head; n != nil; n = n.Next {
		s := fmt.Sprintf("%+v", n.Data)
		sb.WriteString(s)

		if n.Next != nil {
			sb.WriteRune(',')
		}
	}

	return sb.String()
}
