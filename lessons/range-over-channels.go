package lessons

import (
	"fmt"
	"time"
)

func pumpIntoChannel(channel chan string, seconds time.Duration) {
	time.Sleep(seconds * time.Second)
	channel <- fmt.Sprintf("Pumped after %d seconds", seconds)
}

func RangeOverChannels() {
	channel := make(chan string, 3)


	for i := 1; i <= 3; i++ {
		go pumpIntoChannel(channel, time.Duration(i))
	}

	
	time.Sleep(2 *  time.Second)
	close(channel)


	for i := range channel {
		fmt.Println(i)
	}
}
