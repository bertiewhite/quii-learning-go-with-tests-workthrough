package countdown

import "time"

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (r *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(t time.Duration) {
	s.durationSlept += t
}

type SleeperOperations struct {
	calls []string
}

func (s *SleeperOperations) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SleeperOperations) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return 0, nil
}

const write = "write"
const sleep = "sleep"
