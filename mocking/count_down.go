package mocking

import (
	"fmt"
	"io"
	"time"
)

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



func Countdown(out io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, "Go!")
}