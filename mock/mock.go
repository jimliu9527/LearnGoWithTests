package mock

import (
	"fmt"
	"io"
	// "time"
)

type Sleeper interface {
	Sleep()
}

// type ConfigurableSleeper struct {
// 	duration time.Duration
// }

// func (o *ConfigurableSleeper) Sleep() {
// 	time.Sleep(o.duration)
// }

const finalWord = "Go!"
const countdownStart = 3

// var configurableSleeper = ConfigurableSleeper{duration: time.Second}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
