package strlist

import (
	"strings"
)

// StrList represents a chain of ordered strings which provides references to
// both its head and tail. Note that it's a single linked-list with the tail
// provided as convenience to the last element.
type StrList struct {
	Head *Str
	Tail *Str
	Size int
}

// Str represents a string in a chain of ordered strings. Each instance is
// used as a single node within a StrList. When used within a StrList, it's
// undetermined what happens when the Next value is modified directly so it's
// best to avoid touching it.
type Str struct {
	Value string
	Next  *Str
}

// Add appends 's' to the StrList.
func (ml *StrList) Add(s string) {
	t := &Str{
		Value: s,
	}

	if ml.Head == nil {
		ml.Head = t
		ml.Tail = ml.Head
		ml.Size = 1
		return
	}

	ml.Tail.Next = t
	ml.Tail = t
	ml.Size++
}

// ForEach applies the function to each string in the list.
func (ml *StrList) ForEach(f func(i int, s *Str)) {
	s := ml.Head
	for i := 0; i < ml.Size; i++ {
		f(i, s)
		s = s.Next
	}
}

// Slice returns the list of strings as a string slice.
func (ml *StrList) Slice() []string {
	slice := make([]string, ml.Size)
	ml.ForEach(func(i int, s *Str) {
		slice[i] = s.Value
		s = s.Next
	})
	return slice
}

// String returns the list of strings as a comma delimited string.
func (ml StrList) String() string {
	sb := strings.Builder{}
	ml.ForEach(func(i int, s *Str) {
		if s != ml.Head {
			sb.WriteString(", ")
		}
		sb.WriteString(s.Value)
	})
	return sb.String()
}
