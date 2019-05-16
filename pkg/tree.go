package pkg

import (
	"io/ioutil"
	"os"
	"strings"
)

// Leaf represents a regular file (non-directory) within a tree of files
// and the contents of the file.
type Leaf struct {
	Name    string
	Content string
}

// Branch represents a directory within a tree of files.
type Branch struct {
	Name     string
	Leaves   []Leaf
	Branches []Branch
}

// Tree represents the root branch of a tree whereby the name is absolute or
// relative to the working directory.
type Tree Branch

// CreateTree writes the tree of files with the assumption to root file aready
// exits. Files that already exist are ignored unless 'clean' is true; in which
// case those files are overwritten.
func (tree *Tree) CreateTree(clean bool) error {
	err := tree.createLeaves(tree.Name, tree.Leaves, clean)
	if err == nil {
		err = tree.createBranches(tree.Name, tree.Branches, clean)
	}
	return err
}

// fileExists returns true if the provided file exists and an error if it could
// be determined.
func (tree *Tree) fileExists(f string) (bool, error) {
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

// createLeaves is used by 'CreateTree()' to create regular files. Files that
// already exist are ignored unless 'clean' is true; in which case those files
// are overwritten.
func (tree *Tree) createLeaves(root string, l []Leaf, clean bool) error {
	for _, v := range l {
		f := root + "/" + strings.TrimPrefix(v.Name, "/")
		exists, err := tree.fileExists(f)

		switch {
		case err != nil:
			return err

		case !clean && exists:
			continue
		}

		bytes := []byte(v.Content)
		err = ioutil.WriteFile(f, bytes, 0774)
		if err != nil {
			return err
		}
	}

	return nil
}

// createBranches is used by 'CreateTree()' to create directories. Files that
// already exist are ignored unless 'clean' is true; in which case those files
// are overwritten.
func (tree *Tree) createBranches(root string, b []Branch, clean bool) error {
	for _, v := range b {
		f := root + "/" + strings.TrimPrefix(v.Name, "/")
		exists, err := tree.fileExists(f)

		switch {
		case err != nil:
			return err

		case !clean && exists:
			continue

		case exists:
			err = os.RemoveAll(f)
			if err != nil {
				return err
			}
		}

		err = os.Mkdir(f, 0774)
		if err != nil {
			return err
		}

		tree.createLeaves(f, v.Leaves, clean)
		tree.createBranches(f, v.Branches, clean)
	}

	return nil
}

// String returns the string representation of the tree.
func (tree *Tree) String() string {
	sb := &strings.Builder{}
	sb.WriteString(tree.Name)
	tree.stringLeaves(sb, tree.Leaves, 1)
	tree.stringBranches(sb, tree.Branches, 1)
	return sb.String()
}

// stringLeaves is used by 'String()' to write details about a slice of leaves
// to a supplied 'strings.Builder'.
func (tree *Tree) stringLeaves(sb *strings.Builder, l []Leaf, indent int) {
	p := strings.Repeat("\t", indent)
	for _, v := range l {
		sb.WriteRune('\n')
		sb.WriteString(p)
		sb.WriteString("F: ")
		sb.WriteString(v.Name)
	}
}

// stringBranches is used by 'String()' to write details about a slice of
// branches to a supplied 'strings.Builder'.
func (tree *Tree) stringBranches(sb *strings.Builder, b []Branch, indent int) {
	p := strings.Repeat("\t", indent)
	for _, v := range b {
		sb.WriteRune('\n')
		sb.WriteString(p)
		sb.WriteString("B: ")
		sb.WriteString(v.Name)
		tree.stringLeaves(sb, v.Leaves, indent+1)
		tree.stringBranches(sb, v.Branches, indent+1)
	}
}
