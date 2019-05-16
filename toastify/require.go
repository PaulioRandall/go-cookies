package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// requireTempDir creates a temporary directory using 'ioutil.TempDir()' and
// returns the path to it. If there is an error the test will fail and
// immediately exit.
func requireTempDir(t *testing.T) string {
	f, err := ioutil.TempDir(".", "")
	require.Nil(t, err)
	return f
}

// requireRemoveDir removes a directory using 'os.RemoveAll()'. If there is an
// error the test will fail and immediately exit.
func requireRemoveDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	require.Nil(t, err)
}
