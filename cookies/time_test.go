package cookies

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestToUnixMilli(t *testing.T) {
	in, e := time.Parse(time.RFC3339, "2019-04-15T21:50:33-00:00")
	require.Nil(t, e)
	out := ToUnixMilli(in)
	require.Equal(t, int64(1555365033000), out)
}
