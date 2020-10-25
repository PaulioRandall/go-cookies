package cookies

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Simplify and clean up

// FilePath represents an absolute or relative file path.
type FilePath string

// FileData represents the data within a file or to populate a file.
type FileData []byte

// Files is a list the files and directories in a FileTree with some content to
// populate them (files only). A FilePath ending in '/' indicates a directory.
type Files map[FilePath]FileData

// FileTree represents the root directory of a tree whereby the name is absolute
// or relative to the working directory.
type FileTree struct {
	Root  FilePath
	Files Files
}

// Create creates the files and directories. Files that already exist are
// ignored and parent directories are created if they are missing.
func (tree *FileTree) Create() error {
	return createFiles(tree.Root, tree.Files)
}

// createFiles creates the files and directories specified within 'files'. Files
// that already exist are ignored and parent directories are created if they are
// missing.
func createFiles(root FilePath, files Files) error {
	base := string(root)

	for fp, data := range files {
		p := string(fp)
		f := filepath.Join(base, p)

		if isDir(p) {
			if e := createDir(f); e != nil {
				return e
			}
			continue
		}

		if e := createFile(f, data); e != nil {
			return e
		}
	}

	return nil
}

// deleteFiles deletes the files and directories specified within 'files'.
// Note that deleting a directory also deletes its contents.
func deleteFiles(root FilePath, files Files) error {
	base := string(root)

	for fp := range files {
		p := string(fp)
		f := filepath.Join(base, p)

		if exists, e := FileExists(f); e != nil {
			return e
		} else if !exists {
			continue
		}

		if isDir(p) {
			if e := os.RemoveAll(f); e != nil {
				return e
			}
			continue
		}

		if e := os.Remove(f); e != nil {
			return e
		}
	}

	return nil
}

// createParent creates the parent directory of a file or directory if they
// don't already exist.
func createParent(dir string) error {
	dir = filepath.Dir(dir)
	return os.MkdirAll(dir, 0774)
}

// createDir creates the directory 'd'. If it already exists this function
// returns without action.
func createDir(d string) error {
	exists, e := FileExists(d)
	if e != nil || exists {
		return e
	}
	return os.MkdirAll(d, 0774)
}

// createFile creates the file with the specified 'FileData'. If it already
// exists this function returns without action.
func createFile(f string, d FileData) error {
	exists, e := FileExists(f)
	if e != nil || exists {
		return e
	}

	if e = createParent(f); e != nil {
		return e
	}

	b := []byte(d)
	return ioutil.WriteFile(f, b, 0774)
}

// isDir returns true if the FilePath is a directory.
func isDir(f string) bool {
	return strings.HasSuffix(f, "/")
}

// String returns the string representation of the tree.
func (tree FileTree) String() string {
	sb := &strings.Builder{}

	sb.WriteString(string(tree.Root))
	sb.WriteRune('/')

	for k, _ := range tree.Files {
		sb.WriteString("\n\t")
		sb.WriteString(string(k))
	}

	return sb.String()
}
