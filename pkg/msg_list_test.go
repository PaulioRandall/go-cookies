package pkg

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// assertEntries tests that the supplied MsgList contains entries, each
// with the espected Message, in the order supplied.
func assertEntries(t *testing.T, ml MsgList, expMsg ...string) {

	require.NotNil(t, ml.Head)
	require.NotNil(t, ml.Tail)

	n := ml.Head

	for i, v := range expMsg {
		if i != 0 {
			n = n.Next
			require.NotNil(t, n)
		}

		assert.Equal(t, v, n.Message)
	}

	assert.Nil(t, n.Next)
	assert.Equal(t, n, ml.Tail)

	size := 0
	for m := ml.Head; m != nil; m = m.Next {
		size++
	}
	assert.Equal(t, len(expMsg), size)
	assert.Equal(t, size, ml.Size)
}

// ****************************************************************************
// MsgList.Add()
// ****************************************************************************

func TestMsgList_Add_1(t *testing.T) {
	a := MsgList{}
	a.Add("abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_Add_2(t *testing.T) {
	a := MsgList{}
	a.Add("abc")
	a.Add("xyz")
	assertEntries(t, a, "abc", "xyz")
}

// ****************************************************************************
// MsgList.ForEach()
// ****************************************************************************

func TestMsgList_ForEach_1(t *testing.T) {
	a := MsgList{}
	a.Add("abc")
	a.Add("xyz")

	total := 0
	a.ForEach(func(i int, m *Msg) {
		total++

		switch i {
		case 0:
			assert.Equal(t, "abc", m.Message)
		case 1:
			assert.Equal(t, "xyz", m.Message)
		}
	})

	assert.Equal(t, 2, total, "Only expected 2 elements")
}

// ****************************************************************************
// MsgList.String()
// ****************************************************************************

func TestMsgList_String_1(t *testing.T) {
	a := MsgList{}
	a.Add("abc")

	s := a.String()
	assert.Equal(t, "abc", s)
}

func TestMsgList_String_2(t *testing.T) {
	a := MsgList{}
	a.Add("abc")
	a.Add("efg")
	a.Add("xyz")

	s := a.String()
	assert.Equal(t, "abc, efg, xyz", s)
}

// ****************************************************************************
// MsgList.Slice()
// ****************************************************************************

func TestMsgList_Slice_1(t *testing.T) {
	a := MsgList{}
	a.Add("abc")

	s := a.Slice()
	assert.Equal(t, []string{"abc"}, s)
}

func TestMsgList_Slice_2(t *testing.T) {
	a := MsgList{}
	a.Add("abc")
	a.Add("efg")
	a.Add("xyz")

	s := a.Slice()
	assert.Equal(t, []string{"abc", "efg", "xyz"}, s)
}
