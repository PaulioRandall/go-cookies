package quickfiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FilePath represents an absolute or relative file path.
type FilePath string

// FileData represents the data within a file or to populate a file.
type FileData string

// FileSet represents the files and directories in a file tree with some
// content to populate them (files only). A FilePath ending in '/' indicates a
// directory.
type FileSet map[FilePath]FileData

// Tree represents the root directory of a tree whereby the name is absolute or
// relative to the working directory.
type Tree struct {
	Root  FilePath
	Files FileSet
}

// CreateFiles creates the files and directories. Files that already exist are
// ignored and parent directories are created if they are missing.
func (tree Tree) CreateFiles() error {
	return createFiles(tree.Root, tree.Files)
}

// createFiles creates the files and directories specified within 'files'. Files
// that already exist are ignored and parent directories are created if they are
// missing.
func createFiles(root FilePath, files FileSet) error {
	err := error(nil)
	base := string(root)

	for fp, data := range files {
		s := string(fp)
		f := filepath.Join(base, s)

		if isDir(s) {
			err = createDir(f)
		} else {
			err = createFile(f, data)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// deleteFiles deletes the files and directories specified within 'files'.
// Note that deleting a directory also deletes its contents.
func deleteFiles(root FilePath, files FileSet) error {
	err := error(nil)
	base := string(root)

	for fp := range files {
		s := string(fp)
		f := filepath.Join(base, s)

		if isDir(s) {
			err = os.RemoveAll(f)
		} else {
			err = os.Remove(f)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// createParent creates the parent directory of a file or directory if they
// don't already exist.
func createParent(s string) error {
	s = filepath.Dir(s)
	return os.MkdirAll(s, 0774)
}

// createDir creates the directory 'd'. If it already exists this function
// returns without action.
func createDir(d string) error {
	exists, err := exists(d)
	if err != nil || exists {
		return err
	}

	return os.MkdirAll(d, 0774)
}

// createFile creates the file with the specified 'FileData'. If it already
// exists this function returns without action.
func createFile(f string, d FileData) error {
	exists, err := exists(f)
	if err != nil || exists {
		return err
	}

	err = createParent(f)
	if err != nil {
		return err
	}

	b := []byte(d)
	return ioutil.WriteFile(f, b, 0774)
}

// exists returns true if 'f' exists and an error if it could not be determined.
func exists(f string) (bool, error) {
	_, err := os.Stat(f)
	switch {
	case err == nil:
		return true, nil
	case os.IsNotExist(err):
		return false, nil
	default:
		return false, err
	}
}

// isDir returns true if the FilePath is a directory.
func isDir(f string) bool {
	if strings.HasSuffix(f, "/") {
		return true
	}
	return false
}

// String returns the string representation of the FilePath.
func (fp FilePath) String() string {
	return string(fp)
}

// String returns the string representation of the tree.
func (tree Tree) String() string {
	sb := &strings.Builder{}

	sb.WriteString(tree.Root.String())
	sb.WriteRune('/')

	for k, _ := range tree.Files {
		sb.WriteString("\n\t")
		sb.WriteString(k.String())
	}

	return sb.String()
}
