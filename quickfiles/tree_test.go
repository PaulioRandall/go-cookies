package quickfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func randomDir(t *testing.T) string {
	f, err := ioutil.TempDir(".", "")
	require.Nil(t, err)
	return f
}

func removeDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	require.Nil(t, err)
}

func requireDirExists(t *testing.T, dir string) {
	_, err := ioutil.ReadDir(dir)

	switch {
	case err == nil:
	case os.IsNotExist(err):
		assert.FailNow(t, dir+"/ was expected but does not exist")
	default:
		assert.FailNow(t, dir+"/ was expected but it's not there or could not be accessed")
	}
}

func assertFileExists(t *testing.T, f string, exp string) {
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

func assertNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)
	switch {
	case err == nil:
		assert.Fail(t, "File or directory still exists")
	case os.IsNotExist(err):
	default:
		assert.Fail(t, "File or directory existence could not be checked")
	}
}

func createTestFile(f string) error {
	s := filepath.Dir(f)
	err := os.MkdirAll(s, 0774)
	if err != nil {
		return err
	}

	b := []byte("")
	return ioutil.WriteFile(f, b, 0774)
}

// ****************************************************************************
// Tests start here!
// ****************************************************************************
func TestCreateParent(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createParent(n + "/parent/file.txt")
	require.Nil(t, err)
	requireDirExists(t, n+"/parent")
}

func TestCreateDir(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createDir(n)
	require.Nil(t, err)
	requireDirExists(t, n)
}

func TestCreateFile(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	f := n + "/parent/file.txt"
	d := "Three little piggies"
	data := FileData(d)

	err := createFile(f, data)
	require.Nil(t, err)
	assertFileExists(t, f, d)
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

	requireDirExists(t, n+"/temp")
	assertFileExists(t, n+"/temp/abc.txt", "Weatherwax")
	assertFileExists(t, n+"/temp/xyz.txt", "Ogg")
	requireDirExists(t, n+"/temp/nested")
	assertFileExists(t, n+"/temp/nested/abc.txt", "Garlick")
	requireDirExists(t, n+"/empty")
}

func TestDeleteFiles(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	require.Nil(t, createTestFile(n+"/temp/abc.txt"))
	require.Nil(t, createTestFile(n+"/temp/xyz.txt"))
	require.Nil(t, createTestFile(n+"/temp/nested/abc.txt"))
	require.Nil(t, os.MkdirAll(n+"/empty/", 0774))

	tree := Tree{
		Root: FilePath(n),
		Files: map[FilePath]FileData{
			"temp/abc.txt":        "",
			"temp/nested/abc.txt": "",
			"empty/":              "",
		},
	}

	t.Log(tree)
	err := deleteFiles(tree.Root, tree.Files)
	require.Nil(t, err)

	requireDirExists(t, n+"/temp")
	assertNotExists(t, n+"/temp/abc.txt")
	assertFileExists(t, n+"/temp/xyz.txt", "")
	requireDirExists(t, n+"/temp/nested")
	assertNotExists(t, n+"/temp/nested/abc.txt")
	assertNotExists(t, n+"/empty")
}
