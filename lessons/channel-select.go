package lessons

import (
	"fmt"
	"time"
)

func printChannelMsg(msg string, start time.Time) {
	fmt.Println("received", msg, time.Since(start))
}

func ChangeSelect() {
	slowChannel := make(chan string)
	fastChannel := make(chan string)

	start := time.Now()

	go func() {
		time.Sleep(5 * time.Second)
		slowChannel <- "slow message"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fastChannel <- "fast message"
	}()

	fmt.Println(time.Since(start))

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-slowChannel:
			printChannelMsg(msg1, start)
		case msg2 := <-fastChannel:
			printChannelMsg(msg2, start)
		}
	}
}
