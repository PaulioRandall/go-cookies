
package pkg

import (
	"bytes"
	"log"
	"regexp"
	"time"
	"unicode"
)

const (
	POSITIVE_INT_CSV_PATTERN = "^([1-9][0-9]*,)*([1-9][0-9]*)$"
)

// StripWhitespace removes all white space from a string.
func StripWhitespace(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if !unicode.IsSpace(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// ToUnixMilli returns the input Time as Unix milliseconds.
func ToUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// LogIfErr checks if the input error is NOT nil. When true, the error is logged
// and the check result is returned.
func LogIfErr(err error) bool {
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		return true
	}
	return false
}

// WarnIfErr checks if the input error is NOT nil. When true, the error is
// logged as a warning and the check result is returned.
func WarnIfErr(err error) bool {
	if err != nil {
		log.Println("[WARNING] " + err.Error())
		return true
	}
	return false
}

// IsPositiveIntCSV returns true if the input is a CSV of positive integers.
func IsPositiveIntCSV(s string) bool {
	match, _ := regexp.MatchString(POSITIVE_INT_CSV_PATTERN, s)
	return match
}