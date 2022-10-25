package lessons

import (
	"fmt"
	"time"
)

func worker(done chan bool, shouldFail bool) {
	time.Sleep(time.Second * 2)

	if shouldFail {
		done <- false
	}

	done <- true
}

func Channels() {
	isDone := make(chan bool)
	go worker(isDone, true)

	isSuccess := <-isDone

	fmt.Println("The worker is done, sir. Outcome successful?", isSuccess)
}
