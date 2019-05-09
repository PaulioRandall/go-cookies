
package pkg

import (
	"errors"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// LogIfErr()
// ****************************************************************************

func TestLogIfErr_1(t *testing.T) {
	act := LogIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

func TestLogIfErr_2(t *testing.T) {
	err := errors.New("Computer says no!")
	act := LogIfErr(err)
	assert.True(t, act)
	// Output:
	// [ERROR] Computer says no!
}

// ****************************************************************************
// WarnIfErr()
// ****************************************************************************

func TestWarnIfErr_1(t *testing.T) {
	act := WarnIfErr(nil)
	assert.False(t, act)
	// Output:
	//
}

func TestWarnIfErr_2(t *testing.T) {
	err := errors.New("Computer says no!")
	act := WarnIfErr(err)
	assert.True(t, act)
	// Output:
	// [warning] Computer says no!
}

// ****************************************************************************
// StripWhitespace()
// ****************************************************************************

func TestStripWhitespace_0(t *testing.T) {
	assert.Equal(t, "Rincewind", StripWhitespace("Rince \n\t\f\r wind"),
		"Failed when whitespace in centre of given string")
	assert.Equal(t, "Rincewind", StripWhitespace("\t \n\t \r\n\n\fRincewind"),
		"Failed when whitespace at start of given string")	
	assert.Equal(t, "Rincewind", StripWhitespace("Rincewind\r\n \t\t\f \r  \v\v"),
		"Failed when whitespace at end of given string")	
	assert.Equal(t, "Rincewind", StripWhitespace("\r\nRi \tn\tc\t\t ew\f \r  in\vd\v"),
		"Failed when whitespace at the start, middle, and end of given string")
	assert.Equal(t, "Rincewind", StripWhitespace("Rincewind"),
		"Failed when non-whitespace string given")
	assert.Equal(t, "", StripWhitespace(""),
		"Failed when empty string given")
	assert.Equal(t, "", StripWhitespace("\r\n \t\t \t\t \f \r  \v\v  "),
		"Failed when whitespace only string given")
}

// ****************************************************************************
// IsPositiveInt()
// ****************************************************************************

func TestIsPositiveIntCSV_1(t *testing.T) {
	act := IsPositiveIntCSV("5")
	assert.True(t, act)
}

func TestIsPositiveIntCSV_2(t *testing.T) {
	act := IsPositiveIntCSV("1,2,3,4")
	assert.True(t, act)
}

func TestIsPositiveIntCSV_3(t *testing.T) {
	act := IsPositiveIntCSV("")
	assert.False(t, act)
}

func TestIsPositiveIntCSV_4(t *testing.T) {
	act := IsPositiveIntCSV("abc")
	assert.False(t, act)
}

func TestIsPositiveIntCSV_5(t *testing.T) {
	act := IsPositiveIntCSV("abc,efg,xyz")
	assert.False(t, act)
}

// ****************************************************************************
// ToUnixMilli()
// ****************************************************************************

func TestToUnixMilli_1(t *testing.T) {
	aIn, err := time.Parse(time.RFC3339, "2019-04-15T21:50:33-00:00")
	require.Nil(t, err)
	aOut := ToUnixMilli(aIn)
	assert.Equal(t, int64(1555365033000), aOut)
}