package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// createTmpDir creates a directory with the specified 'pattern' in the
// specified 'dir' returning File.Name().
func createTmpDir(t *testing.T, dir, pattern string) string {
	d, err := ioutil.TempDir(dir, pattern)
	require.Nil(t, err, "Error making temp directory")
	return d
}

// createNamedTmpDir creates a directory with the specified 'name' in the
// specified 'dir' returning File.Name().
func createNamedTmpDir(t *testing.T, dir, name string) string {
	n := dir + "/" + name
	err := os.MkdirAll(n, 0777)
	require.Nil(t, err, "Error making named temp directory or it's parents")
	return n
}

// createTmpFile creates a file with the specified 'pattern' in the specified
// 'dir' containing the specified 'content' returning File.Name().
func createTmpFile(t *testing.T, dir, pattern, content string) string {
	f, err := ioutil.TempFile(dir, pattern)
	require.Nil(t, err, "Error creating temp file")
	defer f.Close()
	f.WriteString(content)
	return f.Name()
}

// createNamedTmpFile creates a file with the specified 'name' in the specified
// 'dir' containing the specified 'content' returning File.Name().
func createNamedTmpFile(t *testing.T, dir, name, content string) string {
	err := os.MkdirAll(dir, 0777)
	require.Nil(t, err, "Error making parent directories of named temp file")

	n := dir + "/" + name
	b := []byte(content)
	err = ioutil.WriteFile(n, b, 0777)

	if err != nil {
		fmt.Println(err)
	}

	require.Nil(t, err, "Error creating named temp file")
	return n
}

func TestCompile_1(t *testing.T) {
	d := createTmpDir(t, ".", "")
	defer os.RemoveAll(d)

	createNamedTmpFile(t, d, "abc.json", `"abc": {
	"a": "Weatherwax",
	"b": "Ogg",
	"c": "Garlick"
}`)

	d2 := d + "/nested"
	createNamedTmpFile(t, d2, "xyz.json", `"array": [
	"Wyrd sisters",
	"Witches abroad",
	"Lords & ladies"
]`)

	tmp := createNamedTmpFile(t, d, "template", `{
	{{- "\n"}}{{ .Inject "/abc.json" 2}},
	{{- "\n"}}{{ .Inject "/nested/xyz.json" 2}}
}`)

	inj := Injector{
		Template:  tmp,
		Resources: d,
	}

	out := d + "/output"
	err := inj.Compile(out)
	require.Nil(t, err)

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
