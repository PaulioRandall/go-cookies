package toastify

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertHeaderExists asserts that header 'k' exists within header set 'h'.
// Returns true if all assertions passed.
func AssertHeaderExists(t *testing.T, k string, h http.Header) bool {
	return assert.NotEmpty(t, h.Get(k))
}

// AssertHeaderEqual asserts that header 'k' exists within header set 'h' with
// the value 'exp'. Returns true if all assertions passed.
func AssertHeaderEqual(t *testing.T, k string, h http.Header, exp string) bool {
	return assert.Equal(t, exp, h.Get(k))
}

// AssertHeaderNotEqual asserts that header 'k' either does not exist within
// header set 'h' or that it does NOT equal the value 'notExp'. Returns true if
// all assertions passed.
func AssertHeaderNotEqual(t *testing.T, k string, h http.Header, notExp string) bool {
	return assert.NotEqual(t, notExp, h.Get(k))
}

// AssertHeadersEqual asserts that headers 'h' contains the entries within
// 'exp'. Returns true if all assertions passed.
func AssertHeadersEqual(t *testing.T, h http.Header, exp map[string]string) bool {
	ok := true
	for k, v := range exp {
		ok = ok && AssertHeaderEqual(t, k, h, v)
	}
	return ok
}

// AssertHeaderMatches asserts that 'k' exists in headers 'h' and its values
// matches the regex pattern 'p'. Returns true if all assertions passed.
func AssertHeaderMatches(t *testing.T, k string, h http.Header, p string) bool {
	ok := AssertHeaderExists(t, k, h)
	return ok && assert.Regexp(t, p, h.Get(k))
}

// AssertHeadersMatch asserts that 'h' contains all keys of 'p' and each header
// entry matches the regex pattern under the key in 'p'. Returns true if all
// assertions passed.
func AssertHeadersMatch(t *testing.T, h http.Header, p map[string]string) bool {
	ok := true
	for k, reg := range p {
		ok = ok && AssertHeaderMatches(t, k, h, reg)
	}
	return ok
}
