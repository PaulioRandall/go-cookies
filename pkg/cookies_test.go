package pkg

import (
	"errors"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestLogIfErr(t *testing.T) {
	assert.False(t, LogIfErr(nil))
	err := errors.New("Computer says no!")
	assert.True(t, LogIfErr(err))
}

func TestWarnIfErr(t *testing.T) {
	assert.False(t, WarnIfErr(nil))
	err := errors.New("Computer says no!")
	assert.True(t, WarnIfErr(err))
}

func TestStripWhitespace(t *testing.T) {
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

func TestIsUintCSV(t *testing.T) {
	assert.True(t, IsUintCSV("5"),
		"Failed when single digit given")
	assert.True(t, IsUintCSV("1,2,3,4"),
		"Failed when multi-digit CSV given")
	assert.False(t, IsUintCSV(""),
		"Failed when empty string given")
	assert.False(t, IsUintCSV("a"),
		"Failed when single non-numeric char given")
	assert.False(t, IsUintCSV("abc,efg,xyz"),
		"Failed when multiple non-numeric chars given")
}

func TestToUnixMilli(t *testing.T) {
	aIn, err := time.Parse(time.RFC3339, "2019-04-15T21:50:33-00:00")
	require.Nil(t, err)
	aOut := ToUnixMilli(aIn)
	assert.Equal(t, int64(1555365033000), aOut)
}
