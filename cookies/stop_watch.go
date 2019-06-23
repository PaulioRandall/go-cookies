//usr/bin/env go run "$0" "$@"; exit "$?"

package cookies

import (
	"fmt"
	"time"
)

// StopWatch represents a process timer with a few common operations.
type StopWatch struct {
	Started time.Time
	Stopped time.Time
}

// Start starts the stop watch overwritting any previously start time.
func (sw *StopWatch) Start() {
	sw.Started = time.Now().UTC()
}

// Stop stops the stop watch overwriting any previous stop time.
func (sw *StopWatch) Stop() {
	sw.Stopped = time.Now().UTC()
}

// Lap restarts the stop watch but returns a copy of the stop watch with the
// previous laps times.
func (sw *StopWatch) Lap() StopWatch {
	sw.Stop()
	r := StopWatch{
		Started: sw.Started,
		Stopped: sw.Stopped,
	}
	sw.Started = sw.Stopped
	return r
}

// Elapsed returns the elapsed time between the started and stopped times.
func (sw *StopWatch) Elapsed() time.Duration {
	return sw.Stopped.Sub(sw.Started)
}

// PrintElapsed prints the elapsed time as returned by Elapsed(). 'radix' may be
// passed to print the result in a more appropriate manner.
//
// e.g. PrintElapsed(time.Millisecond) will print as milliseconds.
//
// e.g. PrintElapsed(10 * 1000 * 1000) will print using a custom radix,
// microseconds * 10 in this case.
func (sw *StopWatch) PrintElapsed(radix time.Duration) {
	t := sw.Elapsed()
	f := float64(t) / float64(radix)

	switch radix {
	case time.Nanosecond:
		fmt.Printf("%.3f ns\n", f)
	case time.Microsecond:
		fmt.Printf("%.3f us\n", f)
	case time.Millisecond:
		fmt.Printf("%.3f ms\n", f)
	case time.Second:
		fmt.Printf("%.3f s\n", f)
	case time.Minute:
		fmt.Printf("%.3f m\n", f)
	case time.Hour:
		fmt.Printf("%.3f hr\n", f)
	default:
		fmt.Printf("%.3f\n", f)
	}
}
