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

func TestPushd_AND_Popd(t *testing.T) {
	home, temp := startFileTest()
	defer endFileTest(home, temp)

	a, e := ioutil.TempDir(temp, "a")
	if e != nil {
		panic(e)
	}
	println(a)
	e = Pushd(a)
	require.Nil(t, e)
	require.Equal(t, 1, len(WorkDirHistory))
	require.Equal(t, temp, WorkDirHistory[0])

	b, e := ioutil.TempDir(a, "b")
	if e != nil {
		panic(e)
	}
	e = Pushd(b)
	require.Nil(t, e)
	require.Equal(t, 2, len(WorkDirHistory))
	require.Equal(t, temp, WorkDirHistory[0])
	require.Equal(t, a, WorkDirHistory[1])

	c, e := ioutil.TempDir(a, "c")
	if e != nil {
		panic(e)
	}
	e = Pushd(c)
	require.Nil(t, e)
	require.Equal(t, 3, len(WorkDirHistory))
	require.Equal(t, temp, WorkDirHistory[0])
	require.Equal(t, a, WorkDirHistory[1])
	require.Equal(t, b, WorkDirHistory[2])
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
