package pkg

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	tree := Tree{
		".",
		[]Leaf{
			Leaf{"abc", "Weatherwax"},
			Leaf{"xyz", "Ogg"},
		},
		[]Branch{
			Branch{
				"nested",
				[]Leaf{
					Leaf{"abc", "Garlick"},
				},
				nil,
			},
		},
	}

	fmt.Println(tree.String())
}
