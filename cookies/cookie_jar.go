package cookies

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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

// IsUint returns true if the input is an unsigned integer. Note that false will
// be returned if the string cannot be parsed to an integer.
func IsUint(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		return false
	}
	return true
}

// IsUintCSV returns true if the input is a CSV of positive integers.
func IsUintCSV(s string) bool {
	match, _ := regexp.MatchString("^([1-9][0-9]*,)*([1-9][0-9]*)$", s)
	return match
}

// Indent creates a prefix of 'n' instances of 'p' which it prepends to each
// line of 's'. Returns 's' unchanged if 'n' is 0 or p is empty but panics if
// 'n' is negative.
func Indent(n int, p string, s string) string {
	if n < 0 {
		panic("'n', the number of prefix instances, must not be negative")
	}

	if n == 0 || p == "" {
		return s
	}

	lines := strings.Split(s, "\n")
	pre := strings.Repeat(p, n)
	sb := strings.Builder{}

	for i, l := range lines {
		if i != 0 {
			sb.WriteRune('\n')
		}

		sb.WriteString(pre)
		sb.WriteString(l)
	}

	return sb.String()
}

// TrimPrefixSpace trims the whitespace from the beginning of 's'.
func TrimPrefixSpace(s string) string {
	for i, v := range s {
		if !unicode.IsSpace(v) {
			return s[i:]
		}
	}
	return ""
}

// ForEachToken applies to each token within 's', that is delimited by 'sep',
// the function 'f'. The modified string is then returned.
//
// 'sep' and tokenisation behave exactly the as if calling
// 'strings.Split(s, sep)'.
func ForEachToken(s string, sep string, f func(i int, l string) string) string {
	tokens := strings.Split(s, sep)
	for i, l := range tokens {
		tokens[i] = f(i, l)
	}
	return strings.Join(tokens, sep)
}