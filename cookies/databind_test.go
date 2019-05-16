package cookies

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestFileToQuote(t *testing.T) {
	f, err := ioutil.TempFile(".", "*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	data := []byte("What you see is all there is.")
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	a, err := fileToQuote("abc", f.Name())
	require.Nil(t, err, "Did not expect error to be returned")

	expect := []byte("\"What you see is all there is.\"")
	actual := []byte(a)
	assert.Equal(t, expect, actual)
}
