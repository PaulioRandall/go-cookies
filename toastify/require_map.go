package toastify

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// RequireKeys asserts that all 'keys' are present in 'act'.
func RequireKeys(t *testing.T, keys []string, act map[string]interface{}) {
	for _, k := range keys {
		require.Contains(t, act, k)
	}
}

// RequireExactKeys asserts that all 'keys' are present in 'act' and that no
// other keys exist.
func RequireExactKeys(t *testing.T, keys []string, act map[string]interface{}) {
	RequireKeys(t, keys, act)
	require.Equal(t, len(keys), len(act))
}
