package pkg

import (
	"strings"
)

// MsgList represents a chain of ordered messages which provides references to
// both its head and tail. Note that it is a single linked-list with the tail
// provided as convenience to the last element.
type MsgList struct {
	Head *Msg
	Tail *Msg
	Size int
}

// Msg represents a message in a chain of ordered messages. Each instance is
// used as a single node within a MsgList. When used within a MsgList, it's
// undetermined what happens when the Next value is modified directly.
type Msg struct {
	Message string
	Next    *Msg
}

// Add appends 'm' to the MsgList.
func (ml *MsgList) Add(m string) {
	t := &Msg{
		Message: m,
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

// ForEach applies the function to each message in the list.
func (ml *MsgList) ForEach(f func(i int, m *Msg)) {
	m := ml.Head
	for i := 0; i < ml.Size; i++ {
		f(i, m)
		m = m.Next
	}
}

// String returns the list of messages as a comma delimited string.
func (ml *MsgList) String() string {
	sb := strings.Builder{}
	ml.ForEach(func(i int, m *Msg) {
		if m != ml.Head {
			sb.WriteString(", ")
		}
		sb.WriteString(m.Message)
	})
	return sb.String()
}

// Slice returns the list of messages as a string slice.
func (ml *MsgList) Slice() []string {
	s := make([]string, ml.Size)
	ml.ForEach(func(i int, m *Msg) {
		s[i] = m.Message
		m = m.Next
	})
	return s
}
