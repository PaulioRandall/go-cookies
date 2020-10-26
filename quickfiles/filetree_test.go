package quickfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func createTestFile(f string) error {
	s := filepath.Dir(f)
	err := os.MkdirAll(s, 0774)
	if err != nil {
		return err
	}

	b := []byte("")
	return ioutil.WriteFile(f, b, 0774)
}

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

func AssertNotExists(t *testing.T, f string) {
	_, err := os.Stat(f)
	pre := "'" + f + "' "

	if err == nil {
		assert.Fail(t, pre+"File or directory still exists")
		return
	}

	assert.True(t, os.IsNotExist(err), pre+"Unable to determine if exists")
}

// ****************************************************************************
// Tests start here!
// ****************************************************************************
func TestCreateParent(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createParent(n + "/parent/file.txt")
	require.Nil(t, err)
	require.DirExists(t, n+"/parent")
}

func TestCreateDir(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	err := createDir(n)
	require.Nil(t, err)
	require.DirExists(t, n)
}

func TestCreateFile(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	f := n + "/parent/file.txt"
	d := "Three little piggies"
	data := FileData(d)

	err := createFile(f, data)
	require.Nil(t, err)
	AssertFile(t, f, d)
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

	require.DirExists(t, n+"/temp")
	AssertFile(t, n+"/temp/abc.txt", "Weatherwax")
	AssertFile(t, n+"/temp/xyz.txt", "Ogg")
	require.DirExists(t, n+"/temp/nested")
	AssertFile(t, n+"/temp/nested/abc.txt", "Garlick")
	require.DirExists(t, n+"/empty")
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

	require.DirExists(t, n+"/temp")
	AssertNotExists(t, n+"/temp/abc.txt")
	AssertFile(t, n+"/temp/xyz.txt", "")
	require.DirExists(t, n+"/temp/nested")
	AssertNotExists(t, n+"/temp/nested/abc.txt")
	AssertNotExists(t, n+"/empty")
}