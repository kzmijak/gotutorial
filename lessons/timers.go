package lessons

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 is done")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 is done")
	}()

	time.Sleep(1 * time.Second)
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 aborted")
	}

	time.Sleep(2 * time.Second)
}
