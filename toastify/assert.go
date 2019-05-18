package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFile asserts that file 'f' exists and that it contains 'exp'. Returns
// true if all assertions passed.
func AssertFile(t *testing.T, f string, exp string) bool {
	ok := assert.FileExists(t, f)
	if !ok {
		return false
	}

	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		assert.Fail(t, "Unable to read file "+f)
		return false
	}

	act := string(bytes)
	return assert.Equal(t, exp, act)
}

// assertDir asserts that a directory 'dir' exists and contains the expected
// 'filenames'. Returns the list of files within the directory and a bool which
// is true if all assertions passed.
func assertDir(t *testing.T, dir string, filenames []string) ([]os.FileInfo, bool) {
	ok := assert.DirExists(t, dir)
	if !ok {
		return nil, false
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		assert.Fail(t, "Unable to read directory "+dir)
		return nil, false
	}

	m := make([]string, len(files))
	for i, f := range files {
		m[i] = f.Name()
	}

	for _, fn := range filenames {
		ok = ok && assert.Contains(t, m, fn)
	}

	return files, ok
}

// AssertDir asserts that a directory 'dir' exists and contains the expected
// 'filenames'. Returns true if all assertions passed.
func AssertDir(t *testing.T, dir string, filenames []string) bool {
	_, ok := assertDir(t, dir, filenames)
	return ok
}

// AssertStrictDir asserts that a directory 'dir' exists and contains the
// expected 'filenames' and no other files. Returns true if all assertions
// passed.
func AssertStrictDir(t *testing.T, dir string, filenames []string) bool {
	files, ok := assertDir(t, dir, filenames)
	ok = ok && assert.Equal(t, len(filenames), len(files))
	return ok
}

// AssertNotExists asserts that a file or directory does NOT exist. Returns true
// if all assertions passed.
func AssertNotExists(t *testing.T, f string) bool {
	_, err := os.Stat(f)
	pre := "'" + f + "' "

	if err == nil {
		assert.Fail(t, pre+"File or directory still exists")
		return false
	}

	return assert.True(t, os.IsNotExist(err), pre+"Unable to determine if exists")
}
