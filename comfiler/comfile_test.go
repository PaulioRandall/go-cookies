package comfiler

import (
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

func createTestFile(f string, data string) error {
	s := filepath.Dir(f)
	err := os.MkdirAll(s, 0774)
	if err != nil {
		return err
	}

	b := []byte(data)
	return ioutil.WriteFile(f, b, 0774)
}

const abc = `"abc": {
	"a": "Weatherwax",
	"b": "Ogg",
	"c": "Garlick"
}`

const xyz = `"array": [
	"Wyrd sisters",
	"Witches abroad",
	"Lords & ladies"
]`

const tmp = `{
	{{- "\n"}}{{ .Inject "/abc.json" 2}},
	{{- "\n"}}{{ .Inject "/nested/xyz.json" 2}}
}`

func TestCompile(t *testing.T) {
	n := randomDir(t)
	defer removeDir(t, n)

	require.Nil(t, createTestFile(n+"/abc.json", abc))
	require.Nil(t, createTestFile(n+"/nested/xyz.json", xyz))
	require.Nil(t, createTestFile(n+"/template", tmp))

	com := Comfile{
		Template:  n + "/template",
		Resources: n,
	}

	out := n + "/output"
	err := com.Compile(out)
	require.Nil(t, err)
	require.FileExists(t, out)

	exp := `{
		"abc": {
			"a": "Weatherwax",
			"b": "Ogg",
			"c": "Garlick"
		},
		"array": [
			"Wyrd sisters",
			"Witches abroad",
			"Lords & ladies"
		]
}`
	exp = string(exp)

	b, err := ioutil.ReadFile(out)
	require.Nil(t, err)
	act := string(b)

	assert.Equal(t, exp, act)
}
