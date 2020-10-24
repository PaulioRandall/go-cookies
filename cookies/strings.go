package cookies

import (
	"bytes"
	"strings"
	"unicode"
)

// StripSpace removes all white space from a string.
func StripSpace(s string) string {
	var buf bytes.Buffer
	for _, ru := range s {
		if !unicode.IsSpace(ru) {
			buf.WriteRune(ru)
		}
	}
	return buf.String()
}

// Indent creates a prefix of 'n' instances of 'v' which it prepends to each
// line of 's'. Returns 's' unchanged if 'n' is 0 or 'v' is empty. A panic
// occurs if 'n' is negative.
func IndentLines(n int, v string, s string) string {
	if n < 0 {
		panic("Negative indention count passed")
	}

	if n == 0 || v == "" {
		return s
	}

	lines := strings.Split(s, "\n")
	pre := strings.Repeat(v, n)
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
