package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const timerCount = 3

// Sleeper interface with a Sleep method
type Sleeper interface {
	Sleep()
}

// SpySleeper struct for testing
type SpySleeper struct {
	Calls int
}

// DefaultSleeper struct for actual sleeping
type DefaultSleeper struct{}

// Sleep method implementation for DefaultSleeper
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep method implementation for SpySleeper to count calls
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Countdown function to count down and print final word
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := timerCount; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

// main function to run the countdown
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	// For testing with SpySleeper
	// spySleeper := &SpySleeper{}
	// Countdown(os.Stdout, spySleeper)
	// fmt.Printf("Sleep calls: %d\n", spySleeper.Calls)
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

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
