package stack

import (
	"fmt"
	"strings"
)

// Thing is used as a placeholder for some other data type.
type Thing string

// Stack is what it is.
type Stack struct {
	Top  *Node
	Size int
}

// Node is a node in a Stack.
type Node struct {
	Data Thing
	Next *Node
}

// Empty returns true if the stack is empty.
func (s *Stack) Empty() bool {
	return s.Size == 0
}

// Push adds the specified item to the top of the stack.
func (s *Stack) Push(t Thing) {
	s.Top = &Node{Data: t, Next: s.Top}
	s.Size++
}

// Peek returns the top of the stack without removing it. If the stack is not
// empty then the it is returned along with true else the zero value of the
// data type is returned along with false.
func (s *Stack) Peek() (_ Thing, _ bool) {
	if s.Size == 0 {
		return
	}
	return s.Top.Data, true
}

// Pop removes and returns the top of the stack. If the stack is the zero value
// of the data type is returned along with false.
func (s *Stack) Pop() (t Thing, ok bool) {
	if t, ok = s.Peek(); ok {
		s.Top = s.Top.Next
		s.Size--
	}
	return
}

// String returns the stack items as a comma delimited string where the top
// is first and bottom is last.
func (s *Stack) String() string {
	sb := strings.Builder{}

	for n := s.Top; n != nil; n = n.Next {
		s := fmt.Sprintf("%+v", n.Data)
		sb.WriteString(s)

		if n.Next != nil {
			sb.WriteRune(',')
		}
	}

	return sb.String()
}
