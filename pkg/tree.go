package pkg

import "strings"

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

// Create writes the tree of files. Files that already exist are ignored unless
// 'clean' is true; in which case those files are overwritten.
func (tree *Tree) Create(clean bool) {

}

// String returns the string representation of the tree.
func (tree *Tree) String() string {
	sb := &strings.Builder{}
	sb.WriteString(tree.Name)
	tree.writeLeaves(sb, tree.Leaves, 1)
	tree.writebranches(sb, tree.Branches, 1)
	return sb.String()
}

// writeLeaves is used by 'String()' to write details about a slice of leaves
// to a supplied 'strings.Builder'.
func (tree *Tree) writeLeaves(sb *strings.Builder, l []Leaf, indent int) {
	p := strings.Repeat("\t", indent)
	for _, v := range l {
		sb.WriteRune('\n')
		sb.WriteString(p)
		sb.WriteString("F: ")
		sb.WriteString(v.Name)
	}
}

// writebranches is used by 'String()' to write details about a slice of
// branches to a supplied 'strings.Builder'.
func (tree *Tree) writebranches(sb *strings.Builder, b []Branch, indent int) {
	p := strings.Repeat("\t", indent)
	for _, v := range b {
		sb.WriteRune('\n')
		sb.WriteString(p)
		sb.WriteString("B: ")
		sb.WriteString(v.Name)
		tree.writeLeaves(sb, v.Leaves, indent+1)
		tree.writebranches(sb, v.Branches, indent+1)
	}
}
