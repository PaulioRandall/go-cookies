package cookies

import (
	"fmt"
	"time"
)

// ToUnixMilli returns the input Time as Unix milliseconds.
func ToUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// FmtDuration returns the duration as a string using 'dp' to specify
// decimal points and 'radix' to specify units. Unit suffix is added only if it
// matches a metric time unit between (inclusive) nanoseconds and hours.
//
// Print in milliseconds:
// ```
// DurationString(time.Millisecond)
// ```
//
// Custom radix in relation to time.Nanosecond:
// ```
// // tens of microseconds
// DurationString(10 * 1000 * 1000)
// ```
func FmtDuration(t time.Duration, dp uint, radix time.Duration) string {
	switch f := float64(t) / float64(radix); radix {
	case time.Nanosecond:
		return fmt.Sprintf("%.*f ns", dp, f)
	case time.Microsecond:
		return fmt.Sprintf("%.*f us", dp, f)
	case time.Millisecond:
		return fmt.Sprintf("%.*f ms", dp, f)
	case time.Second:
		return fmt.Sprintf("%.*f s", dp, f)
	case time.Minute:
		return fmt.Sprintf("%.*f m", dp, f)
	case time.Hour:
		return fmt.Sprintf("%.*f hr", dp, f)
	default:
		return fmt.Sprintf("%.*f", dp, f)
	}
}
