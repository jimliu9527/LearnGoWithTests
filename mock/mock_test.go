package mock

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func TestCountDown(t *testing.T) {

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &CountdownOperationsSpy{}
		CountDown(buffer, spy)
		fmt.Println(spy.Calls)
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		if got != want {
			t.Errorf("got is '%s' but want is '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("want calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
