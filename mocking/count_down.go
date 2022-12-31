package mocking

import (
	"fmt"
	"io"
	"time"
)
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return 
}

// type ConfigurableSleeper struct {
// 	duration time.Duration
// 	sleep    func(time.Duration)
// }





func Countdown(out io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, "Go!")
}