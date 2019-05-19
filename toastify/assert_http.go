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

// AssertHeaderValue asserts that header 'k' exists within header set 'h' with
// the value 'exp'. Returns true if all assertions passed.
func AssertHeaderValue(t *testing.T, h http.Header, k string, exp string) bool {
	return assert.Equal(t, exp, h.Get(k))
}
