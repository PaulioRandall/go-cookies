package strlist

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// assertEntries tests that the supplied StrList contains entries, each
// with the espected Message, in the order supplied.
func assertEntries(t *testing.T, ml StrList, expStr ...string) {

	require.NotNil(t, ml.Head)
	require.NotNil(t, ml.Tail)

	n := ml.Head

	for i, v := range expStr {
		if i != 0 {
			n = n.Next
			require.NotNil(t, n)
		}

		assert.Equal(t, v, n.Value)
	}

	assert.Nil(t, n.Next)
	assert.Equal(t, n, ml.Tail)

	size := 0
	for m := ml.Head; m != nil; m = m.Next {
		size++
	}
	assert.Equal(t, len(expStr), size)
	assert.Equal(t, size, ml.Size)
}

// ****************************************************************************
// StrList.Add()
// ****************************************************************************

func TestStrList_Add_1(t *testing.T) {
	a := StrList{}
	a.Add("abc")
	assertEntries(t, a, "abc")
}

func TestStrList_Add_2(t *testing.T) {
	a := StrList{}
	a.Add("abc")
	a.Add("xyz")
	assertEntries(t, a, "abc", "xyz")
}

// ****************************************************************************
// StrList.ForEach()
// ****************************************************************************

func TestStrList_ForEach_1(t *testing.T) {
	a := StrList{}
	a.Add("abc")
	a.Add("xyz")

	total := 0
	a.ForEach(func(i int, m *Str) {
		total++

		switch i {
		case 0:
			assert.Equal(t, "abc", m.Value)
		case 1:
			assert.Equal(t, "xyz", m.Value)
		}
	})

	assert.Equal(t, 2, total, "Only expected 2 elements")
}

// ****************************************************************************
// StrList.String()
// ****************************************************************************

func TestStrList_String(t *testing.T) {
	a := StrList{}
	a.Add("abc")
	s := a.String()
	assert.Equal(t, "abc", s)

	a.Add("efg")
	a.Add("xyz")
	s = a.String()
	assert.Equal(t, "abc, efg, xyz", s)
}

// ****************************************************************************
// StrList.Slice()
// ****************************************************************************

func TestStrList_Slice(t *testing.T) {
	a := StrList{}
	a.Add("abc")
	s := a.Slice()
	assert.Equal(t, []string{"abc"}, s)

	a.Add("efg")
	a.Add("xyz")
	s = a.Slice()
	assert.Equal(t, []string{"abc", "efg", "xyz"}, s)
}
