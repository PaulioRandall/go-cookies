package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// RequireFile asserts that a file exists and that it has the expected content.
func RequireFile(t *testing.T, f string, exp string) {
	require.FileExists(t, f)

	bytes, err := ioutil.ReadFile(f)

	if err != nil {
		require.Fail(t, "Unable to read "+f)
	}

	act := string(bytes)
	require.Equal(t, exp, act)
}

// RequireNotExists asserts that a file or directory does NOT exist.
func RequireNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)

	if err == nil {
		require.Fail(t, "File or directory should NOT exist "+f)
	}

	require.True(t, os.IsNotExist(err), "Unable to determine if exists "+f)
}
