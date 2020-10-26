package list

import (
	"fmt"
	"strings"
)

type (
	// Thing is used as a placeholder for some other data type.
	Thing string

	// List is a single linked list.
	List struct {
		Head *Node
		Tail *Node // Used for quick appending
		Size int
	}

	// Node is a node in a List.
	Node struct {
		Data Thing
		Next *Node
	}
)

// Zero is an empty Thing.
var Zero = Thing("")

// Empty returns true if the list is empty.
func (l *List) Empty() bool {
	return l.Size == 0
}

// InRange returns true if 'i' is a valid index for List.Get, List.Set and
// List.Insert.
func (l *List) InRange(i int) bool {
	return i >= 0 && i < l.Size
}

// Get returns the item at index 'i'.
func (l *List) Get(i int) (Thing, error) {
	for idx, n := 0, l.Head; n != nil; idx, n = idx+1, n.Next {
		if idx == i {
			return n.Data, nil
		}
	}
	return Zero, l.checkRange(i)
}

// Set sets 't' as the item at index 'i' returning the previously held item.
func (l *List) Set(i int, t Thing) (Thing, error) {
	for idx, n := 0, l.Head; n != nil; idx, n = idx+1, n.Next {
		if idx == i {
			prev := n.Data
			n.Data = t
			return prev, nil
		}
	}
	return Zero, l.checkRange(i)
}

// Prepend prepends 't' to the front of the list.
func (l *List) Prepend(t Thing) {
	n := &Node{Data: t, Next: l.Head}

	l.Head = n
	if l.Size == 0 {
		l.Tail = n
	}

	l.Size++
}

// Append appends 't' to the end of the list.
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

// Insert inserts 't' at index 'i' within the list.
func (l *List) Insert(i int, t Thing) error {

	if e := l.checkRange(i); e != nil {
		return e
	}

	if i == 0 {
		l.Prepend(t)
		return nil
	}

	if i == l.Size {
		l.Append(t)
		return nil
	}

	n := l.Head
	for idx := 0; idx < i; idx++ {
		n = n.Next
	}

	n.Next = &Node{Data: t, Next: n.Next}
	l.Size++
	return nil
}

// IndexOf returns the first instance of the specified item or -1 if the item
// was not in the list.
func (l *List) IndexOf(t Thing) int {
	i := 0
	for n := l.Head; n != nil; n = n.Next {
		if n.Data == t {
			return i
		}
		i++
	}
	return -1
}

// Remove removes the item at index 'i'.
func (l *List) Remove(i int) error {

	if e := l.checkRange(i); e != nil {
		return e
	}

	var prev, n *Node = nil, l.Head
	for idx := 0; idx != i; idx++ {
		n = n.Next
		prev = n
	}

	switch {
	case l.Size == 1:
		l.Head, l.Tail = nil, nil
	case i == 0:
		l.Head = n.Next
	case i == l.Size-1:
		l.Tail = prev
		l.Tail.Next = nil
	default:
		prev.Next = n.Next
	}

	l.Size--
	return nil
}

// Foreach applies function 'f' to each item in the list returning the moment
// a result is false.
func (l *List) Foreach(f func(int, Thing) bool) {
	for i, n := 0, l.Head; n != nil; i, n = i+1, n.Next {
		if !f(i, n.Data) {
			return
		}
	}
}

// Slice returns a range of items from 'begin' (inc) to 'end' (exc) as a slice,
// preserving order.
func (l *List) Slice(begin, end int) ([]Thing, error) {

	switch {
	case begin < 0:
		return nil, fmt.Errorf("Out of range: begin is negative")
	case begin > end:
		return nil, fmt.Errorf("Invalid range: begin exceeds end")
	case end > l.Size:
		return nil, fmt.Errorf("Out of range: end exceeds list size")
	}

	s := make([]Thing, 0, end-begin)

	for i, n := 0, l.Head; i < end; i, n = i+1, n.Next {
		if i >= begin {
			s = append(s, n.Data)
		}
	}

	return s, nil
}

// SliceAll returns all list items as a slice, preserving order.
func (l *List) SliceAll() []Thing {
	s := make([]Thing, 0, l.Size)
	for n := l.Head; n != nil; n = n.Next {
		s = append(s, n.Data)
	}
	return s
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

func (l *List) checkRange(i int) error {
	if i < 0 {
		return fmt.Errorf("Out of range: index is negative")
	}
	if i >= l.Size {
		return fmt.Errorf("Out of range: index matches or exceeds list size")
	}
	return nil
}
