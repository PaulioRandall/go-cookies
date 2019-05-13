package pkg

import (
	"strconv"
)

// MsgList represents a chain of ordered messages by providing the head and tail
// of the list.
type MsgList struct {
	Head *Msg
	Tail *Msg
}

// Msg represents a message in a chain of ordered messages. When used with the
// MsgList type, do not modify the Next value directly as the affect is
// undetermined.
type Msg struct {
	Message string
	Next    *Msg
}

// AddIfEmpty appends 'm' to the MsgList.
func (ml *MsgList) Add(m string) {
	t := &Msg{
		Message: m,
	}

	if ml.Head == nil {
		ml.Head = t
		ml.Tail = ml.Head
		return
	}

	ml.Tail.Next = t
	ml.Tail = t
}

// AddIfEmpty appends 'm' to the MsgList if 's' is empty.
func (ml *MsgList) AddIfEmpty(s string, m string) {
	if s == "" {
		ml.Add(m)
	}
}

// AddIfNotUint appends 'm' if 's' is NOT a positive integer.
func (ml *MsgList) AddIfNotUint(s string, m string) {
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		ml.Add(m)
	}
}

// AddIfNotUintCSV appends 'm' if 's' is NOT a CSV of positive integers.
func (ml *MsgList) AddIfNotUintCSV(s string, m string) {
	if s != "" && !IsUintCSV(s) {
		ml.Add(m)
	}
}
