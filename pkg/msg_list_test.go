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
// MsgList.AddIfEmpty()
// ****************************************************************************

func TestMsgList_AddIfEmpty_1(t *testing.T) {
	a := MsgList{}
	a.AddIfEmpty("", "abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_AddIfEmpty_2(t *testing.T) {
	a := MsgList{}
	a.AddIfEmpty("not empty", "abc")

	assert.Nil(t, a.Head)
	assert.Nil(t, a.Tail)
}

// ****************************************************************************
// MsgList.AddIfNotUint()
// ****************************************************************************

func TestMsgList_AddIfNotUint_1(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUint("not uint", "abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_AddIfNotUint_2(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUint("1.1", "abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_AddIfNotUint_3(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUint("9", "abc")

	assert.Nil(t, a.Head)
	assert.Nil(t, a.Tail)
}

// ****************************************************************************
// MsgList.AddIfNotUintCSV()
// ****************************************************************************

func TestMsgList_AddIfNotUintCSV_1(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUintCSV("not uint csv", "abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_AddIfNotUintCSV_2(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUintCSV("1,2.2,3", "abc")
	assertEntries(t, a, "abc")
}

func TestMsgList_AddIfNotUintCSV_3(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUintCSV("1,2,3", "abc")

	assert.Nil(t, a.Head)
	assert.Nil(t, a.Tail)
}

func TestMsgList_AddIfNotUintCSV_4(t *testing.T) {
	a := MsgList{}
	a.AddIfNotUintCSV("", "abc")

	assert.Nil(t, a.Head)
	assert.Nil(t, a.Tail)
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
