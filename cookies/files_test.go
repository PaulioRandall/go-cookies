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

	checkErr := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	requireHistory := func(exps ...string) {
		require.Equal(t, len(exps), len(WorkDirHistory))
		for i, exp := range exps {
			require.Equal(t, exp, WorkDirHistory[i])
		}
	}

	a, e := ioutil.TempDir(temp, "a")
	checkErr(e)
	require.Nil(t, Pushd(a))
	requireHistory(temp)

	b, e := ioutil.TempDir(a, "b")
	checkErr(e)
	require.Nil(t, Pushd(b))
	requireHistory(temp, a)

	c, e := ioutil.TempDir(a, "c")
	checkErr(e)
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