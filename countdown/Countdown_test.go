package countdown

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SleeperOperations{}

	Countdown(buffer, sleeper)

	exp := `3
2
1
Go!`
	assert.Equal(t, exp, buffer.String())
	assert.Equal(t, 3, len(sleeper.calls))
}

func TestCountDownOrdering(t *testing.T) {
	sleeperWriter := &SleeperOperations{}

	Countdown(sleeperWriter, sleeperWriter)
	expCalls := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	assert.Equal(t, expCalls, sleeperWriter.calls)
}

func TestConfigurableSleeper_Sleep(t *testing.T) {
	spyTime := SpyTime{durationSlept: 0}
	sleeper := ConfigurableSleeper{
		duration: 5 * time.Second,
		sleep:    spyTime.Sleep,
	}

	sleeper.Sleep()

	expSleep := 5 * time.Second
	assert.Equal(t, expSleep, spyTime.durationSlept)

	sleeper.Sleep()
	sleeper.Sleep()
	expSleep = 15 * time.Second
	assert.Equal(t, expSleep, spyTime.durationSlept)
}
