package cookies

import (
	"fmt"

	"github.com/pkg/errors"
)

// Wrap wraps an error 'e' with a another message 'm'.
func Wrap(e error, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	return errors.Wrap(e, m)
}
