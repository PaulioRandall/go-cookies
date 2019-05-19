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

// RequireHeaderValue asserts that header 'k' exists within header set 'h' with
// the value 'exp'.
func RequireHeaderValue(t *testing.T, h http.Header, k string, exp string) {
	require.Equal(t, exp, h.Get(k))
}
