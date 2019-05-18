package cookies

import (
	"errors"
	"fmt"
	"strconv"
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

func TestToUnixMilli(t *testing.T) {
	aIn, err := time.Parse(time.RFC3339, "2019-04-15T21:50:33-00:00")
	require.Nil(t, err)
	aOut := ToUnixMilli(aIn)
	assert.Equal(t, int64(1555365033000), aOut)
}

func TestIsUint(t *testing.T) {
	assert.True(t, IsUint("9"), "Failed when valid unsigned integer given")
	assert.False(t, IsUint("1.1"), "Failed when float given")
	assert.False(t, IsUint(""), "Failed when empty string given")
	assert.False(t, IsUint("a"), "Failed when non-numeric char given")
}

func TestIsUintCSV(t *testing.T) {
	assert.True(t, IsUintCSV("5"), "Failed when single digit given")
	assert.True(t, IsUintCSV("1,2,3,4"), "Failed when multi-digit CSV given")
	assert.False(t, IsUintCSV(""), "Failed when empty string given")
	assert.False(t, IsUintCSV("a"), "Failed when single non-numeric char given")
	assert.False(t, IsUintCSV("abc,efg,xyz"),
		"Failed when multiple non-numeric chars given")
}

func TestIndent(t *testing.T) {
	exp := "\t\t\n\t\tMoonglow\n\t\tMoonglow\n\t\t"
	act := Indent(2, "\t", "\nMoonglow\nMoonglow\n")
	assert.Equal(t, exp, act)

	assert.Equal(t, "Moonglow", Indent(1, "", "Moonglow"))
	assert.Equal(t, "Moonglow", Indent(0, "\t", "Moonglow"))
	assert.Equal(t, "\t", Indent(1, "\t", ""))

	assert.Panics(t, func() {
		Indent(-5, "\t", "Moonglow")
	})
}

func TestTrimPrefixSpace(t *testing.T) {
	assert.Equal(t, "abc", TrimPrefixSpace(" \n\r\t\vabc"))
	assert.Equal(t, "", TrimPrefixSpace(" \n\r\t\v"))
	assert.Equal(t, "abc \n\r\t\v", TrimPrefixSpace("abc \n\r\t\v"))
	assert.Equal(t, "", TrimPrefixSpace(""))
}

func TestForEachToken(t *testing.T) {
	f := func(i int, l string) string {
		switch i {
		case 0:
			assert.Equal(t, "a", l)
		case 1:
			assert.Equal(t, "b", l)
		default:
			assert.Fail(t, "Only expected 2 tokens")
		}
		return fmt.Sprintf("%s%d", l, i)
	}

	a := ForEachToken("a\nb", "\n", f)
	assert.Equal(t, "a0\nb1", a)

	b := ForEachToken("a", "\n", f)
	assert.Equal(t, "a0", b)

	c := ForEachToken("ab", "", f)
	assert.Equal(t, "a0b1", c)
}

func TestMapStrings(t *testing.T) {
	f := func(i int, s string) (string, string) {
		switch i {
		case 0:
			assert.Equal(t, "a", s)
		case 1:
			assert.Equal(t, "b", s)
		default:
			assert.Fail(t, "Only expected 2 items at most")
		}
		return s, strconv.Itoa(i)
	}

	aAct := MapStrings([]string{"a", "b"}, f)
	aExp := map[string]string{"a": "0", "b": "1"}
	assert.Equal(t, aExp, aAct)

	bAct := MapStrings([]string{"a"}, f)
	bExp := map[string]string{"a": "0"}
	assert.Equal(t, bExp, bAct)
}
