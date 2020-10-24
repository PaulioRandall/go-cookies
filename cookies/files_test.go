package cookies

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileToQuote(t *testing.T) {

	f, e := ioutil.TempFile(".", "*")
	if e != nil {
		log.Fatal(e)
	}
	defer os.Remove(f.Name())

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
