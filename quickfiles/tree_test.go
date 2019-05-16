package quickfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func requireDir(t *testing.T, dir string) {
	_, err := ioutil.ReadDir(dir)

	switch {
	case err == nil:
	case os.IsNotExist(err):
		assert.FailNow(t, dir+"/ was expected but does not exist")
	default:
		assert.FailNow(t, dir+"/ was expected but it's not there or could not be accessed")
	}
}

func randomDir(t *testing.T) string {
	f, err := ioutil.TempDir(".", "")
	require.Nil(t, err)
	return f
}

func removeDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	require.Nil(t, err)
}

func assertFile(t *testing.T, f string, exp string) {
	bytes, err := ioutil.ReadFile(f)
	act := string(bytes)

	switch {
	case err == nil:
		assert.Equal(t, exp, act)
	case os.IsNotExist(err):
		assert.Fail(t, f+" was expected but does not exist")
	default:
		assert.Fail(t, f+" was expected but find, access, or open")
	}
}

func TestCreateParent(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createParent(n + "/parent/file.txt")
	require.Nil(t, err)
	requireDir(t, n+"/parent")
}

func TestCreateDir(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createDir(n)
	require.Nil(t, err)
	requireDir(t, n)
}

func TestCreateFile(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	f := n + "/parent/file.txt"
	d := "Three little piggies"
	data := FileData(d)

	err := createFile(f, data)
	require.Nil(t, err)
	assertFile(t, f, d)
}

func TestIsDir(t *testing.T) {
	assert.True(t, isDir("/"))
	assert.True(t, isDir("abc/"))
	assert.False(t, isDir("abc"))
	assert.False(t, isDir("/abc"))
}

func TestCreateFiles(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	tree := Tree{
		Root: FilePath(n),
		Files: map[FilePath]FileData{
			"temp/abc.txt":        "Weatherwax",
			"temp/xyz.txt":        "Ogg",
			"temp/nested/abc.txt": "Garlick",
			"empty/":              "",
		},
	}

	fmt.Println(tree)
	err := createFiles(tree.Root, tree.Files)
	require.Nil(t, err)

	requireDir(t, n+"/temp")
	assertFile(t, n+"/temp/abc.txt", "Weatherwax")
	assertFile(t, n+"/temp/xyz.txt", "Ogg")
	requireDir(t, n+"/temp/nested")
	assertFile(t, n+"/temp/nested/abc.txt", "Garlick")
	requireDir(t, n+"/empty")
}
