package cookies

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	twoMS   time.Duration = time.Duration(2000000)
	threeMS time.Duration = time.Duration(3000000)
)

func TestStopWatch_Start_Stop(t *testing.T) {
	sw := StopWatch{}

	sw.Start()
	time.Sleep(twoMS)
	sw.Stop()

	require.NotEmpty(t, sw.Started)
	require.NotEmpty(t, sw.Stopped)

	assert.True(t, sw.Started.UnixNano() < sw.Stopped.UnixNano())
}

func TestStopWatch_Elapsed(t *testing.T) {
	sw := StopWatch{}

	sw.Start()
	time.Sleep(twoMS)
	sw.Stop()

	elapsed := sw.Elapsed()

	minExpected := twoMS
	assert.True(t, elapsed >= minExpected)

	maxExpected := threeMS
	assert.True(t, elapsed <= maxExpected)
}

func TestStopWatch_Lap(t *testing.T) {
	sw := StopWatch{}
	laps := make([]StopWatch, 3)

	sw.Start()

	time.Sleep(twoMS)
	laps[0] = sw.Lap()

	time.Sleep(twoMS)
	laps[1] = sw.Lap()

	time.Sleep(twoMS)
	laps[2] = sw.Lap()

	prev := StopWatch{
		Stopped: laps[0].Started,
	}
	for _, lap := range laps {
		assert.True(t, prev.Stopped.UnixNano() == lap.Started.UnixNano())
		prev = lap
	}
}
