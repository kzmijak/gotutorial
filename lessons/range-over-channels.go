package lessons

import "fmt"

func RangeOverChannels() {
	channel := make(chan string, 2)

	channel <- "msg1"
	channel <- "msg2"

	close(channel)

	for i := range channel {
		fmt.Println(i)
	}
}
