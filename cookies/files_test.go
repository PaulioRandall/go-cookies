package cookies

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func startFileTest() (home, temp string) {
	var e error
	if home, e = os.Getwd(); e != nil {
		panic(e)
	}
	if temp, e = filepath.Abs("."); e != nil {
		panic(e)
	}
	if temp, e = ioutil.TempDir(temp, "*"); e != nil {
		panic(e)
	}
	if e = os.Chdir(temp); e != nil {
		panic(e)
	}
	return home, temp
}

func endFileTest(home, temp string) {
	if e := os.Chdir(home); e != nil {
		panic(e)
	}
	if e := os.RemoveAll(temp); e != nil {
		panic(e)
	}
}

func requireFile(t *testing.T, f string, exp string) {
	require.FileExists(t, f)
	bytes, e := ioutil.ReadFile(f)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, string(bytes))
}

func requireNotExists(t *testing.T, f string) {
	_, e := os.Stat(f)
	require.NotNil(t, e, "File still exists: %s", f)
	require.True(t, os.IsNotExist(e), "File still exists: %s", f)
}

func TestPushd_AND_Popd(t *testing.T) {
	home, temp := startFileTest()
	defer endFileTest(home, temp)

	tempDir := func(dir string) string {
		r, e := ioutil.TempDir(temp, dir)
		if e != nil {
			panic(e)
		}
		return r
	}

	requireHistory := func(exps ...string) {
		require.Equal(t, len(exps), len(WorkDirHistory))
		for i, exp := range exps {
			require.Equal(t, exp, WorkDirHistory[i])
		}
	}

	a := tempDir("a")
	require.Nil(t, Pushd(a))
	requireHistory(temp)

	b := tempDir("b")
	require.Nil(t, Pushd(b))
	requireHistory(temp, a)

	c := tempDir("c")
	require.Nil(t, Pushd(c))
	requireHistory(temp, a, b)

	require.Nil(t, Popd())
	requireHistory(temp, a)

	require.Nil(t, Popd())
	requireHistory(temp)

	require.Nil(t, Popd())
	requireHistory()
}

func TestFileToQuote(t *testing.T) {
	home, temp := startFileTest()
	defer endFileTest(home, temp)

	f, e := ioutil.TempFile(temp, "*")
	if e != nil {
		log.Fatal(e)
	}

	data := []byte("What you see is all there is.")
	if _, e = f.Write(data); e != nil {
		log.Fatal(e)
	}

	a, e := FileToQuote(f.Name())
	require.Nil(t, e)

	exp := []byte("\"What you see is all there is.\"")
	act := []byte(a)
	require.Equal(t, exp, act)
}

func TestCreateFiles(t *testing.T) {
	home, temp := startFileTest()
	defer endFileTest(home, temp)

	e := CreateFiles(temp, os.ModePerm, map[string][]byte{
		"abc.txt":        []byte("Weatherwax"),
		"xyz.txt":        []byte("Ogg"),
		"nested/abc.txt": []byte("Garlick"),
		"empty/":         nil,
	})
	require.Nil(t, e)

	requireFile(t, temp+"/abc.txt", "Weatherwax")
	requireFile(t, temp+"/xyz.txt", "Ogg")
	require.DirExists(t, temp+"/nested")
	requireFile(t, temp+"/nested/abc.txt", "Garlick")
	require.DirExists(t, temp+"/empty")
}
