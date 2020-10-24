//usr/bin/env go run "$0" "$@"; exit "$?"

package cookies

import (
	"time"

	"github.com/PaulioRandall/go-cookies/cookies"
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

// Stop stops the stop watch overwriting any previous stop time. A copy of the
// stopwatch is returned.
func (sw *StopWatch) Stop() StopWatch {
	sw.Stopped = time.Now().UTC()
	return *sw
}

// Lap stops the stopwatch, copies it, then sets the start as the copies stop
// time. The copy is returned.
func (sw *StopWatch) Lap() StopWatch {
	r := sw.Stop()
	sw.Started = r.Stopped
	return r
}

// Elapsed returns the elapsed time between the started and stopped times.
func (sw *StopWatch) Elapsed() time.Duration {
	return sw.Stopped.Sub(sw.Started)
}

// ElapsedString returns StopWatch.Elapsed as a string with appropriate units.
func (sw *StopWatch) ElapsedString() string {
	return cookies.FmtDuration(sw.Elapsed(), 3, time.Nanosecond)
}
