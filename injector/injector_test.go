package injector

import (
	"io/ioutil"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"

	q "github.com/PaulioRandall/go-cookies/quickfiles"
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

	tree := q.Tree{
		Root: q.FilePath(n),
		Files: map[q.FilePath]q.FileData{
			"abc.json":        abc,
			"nested/xyz.json": xyz,
			"template":        tmp,
		},
	}

	tree.CreateFiles()

	inj := Injector{
		Template:  n + "/template",
		Resources: n,
	}

	out := n + "/output"
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
