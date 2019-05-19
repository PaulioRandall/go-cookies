package toastify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertKeys asserts that all 'keys' are present in 'act'. Returns true if all
// assertions passed.
func AssertKeys(t *testing.T, keys []string, act map[string]interface{}) bool {
	ok := true
	for _, k := range keys {
		ok = ok && assert.Contains(t, act, k)
	}
	return ok
}

// AssertExactKeys asserts that all 'keys' are present in 'act' and that no
// other keys exist. Returns true if all assertions passed.
func AssertExactKeys(t *testing.T, keys []string, act map[string]interface{}) bool {
	ok := AssertKeys(t, keys, act)
	return ok && assert.Equal(t, len(keys), len(act))
}
