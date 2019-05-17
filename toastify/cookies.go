package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TempDir creates a temporary directory using 'ioutil.TempDir()' and returns
// the path to it. If there is an error the test will fail and immediately exit.
func TempDir(t *testing.T) string {
	f, err := ioutil.TempDir(".", "")
	require.Nil(t, err)
	return f
}

// RemoveDir removes a directory using 'os.RemoveAll()'. If there is an error
// the test will fail and immediately exit.
func RemoveDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	require.Nil(t, err)
}
