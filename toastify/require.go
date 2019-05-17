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
		pre := "'" + f + "' "
		require.Fail(t, pre+"Unable to determine if exists")
		return
	}

	act := string(bytes)
	require.Equal(t, exp, act)
}

// RequireNotExists asserts that a file or directory does NOT exist.
func RequireNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)
	pre := "'" + f + "' "

	if err == nil {
		require.Fail(t, pre+"Should NOT exist")
		return
	}

	require.True(t, os.IsNotExist(err), pre+"Unable to determine if exists")
}
