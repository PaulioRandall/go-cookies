
package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_1(t *testing.T) {
	a := add(2, 3)
	assert.Equal(t, 5, a)
}