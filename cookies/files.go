package cookies

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// WorkDirStack is the working directory stack used by Pushd and Popd functions.
var WorkDirStack = []string{}

// Wrap wraps an error 'e' with a another message 'm'.
func Wrap(e error, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	return errors.Wrap(e, m)
}

// CD is an alias for os.Chdir.
func CD(dir string) error {
	return os.Chdir(dir)
}

// Pushd emulates pushd bash functionality.
func Pushd(dir string) error {
	curr, e := os.Getwd()
	if e != nil {
		return e
	}
	if e = os.Chdir(dir); e != nil {
		return e
	}
	WorkDirStack = append(WorkDirStack, curr)
	return nil
}

// Popd emulates the popd bash functionality.
func Popd() error {
	last := len(WorkDirStack) - 1
	dir := WorkDirStack[last]
	if e := os.Chdir(dir); e != nil {
		return e
	}
	WorkDirStack = WorkDirStack[:last]
	return nil
}

// RemoveDir recursively removes directory 'dir'. If something was removed then
// true is returned rather than an error.
func RemoveDir(dir string) (bool, error) {
	switch _, e := os.Stat(dir); {
	case os.IsNotExist(e):
		return false, nil
	case e != nil:
		return false, Wrap(e, "Could not access %s", dir)
	case os.RemoveAll(dir) != nil:
		return false, Wrap(e, "Unable to remove %s", dir)
	default:
		return true, nil
	}
}

// FileExists returns true if the file exists, false if not, and an error if
// file existence could not be determined.
func FileExists(f string) (bool, error) {
	_, e := os.Stat(f)
	if os.IsNotExist(e) {
		return false, nil
	}
	return e == nil, e
}

// IsDir returns true if the file exists and is a directory. An error is
// returned if this could not be determined.
func IsDir(f string) (bool, error) {
	stat, e := os.Stat(f)
	if os.IsNotExist(e) {
		return false, nil
	}
	if e != nil {
		return false, e
	}
	return stat.Mode().IsDir(), nil
}

// IsRegFile returns true if the file exists and is a regular file. An error is
// returned if this could not be determined.
func IsRegFile(f string) (bool, error) {
	stat, e := os.Stat(f)
	if os.IsNotExist(e) {
		return false, nil
	}
	if e != nil {
		return false, e
	}
	return stat.Mode().IsRegular(), nil
}

// SameFile returns true if the two files 'a' and 'b' describe the same files
// as determined by os.SameFile. An error is returned if the file info could
// not be retreived for either file.
func SameFile(a, b string) (bool, error) {
	aStat, e := os.Stat(a)
	if e != nil {
		return false, e
	}
	bStat, e := os.Stat(b)
	if e != nil {
		return false, e
	}
	return os.SameFile(aStat, bStat), nil
}

// CopyFile copies the single file 'src' to 'dst'.
func CopyFile(src, dst string, overwrite bool) error {

	if ok, e := IsRegFile(src); e != nil || !ok {
		return fmt.Errorf("Missing or not a regular file: %s", src)
	}

	if !overwrite {
		ok, e := FileExists(dst)
		if e != nil {
			return e
		}
		if ok {
			return fmt.Errorf("Destination already exists: %s", dst)
		}
	}

	same, e := SameFile(src, dst)
	if e == nil && same {
		return fmt.Errorf("Destination is the same as source: %s == %s", dst, src)
	}

	return NoCheckCopyFile(src, dst)
}

// NoCheckCopyFile copies the single file 'src' to 'dst' and doesn't make any
// attempt to check the file paths before hand.
func NoCheckCopyFile(src, dst string) error {

	srcFile, e := os.Open(src)
	if e != nil {
		return e
	}
	defer srcFile.Close()

	dstFile, e := os.Create(dst)
	if e != nil {
		return e
	}
	defer dstFile.Close()

	_, e = io.Copy(dstFile, srcFile)
	if e != nil {
		return e
	}

	return nil
}

// FileToQuote returns the bytes of the input file as as a quoted string so it
// may be embedded in source code. Use []byte(quotedString) to decode.
func FileToQuote(file string) (string, error) {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	s := strconv.Quote(string(b))
	return s, nil
}
