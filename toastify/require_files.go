package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// RequireFile asserts that file 'f' exists and that it contains 'exp'.
func RequireFile(t *testing.T, f string, exp string) {
	require.FileExists(t, f)

	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		require.Fail(t, "Unable to read file "+f)
	}

	act := string(bytes)
	require.Equal(t, exp, act)
}

// RequireDir asserts that a directory 'dir' exists and contains the expected
// 'filenames'.
func RequireDir(t *testing.T, dir string, filenames []string) {
	requireDir(t, dir, filenames)
}

// RequireExactDir asserts that a directory 'dir' exists and contains the
// expected 'filenames' and no other files.
func RequireExactDir(t *testing.T, dir string, filenames []string) {
	files := requireDir(t, dir, filenames)
	require.Equal(t, len(filenames), len(files))
}

// RequireNotExists asserts that a file or directory does NOT exist.
func RequireNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)
	pre := "'" + f + "' "

	require.NotNil(t, err, pre+"File or directory still exists")
	require.True(t, os.IsNotExist(err), pre+"Unable to determine if exists")
}

// requireDir asserts that a directory 'dir' exists and contains the expected
// 'filenames'. Returns the list of files within the directory.
func requireDir(t *testing.T, dir string, filenames []string) []os.FileInfo {
	require.DirExists(t, dir)

	files, err := ioutil.ReadDir(dir)
	require.Nil(t, err, "Unable to read directory "+dir)

	m := make([]string, len(files))
	for i, f := range files {
		m[i] = f.Name()
	}

	for _, fn := range filenames {
		require.Contains(t, m, fn)
	}

	return files
}
