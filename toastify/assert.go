package toastify

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFile asserts that a file exists and that it has the expected content.
func AssertFile(t *testing.T, f string, exp string) {
	assert.FileExists(t, f)

	bytes, err := ioutil.ReadFile(f)

	if err != nil {
		pre := "'" + f + "' "
		assert.Fail(t, pre+"Unable to determine if exists")
		return
	}

	act := string(bytes)
	assert.Equal(t, exp, act)
}

// AssertNotExists asserts that a file or directory does NOT exist.
func AssertNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)
	pre := "'" + f + "' "

	if err == nil {
		assert.Fail(t, pre+"File or directory still exists")
		return
	}

	assert.True(t, os.IsNotExist(err), pre+"Unable to determine if exists")
}
