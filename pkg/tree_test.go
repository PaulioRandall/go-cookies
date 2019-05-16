package pkg

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
		assert.FailNow(t, dir+"/ was expected but find, access, or open")
	}
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

func TestCreateTree(t *testing.T) {
	tree := Tree{
		".",
		nil,
		[]Branch{
			Branch{
				"temp/",
				[]Leaf{
					Leaf{"abc.txt", "Weatherwax"},
					Leaf{"xyz.txt", "Ogg"},
				},
				[]Branch{
					Branch{
						"nested/",
						[]Leaf{
							Leaf{"abc.txt", "Garlick"},
						},
						nil,
					},
				},
			},
		},
	}

	fmt.Println(tree.String())
	tree.CreateTree(true)

	requireDir(t, "./temp")
	assertFile(t, "./temp/abc.txt", "Weatherwax")
	assertFile(t, "./temp/xyz.txt", "Ogg")
	requireDir(t, "./temp/nested")
	assertFile(t, "./temp/nested/abc.txt", "Garlick")

	err := os.RemoveAll("./temp")
	require.Nil(t, err)
}
