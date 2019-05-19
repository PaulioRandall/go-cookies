package toastify

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// RequireHeaderExists asserts that header 'k' exists within header set 'h'.
func RequireHeaderExists(t *testing.T, k string, h http.Header) {
	require.NotEmpty(t, h.Get(k))
}

// RequireHeaderEqual asserts that header 'k' exists within header set 'h' with
// the value 'exp'.
func RequireHeaderEqual(t *testing.T, k string, h http.Header, exp string) {
	require.Equal(t, exp, h.Get(k))
}

// RequireHeaderNotEqual asserts that header 'k' either does not exist within
// header set 'h' or that it does NOT equal the value 'notExp'.
func RequireHeaderNotEqual(t *testing.T, k string, h http.Header, notExp string) {
	require.NotEqual(t, notExp, h.Get(k))
}

// RequireHeadersEqual asserts that headers 'h' contains the entries within
// 'exp'.
func RequireHeadersEqual(t *testing.T, h http.Header, exp map[string]string) {
	for k, v := range exp {
		RequireHeaderEqual(t, k, h, v)
	}
}

// RequireHeaderMatches asserts that 'k' exists in headers 'h' and its values
// matches the regex pattern 'p'.
func RequireHeaderMatches(t *testing.T, k string, h http.Header, p string) {
	RequireHeaderExists(t, k, h)
	require.Regexp(t, p, h.Get(k))
}

// RequireHeadersMatch asserts that 'h' contains all keys of 'p' and each header
// entry matches the regex pattern under the key in 'p'.
func RequireHeadersMatch(t *testing.T, h http.Header, p map[string]string) {
	for k, reg := range p {
		RequireHeaderMatches(t, k, h, reg)
	}
}
