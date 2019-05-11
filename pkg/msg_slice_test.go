package pkg

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// AppendIfEmpty()
// ****************************************************************************

func TestAppendIfEmpty_1(t *testing.T) {
	act := AppendIfEmpty("", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

func TestAppendIfEmpty_2(t *testing.T) {
	act := AppendIfEmpty("", []string{"xyz"}, "abc")
	assert.Len(t, act, 2)
	assert.Contains(t, act, "xyz")
	assert.Contains(t, act, "abc")
}

func TestAppendIfEmpty_3(t *testing.T) {
	act := AppendIfEmpty("NOT-EMPTY", []string{}, "abc")
	assert.Len(t, act, 0)
}

// ****************************************************************************
// AppendIfNotUint()
// ****************************************************************************

func TestAppendIfNotUint___1(t *testing.T) {
	act := AppendIfNotUint("5", []string{}, "abc")
	assert.Len(t, act, 0)
}

func TestAppendIfNotUint___2(t *testing.T) {
	act := AppendIfNotUint("0", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

func TestAppendIfNotUint___3(t *testing.T) {
	act := AppendIfNotUint("-5", []string{}, "abc")
	assert.Len(t, act, 1)
	assert.Contains(t, act, "abc")
}

func TestAppendIfNotUint___4(t *testing.T) {
	act := []string{}
	act = AppendIfNotUint("-1", act, "abc")
	act = AppendIfNotUint("-1", act, "efg")
	act = AppendIfNotUint("-1", act, "hij")
	assert.Len(t, act, 3)
	assert.Contains(t, act, "abc")
	assert.Contains(t, act, "efg")
	assert.Contains(t, act, "hij")
}

// ****************************************************************************
// AppendIfNotUintCSV()
// ****************************************************************************

func TestAppendIfNotUintCSV_1(t *testing.T) {
	act := AppendIfNotUintCSV("1,2,99,4,3", []string{}, "abc")
	assert.Len(t, act, 0)
}

func TestAppendIfNotUintCSV_2(t *testing.T) {
	act := AppendIfNotUintCSV("4", []string{}, "abc")
	assert.Len(t, act, 0)
}

func TestAppendIfNotUintCSV_3(t *testing.T) {
	act := AppendIfNotUintCSV("0", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotUintCSV_4(t *testing.T) {
	act := AppendIfNotUintCSV("-99", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotUintCSV_5(t *testing.T) {
	act := AppendIfNotUintCSV("3,2,1,0", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotUintCSV_6(t *testing.T) {
	act := AppendIfNotUintCSV(",1,2", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}

func TestAppendIfNotUintCSV_7(t *testing.T) {
	act := AppendIfNotUintCSV("", []string{}, "abc")
	require.Len(t, act, 1)
	assert.Equal(t, "abc", act[0])
}
