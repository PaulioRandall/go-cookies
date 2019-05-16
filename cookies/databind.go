package cookies

import (
	"io/ioutil"
	"strconv"
)

// fileToQuote returns the bytes of the input file as as a quoted string so it
// may be embedded in source code. Use []byte(quotedString) to decode.
func fileToQuote(varName string, file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	r := strconv.Quote(string(b))
	return r, nil
}
